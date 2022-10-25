package csiconfigobservercontroller

import (
	"strings"

	"k8s.io/client-go/tools/cache"

	configinformers "github.com/uccps-samples/client-go/config/informers/externalversions"
	configlistersv1 "github.com/uccps-samples/client-go/config/listers/config/v1"

	"github.com/uccps-samples/library-go/pkg/controller/factory"
	"github.com/uccps-samples/library-go/pkg/operator/configobserver"
	"github.com/uccps-samples/library-go/pkg/operator/configobserver/proxy"
	"github.com/uccps-samples/library-go/pkg/operator/events"
	"github.com/uccps-samples/library-go/pkg/operator/resourcesynccontroller"
	"github.com/uccps-samples/library-go/pkg/operator/v1helpers"
)

// ProxyConfigPath returns the path for the observed proxy config. This is a
// function to avoid exposing a slice that could potentially be appended.
func ProxyConfigPath() []string {
	return []string{"targetcsiconfig", "proxy"}
}

// Listers implement the configobserver.Listers interface.
type Listers struct {
	ProxyLister_ configlistersv1.ProxyLister

	ResourceSync       resourcesynccontroller.ResourceSyncer
	PreRunCachesSynced []cache.InformerSynced
}

func (l Listers) ProxyLister() configlistersv1.ProxyLister {
	return l.ProxyLister_
}

func (l Listers) ResourceSyncer() resourcesynccontroller.ResourceSyncer {
	return l.ResourceSync
}

func (l Listers) PreRunHasSynced() []cache.InformerSynced {
	return l.PreRunCachesSynced
}

// CISConfigObserverController watches information that's relevant to CSI driver operators.
// For now it only observes proxy information, (through the proxy.config.uccp.io/cluster
// object), but more will be added.
type CSIConfigObserverController struct {
	factory.Controller
}

// NewCSIConfigObserverController returns a new CSIConfigObserverController.
func NewCSIConfigObserverController(
	name string,
	operatorClient v1helpers.OperatorClient,
	configinformers configinformers.SharedInformerFactory,
	eventRecorder events.Recorder,
) *CSIConfigObserverController {
	informers := []factory.Informer{
		operatorClient.Informer(),
		configinformers.Config().V1().Proxies().Informer(),
	}

	c := &CSIConfigObserverController{
		Controller: configobserver.NewConfigObserver(
			operatorClient,
			eventRecorder.WithComponentSuffix("csi-config-observer-controller-"+strings.ToLower(name)),
			Listers{
				ProxyLister_: configinformers.Config().V1().Proxies().Lister(),
				PreRunCachesSynced: append([]cache.InformerSynced{},
					operatorClient.Informer().HasSynced,
					configinformers.Config().V1().Proxies().Informer().HasSynced,
				),
			},
			informers,
			proxy.NewProxyObserveFunc(ProxyConfigPath()),
		),
	}

	return c
}
