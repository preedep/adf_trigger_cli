package azure

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
	"strings"
)

type ADFStatus string

const (
	SUCCESSES ADFStatus = "succeeded"
	FAILED              = "failed"
	CANCELED            = "canceled"
)

type DataFactories struct {
	subscriptionId string
	resourceGroup  string
	factoryName    string
	pipelineName   string
}

func CreateDataFactories(
	subscription_id string,
	resource_group string,
	factory_name string,
) DataFactories {
	return DataFactories{
		subscriptionId: subscription_id,
		resourceGroup:  resource_group,
		factoryName:    factory_name,
	}
}

func (d *DataFactories) RunPipeLine(pipeline_name string, callback func(ADFStatus /**/, string)) error {
	d.pipelineName = pipeline_name
	run_id, err := d.runPipelineClientCreateRun()
	if err == nil {
		fmt.Printf("Pipeline run id : %v\r\n", run_id)
		for {
			status, err := d.waitForStatus(run_id)
			var message string = ""
			if err != nil {
				message = err.Error()
			}
			if callback != nil {
				callback(status, message)
			}
			switch status {
			case SUCCESSES, CANCELED:
				{
					return nil
				}
			case FAILED:
				{
					return err
				}
			}
		}
	}
	return err
}

func (d *DataFactories) runPipelineClientCreateRun() (string, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("failed azure credential %v", err))
		return "", err
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient(d.subscriptionId, cred, nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("failed new pipeline client  %v", err))
		return "", err
	}

	res, err := client.CreateRun(ctx,
		d.resourceGroup,
		d.factoryName,
		d.pipelineName,
		&armdatafactory.PipelinesClientCreateRunOptions{ReferencePipelineRunID: nil,
			IsRecovery:        nil,
			StartActivityName: nil,
			StartFromFailure:  nil,
			Parameters: map[string]interface{}{
				"OutputBlobNameList": []interface{}{
					"exampleoutput.csv",
				},
			},
		})
	if err != nil {
		err = errors.New(fmt.Sprintf("failed run pipeline  %v", err))
		return "", err
	}
	return *res.RunID, nil
}
func (d *DataFactories) waitForStatus(run_id string) (ADFStatus, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("failed credential  %v", err))
		return FAILED, err
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelineRunsClient(d.subscriptionId, cred, nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("failed new pipeline client  %v", err))
		return FAILED, err
	}
	res, err := client.Get(ctx, d.resourceGroup, d.factoryName, run_id, nil)
	if err != nil {
		err = errors.New(fmt.Sprintf("failed get status  %v", err))
		return FAILED, err
	}
	return ADFStatus(strings.ToLower(*res.Status)), nil
}
