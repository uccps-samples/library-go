package csidrivercontrollerservicecontroller

import (
	"fmt"
	"os"
	"testing"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	coreinformers "k8s.io/client-go/informers"
	fakecore "k8s.io/client-go/kubernetes/fake"

	"github.com/ghodss/yaml"
	"github.com/google/go-cmp/cmp"

	configv1 "github.com/uccps-samples/api/config/v1"
	fakeconfig "github.com/uccps-samples/client-go/config/clientset/versioned/fake"
	configinformers "github.com/uccps-samples/client-go/config/informers/externalversions"
	"github.com/uccps-samples/library-go/pkg/operator/csi/csiconfigobservercontroller"
	"github.com/uccps-samples/library-go/pkg/operator/resource/resourceread"
)

const (
	defaultContainerName  = "csi-driver"
	defaultHTTPProxyValue = "http://foo.bar.proxy"
)

func TestWithObservedProxyDeploymentHook(t *testing.T) {
	const (
		replica0 = 0
		replica1 = 1
		replica2 = 2
	)
	var (
		argsLevel2 = 2
	)
	testCases := []struct {
		name               string
		initialDriver      *fakeDriverInstance
		initialDeployment  *appsv1.Deployment
		expectedDeployment *appsv1.Deployment
		expectError        bool
	}{
		{
			name:          "no observed proxy config",
			initialDriver: makeFakeDriverInstance(), // CR has no observed proxy config
			initialDeployment: makeDeployment(
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentHTTPProxyAnnotation(defaultContainerName),
				withDeploymentGeneration(1, 0)),
			expectedDeployment: makeDeployment( // no container has proxy ENV set
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentHTTPProxyAnnotation(defaultContainerName),
				withDeploymentGeneration(1, 0)),
			expectError: false,
		},
		{
			name: "observed proxy config, annotation present",
			initialDriver: makeFakeDriverInstance(
				withObservedHTTPProxy(defaultHTTPProxyValue, nil /* config path*/),
			),
			initialDeployment: makeDeployment(
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentHTTPProxyAnnotation(defaultContainerName),
				withDeploymentGeneration(1, 0)),
			expectedDeployment: makeDeployment(
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentGeneration(1, 0),
				withDeploymentHTTPProxyAnnotation(defaultContainerName),
				withDeploymentHTTPProxyEnv(defaultHTTPProxyValue, defaultContainerName)), // proxy ENV was added to container
			expectError: false,
		},
		{
			name: "observed proxy config, annotation present with WRONG container name",
			initialDriver: makeFakeDriverInstance(
				withObservedHTTPProxy(defaultHTTPProxyValue, nil /* config path*/),
			),
			initialDeployment: makeDeployment(
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentHTTPProxyAnnotation("csi-driver-non-existent"), // this container doesn't exist
				withDeploymentGeneration(1, 0)),
			expectedDeployment: makeDeployment( // no container has proxy ENV is set
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentGeneration(1, 0),
				withDeploymentHTTPProxyAnnotation("csi-driver-non-existent")),
			expectError: false,
		},
		{
			name: "observed proxy config, annotation NOT present",
			initialDriver: makeFakeDriverInstance(
				withObservedHTTPProxy(defaultHTTPProxyValue, nil /* config path*/),
			),
			initialDeployment: makeDeployment( // inject-proxy annotation not added
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentGeneration(1, 0)),
			expectedDeployment: makeDeployment( // no container has proxy ENV set
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentGeneration(1, 0)),
			expectError: false,
		},
		{
			name: "invalid observed proxy config",
			initialDriver: makeFakeDriverInstance(
				withInvalidObservedHTTPProxy(defaultHTTPProxyValue, nil /* config path*/),
			),
			initialDeployment: makeDeployment(
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentHTTPProxyAnnotation(defaultContainerName),
				withDeploymentGeneration(1, 0)),
			expectedDeployment: makeDeployment( // no container has proxy ENV set
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentHTTPProxyAnnotation(defaultContainerName),
				withDeploymentGeneration(1, 0)),
			expectError: true, // report an error
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fn := WithObservedProxyDeploymentHook()
			err := fn(&tc.initialDriver.Spec, tc.initialDeployment)
			if err != nil && !tc.expectError {
				t.Errorf("Expected no error running hook function, got: %v", err)

			}
			if !equality.Semantic.DeepEqual(tc.initialDeployment, tc.expectedDeployment) {
				t.Errorf("Unexpected Deployment content:\n%s", cmp.Diff(tc.initialDeployment, tc.expectedDeployment))
			}
		})
	}
}

func TestWithReplicasHook(t *testing.T) {
	var (
		argsLevel2       = 2
		masterNodeLabels = map[string]string{"node-role.kubernetes.io/master": ""}
		workerNodeLabels = map[string]string{"node-role.kubernetes.io/worker": ""}
	)
	testCases := []struct {
		name               string
		initialDriver      *fakeDriverInstance
		initialNodes       []*v1.Node
		initialDeployment  *appsv1.Deployment
		expectedDeployment *appsv1.Deployment
		expectError        bool
	}{
		{
			name:          "three-node control-plane",
			initialDriver: makeFakeDriverInstance(),
			initialNodes: []*v1.Node{
				makeNode("A", masterNodeLabels),
				makeNode("B", masterNodeLabels),
				makeNode("C", masterNodeLabels),
			},
			initialDeployment: makeDeployment(
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentReplicas(1),
				withDeploymentGeneration(1, 0)),
			expectedDeployment: makeDeployment(
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentReplicas(2),
				withDeploymentGeneration(1, 0)),
			expectError: false,
		},
		{
			name:          "three-node control-plane with one worker node",
			initialDriver: makeFakeDriverInstance(),
			initialNodes: []*v1.Node{
				makeNode("A", masterNodeLabels),
				makeNode("B", masterNodeLabels),
				makeNode("C", masterNodeLabels),
				makeNode("D", workerNodeLabels),
			},
			initialDeployment: makeDeployment(
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentReplicas(1),
				withDeploymentGeneration(1, 0)),
			expectedDeployment: makeDeployment(
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentReplicas(2),
				withDeploymentGeneration(1, 0)),
			expectError: false,
		},
		{
			name:          "two-node control-plane",
			initialDriver: makeFakeDriverInstance(),
			initialNodes: []*v1.Node{
				makeNode("A", masterNodeLabels),
				makeNode("B", masterNodeLabels),
			},
			initialDeployment: makeDeployment(
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentReplicas(1),
				withDeploymentGeneration(1, 0)),
			expectedDeployment: makeDeployment(
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentReplicas(2),
				withDeploymentGeneration(1, 0)),
			expectError: false,
		},
		{
			name:          "single-node control-plane with one worker node",
			initialDriver: makeFakeDriverInstance(),
			initialNodes: []*v1.Node{
				makeNode("A", masterNodeLabels),
				makeNode("B", workerNodeLabels),
			},
			initialDeployment: makeDeployment(
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentReplicas(1),
				withDeploymentGeneration(1, 0)),
			expectedDeployment: makeDeployment(
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentReplicas(1),
				withDeploymentGeneration(1, 0)),
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var initialObjects []runtime.Object

			// Add deployment to slice of objects to be added to client and informer
			initialObjects = append(initialObjects, tc.initialDeployment)

			// Do the same with nodes
			for _, node := range tc.initialNodes {
				initialObjects = append(initialObjects, node)
			}

			// Create fake client and informer
			coreClient := fakecore.NewSimpleClientset(initialObjects...)
			coreInformerFactory := coreinformers.NewSharedInformerFactory(coreClient, 0 /*no resync */)

			// Fill the fake informer with the initial deployment and nodes
			coreInformerFactory.Apps().V1().Deployments().Informer().GetIndexer().Add(tc.initialDeployment)
			for _, node := range initialObjects {
				coreInformerFactory.Core().V1().Nodes().Informer().GetIndexer().Add(node)
			}

			fn := WithReplicasHook(coreInformerFactory.Core().V1().Nodes().Lister())
			err := fn(&tc.initialDriver.Spec, tc.initialDeployment)
			if err != nil && !tc.expectError {
				t.Errorf("Expected no error running hook function, got: %v", err)

			}
			if !equality.Semantic.DeepEqual(tc.initialDeployment, tc.expectedDeployment) {
				t.Errorf("Unexpected Deployment content:\n%s", cmp.Diff(tc.initialDeployment, tc.expectedDeployment))
			}
		})
	}
}

func setEnvVars(image images) {
	os.Setenv(driverImageEnvName, image.csiDriver)
	os.Setenv(provisionerImageEnvName, image.provisioner)
	os.Setenv(attacherImageEnvName, image.attacher)
	os.Setenv(snapshotterImageEnvName, image.snapshotter)
	os.Setenv(resizerImageEnvName, image.resizer)
	os.Setenv(livenessProbeImageEnvName, image.livenessProbe)
	os.Setenv(kubeRBACProxyImageEnvName, image.kubeRBACProxy)
}

func TestManifestHooks(t *testing.T) {
	argsLevel2 := 2
	testCases := []struct {
		name               string
		initialOperator    *fakeDriverInstance
		initialManifest    []byte
		image              images
		infraObjectName    string
		expectedDeployment *appsv1.Deployment
		expectError        bool
	}{
		{
			name:            "replace all images",
			initialOperator: makeFakeDriverInstance(),
			initialManifest: makeFakeManifest(),
			image:           defaultImages(),
			infraObjectName: infraConfigName,
			expectedDeployment: makeDeployment(
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentReplicas(1),
				withDeploymentGeneration(0, 0)),
			expectError: false,
		},
		{
			name:            "replace no images, except for CSI Driver image",
			initialOperator: makeFakeDriverInstance(),
			initialManifest: makeFakeManifest(),
			image:           images{csiDriver: "quay.io/openshift/origin-test-csi-driver:latest"},
			infraObjectName: infraConfigName,
			expectedDeployment: makeDeployment(
				defaultClusterID,
				argsLevel2,
				images{csiDriver: "quay.io/openshift/origin-test-csi-driver:latest"},
				withDeploymentReplicas(1),
				withDeploymentGeneration(0, 0)),
			expectError: false,
		},
		{
			name:            "incorrect infra object, so expect an error",
			initialOperator: makeFakeDriverInstance(),
			initialManifest: makeFakeManifest(),
			image:           defaultImages(),
			infraObjectName: "dummy",
			expectedDeployment: makeDeployment(
				defaultClusterID,
				argsLevel2,
				defaultImages(),
				withDeploymentReplicas(1),
				withDeploymentGeneration(0, 0)),
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			setEnvVars(tc.image)
			infraObj := makeInfra()
			infraObj.ObjectMeta.Name = tc.infraObjectName
			initialInfras := []runtime.Object{infraObj}
			configClient := fakeconfig.NewSimpleClientset(initialInfras...)
			configInformerFactory := configinformers.NewSharedInformerFactory(configClient, 0)
			configInformerFactory.Config().V1().Infrastructures().Informer().GetIndexer().Add(initialInfras[0])

			fn := WithPlaceholdersHook(configInformerFactory)
			manifest, err := fn(&tc.initialOperator.Spec, tc.initialManifest)
			if err != nil && !tc.expectError {
				t.Errorf("Expected no error running hook function, got: %v", err)
			}
			if len(manifest) > 0 && !equality.Semantic.DeepEqual(resourceread.ReadDeploymentV1OrDie(manifest), tc.expectedDeployment) {
				t.Errorf("Unexpected Deployment content:\n%s", cmp.Diff(resourceread.ReadDeploymentV1OrDie(manifest), tc.expectedDeployment))
			}
		})
	}
}

func TestWithControlPlaneTopologyHook(t *testing.T) {
	argsLevel2 := 2
	tests := []struct {
		name               string
		infra              *configv1.Infrastructure
		initialDeployment  *appsv1.Deployment
		expectedDeployment *appsv1.Deployment
	}{
		{
			name:               "highly available topology",
			infra:              makeInfraWithCPTopology(configv1.HighlyAvailableTopologyMode),
			initialDeployment:  makeDeployment(defaultClusterID, argsLevel2, defaultImages()),
			expectedDeployment: makeDeployment(defaultClusterID, argsLevel2, defaultImages()),
		},
		{
			name:               "external topology",
			infra:              makeInfraWithCPTopology(configv1.ExternalTopologyMode),
			initialDeployment:  makeDeployment(defaultClusterID, argsLevel2, defaultImages()),
			expectedDeployment: makeDeployment(defaultClusterID, argsLevel2, defaultImages(), withNodeSelector(map[string]string{})),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			initialInfras := []runtime.Object{test.infra}
			configClient := fakeconfig.NewSimpleClientset(initialInfras...)
			configInformerFactory := configinformers.NewSharedInformerFactory(configClient, 0)
			configInformerFactory.Config().V1().Infrastructures().Informer().GetIndexer().Add(initialInfras[0])

			fn := WithControlPlaneTopologyHook(configInformerFactory)
			err := fn(nil, test.initialDeployment)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !equality.Semantic.DeepEqual(test.initialDeployment, test.expectedDeployment) {
				t.Errorf("Unexpected Deployment content:\n%s", cmp.Diff(test.initialDeployment, test.expectedDeployment))
			}
		})
	}
}

func withNodeSelector(nodeSelector map[string]string) deploymentModifier {
	return func(deployment *appsv1.Deployment) *appsv1.Deployment {
		deployment.Spec.Template.Spec.NodeSelector = nodeSelector
		return deployment
	}
}

func makeInfraWithCPTopology(mode configv1.TopologyMode) *configv1.Infrastructure {
	infra := makeInfra()
	infra.Status.ControlPlaneTopology = mode
	return infra
}

func withObservedHTTPProxy(proxy string, path []string) driverModifier {
	if len(path) == 0 {
		path = csiconfigobservercontroller.ProxyConfigPath()
	}
	return func(i *fakeDriverInstance) *fakeDriverInstance {
		observedConfig := map[string]interface{}{}
		unstructured.SetNestedStringMap(observedConfig, map[string]string{"HTTP_PROXY": proxy}, path...)
		d, _ := yaml.Marshal(observedConfig)
		i.Spec.ObservedConfig = runtime.RawExtension{Raw: d, Object: &unstructured.Unstructured{Object: observedConfig}}
		return i
	}
}

func withInvalidObservedHTTPProxy(proxy string, path []string) driverModifier {
	if len(path) == 0 {
		path = csiconfigobservercontroller.ProxyConfigPath()
	}
	return func(i *fakeDriverInstance) *fakeDriverInstance {
		observedConfig := map[string]interface{}{}
		unstructured.SetNestedStringMap(observedConfig, map[string]string{"HTTP_PROXY": proxy}, path...)
		invalidYAML := []byte("[observedConfig:")
		i.Spec.ObservedConfig = runtime.RawExtension{Raw: invalidYAML, Object: &unstructured.Unstructured{Object: observedConfig}}
		return i
	}
}

func withDeploymentHTTPProxyAnnotation(containerName string) deploymentModifier {
	return func(instance *appsv1.Deployment) *appsv1.Deployment {
		instance.Annotations = map[string]string{"config.uccp.io/inject-proxy": containerName}
		return instance
	}
}

func withDeploymentHTTPProxyEnv(proxy, containerName string) deploymentModifier {
	return func(instance *appsv1.Deployment) *appsv1.Deployment {
		containers := instance.Spec.Template.Spec.Containers
		for i := range containers {
			if containers[i].Name == containerName {
				containers[i].Env = append(containers[i].Env, v1.EnvVar{Name: "HTTP_PROXY", Value: proxy})
			}
		}
		return instance
	}
}

func makeNode(suffix string, labels map[string]string) *v1.Node {
	return &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name:   fmt.Sprintf("node-%s", suffix),
			Labels: labels,
		},
	}
}

func TestWithLeaderElectionReplacerHook(t *testing.T) {
	in := `
            - --leader-election-lease-duration=${LEADER_ELECTION_LEASE_DURATION}
            - --leader-election-renew-deadline=${LEADER_ELECTION_RENEW_DEADLINE}
            - --leader-election-retry-period=${LEADER_ELECTION_RETRY_PERIOD}
`
	expected := `
            - --leader-election-lease-duration=1s
            - --leader-election-renew-deadline=2s
            - --leader-election-retry-period=3s
`
	le := configv1.LeaderElection{
		LeaseDuration: metav1.Duration{Duration: time.Second},
		RenewDeadline: metav1.Duration{Duration: 2 * time.Second},
		RetryPeriod:   metav1.Duration{Duration: 3 * time.Second},
	}
	hook := WithLeaderElectionReplacerHook(le)
	out, err := hook(nil, []byte(in))
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(out) != expected {
		t.Errorf("expected %q, got %q", expected, string(out))
	}
}
