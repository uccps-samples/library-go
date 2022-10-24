package staleconditions

import (
	"context"
	"time"

	operatorv1 "github.com/uccps-samples/api/operator/v1"

	"github.com/uccps-samples/library-go/pkg/controller/factory"
	"github.com/uccps-samples/library-go/pkg/operator/events"
	"github.com/uccps-samples/library-go/pkg/operator/v1helpers"
)

type RemoveStaleConditionsController struct {
	conditions     []string
	operatorClient v1helpers.OperatorClient
}

func NewRemoveStaleConditionsController(
	conditions []string,
	operatorClient v1helpers.OperatorClient,
	eventRecorder events.Recorder,
) factory.Controller {
	c := &RemoveStaleConditionsController{
		conditions:     conditions,
		operatorClient: operatorClient,
	}
	return factory.New().ResyncEvery(time.Minute).WithSync(c.sync).WithInformers(operatorClient.Informer()).ToController("RemoveStaleConditionsController", eventRecorder.WithComponentSuffix("remove-stale-conditions"))
}

func (c RemoveStaleConditionsController) sync(ctx context.Context, syncContext factory.SyncContext) error {
	removeStaleConditionsFn := func(status *operatorv1.OperatorStatus) error {
		for _, condition := range c.conditions {
			v1helpers.RemoveOperatorCondition(&status.Conditions, condition)
		}
		return nil
	}

	if _, _, err := v1helpers.UpdateStatus(ctx, c.operatorClient, removeStaleConditionsFn); err != nil {
		return err
	}

	return nil
}
