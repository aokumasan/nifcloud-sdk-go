package datapipeline

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"github.com/awslabs/aws-sdk-go/aws"
)

// ActivatePipelineRequest generates a request for the ActivatePipeline operation.
func (c *DataPipeline) ActivatePipelineRequest(input *ActivatePipelineInput) (req *aws.Request, output *ActivatePipelineOutput) {
	if opActivatePipeline == nil {
		opActivatePipeline = &aws.Operation{
			Name:       "ActivatePipeline",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opActivatePipeline, input, output)
	output = &ActivatePipelineOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) ActivatePipeline(input *ActivatePipelineInput) (output *ActivatePipelineOutput, err error) {
	req, out := c.ActivatePipelineRequest(input)
	output = out
	err = req.Send()
	return
}

var opActivatePipeline *aws.Operation

// CreatePipelineRequest generates a request for the CreatePipeline operation.
func (c *DataPipeline) CreatePipelineRequest(input *CreatePipelineInput) (req *aws.Request, output *CreatePipelineOutput) {
	if opCreatePipeline == nil {
		opCreatePipeline = &aws.Operation{
			Name:       "CreatePipeline",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCreatePipeline, input, output)
	output = &CreatePipelineOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) CreatePipeline(input *CreatePipelineInput) (output *CreatePipelineOutput, err error) {
	req, out := c.CreatePipelineRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreatePipeline *aws.Operation

// DeletePipelineRequest generates a request for the DeletePipeline operation.
func (c *DataPipeline) DeletePipelineRequest(input *DeletePipelineInput) (req *aws.Request) {
	if opDeletePipeline == nil {
		opDeletePipeline = &aws.Operation{
			Name:       "DeletePipeline",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeletePipeline, input, nil)

	return
}

func (c *DataPipeline) DeletePipeline(input *DeletePipelineInput) (err error) {
	req := c.DeletePipelineRequest(input)
	err = req.Send()
	return
}

var opDeletePipeline *aws.Operation

// DescribeObjectsRequest generates a request for the DescribeObjects operation.
func (c *DataPipeline) DescribeObjectsRequest(input *DescribeObjectsInput) (req *aws.Request, output *DescribeObjectsOutput) {
	if opDescribeObjects == nil {
		opDescribeObjects = &aws.Operation{
			Name:       "DescribeObjects",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeObjects, input, output)
	output = &DescribeObjectsOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) DescribeObjects(input *DescribeObjectsInput) (output *DescribeObjectsOutput, err error) {
	req, out := c.DescribeObjectsRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeObjects *aws.Operation

// DescribePipelinesRequest generates a request for the DescribePipelines operation.
func (c *DataPipeline) DescribePipelinesRequest(input *DescribePipelinesInput) (req *aws.Request, output *DescribePipelinesOutput) {
	if opDescribePipelines == nil {
		opDescribePipelines = &aws.Operation{
			Name:       "DescribePipelines",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribePipelines, input, output)
	output = &DescribePipelinesOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) DescribePipelines(input *DescribePipelinesInput) (output *DescribePipelinesOutput, err error) {
	req, out := c.DescribePipelinesRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribePipelines *aws.Operation

// EvaluateExpressionRequest generates a request for the EvaluateExpression operation.
func (c *DataPipeline) EvaluateExpressionRequest(input *EvaluateExpressionInput) (req *aws.Request, output *EvaluateExpressionOutput) {
	if opEvaluateExpression == nil {
		opEvaluateExpression = &aws.Operation{
			Name:       "EvaluateExpression",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opEvaluateExpression, input, output)
	output = &EvaluateExpressionOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) EvaluateExpression(input *EvaluateExpressionInput) (output *EvaluateExpressionOutput, err error) {
	req, out := c.EvaluateExpressionRequest(input)
	output = out
	err = req.Send()
	return
}

var opEvaluateExpression *aws.Operation

// GetPipelineDefinitionRequest generates a request for the GetPipelineDefinition operation.
func (c *DataPipeline) GetPipelineDefinitionRequest(input *GetPipelineDefinitionInput) (req *aws.Request, output *GetPipelineDefinitionOutput) {
	if opGetPipelineDefinition == nil {
		opGetPipelineDefinition = &aws.Operation{
			Name:       "GetPipelineDefinition",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetPipelineDefinition, input, output)
	output = &GetPipelineDefinitionOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) GetPipelineDefinition(input *GetPipelineDefinitionInput) (output *GetPipelineDefinitionOutput, err error) {
	req, out := c.GetPipelineDefinitionRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetPipelineDefinition *aws.Operation

// ListPipelinesRequest generates a request for the ListPipelines operation.
func (c *DataPipeline) ListPipelinesRequest(input *ListPipelinesInput) (req *aws.Request, output *ListPipelinesOutput) {
	if opListPipelines == nil {
		opListPipelines = &aws.Operation{
			Name:       "ListPipelines",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListPipelines, input, output)
	output = &ListPipelinesOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) ListPipelines(input *ListPipelinesInput) (output *ListPipelinesOutput, err error) {
	req, out := c.ListPipelinesRequest(input)
	output = out
	err = req.Send()
	return
}

var opListPipelines *aws.Operation

// PollForTaskRequest generates a request for the PollForTask operation.
func (c *DataPipeline) PollForTaskRequest(input *PollForTaskInput) (req *aws.Request, output *PollForTaskOutput) {
	if opPollForTask == nil {
		opPollForTask = &aws.Operation{
			Name:       "PollForTask",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opPollForTask, input, output)
	output = &PollForTaskOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) PollForTask(input *PollForTaskInput) (output *PollForTaskOutput, err error) {
	req, out := c.PollForTaskRequest(input)
	output = out
	err = req.Send()
	return
}

var opPollForTask *aws.Operation

// PutPipelineDefinitionRequest generates a request for the PutPipelineDefinition operation.
func (c *DataPipeline) PutPipelineDefinitionRequest(input *PutPipelineDefinitionInput) (req *aws.Request, output *PutPipelineDefinitionOutput) {
	if opPutPipelineDefinition == nil {
		opPutPipelineDefinition = &aws.Operation{
			Name:       "PutPipelineDefinition",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opPutPipelineDefinition, input, output)
	output = &PutPipelineDefinitionOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) PutPipelineDefinition(input *PutPipelineDefinitionInput) (output *PutPipelineDefinitionOutput, err error) {
	req, out := c.PutPipelineDefinitionRequest(input)
	output = out
	err = req.Send()
	return
}

var opPutPipelineDefinition *aws.Operation

// QueryObjectsRequest generates a request for the QueryObjects operation.
func (c *DataPipeline) QueryObjectsRequest(input *QueryObjectsInput) (req *aws.Request, output *QueryObjectsOutput) {
	if opQueryObjects == nil {
		opQueryObjects = &aws.Operation{
			Name:       "QueryObjects",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opQueryObjects, input, output)
	output = &QueryObjectsOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) QueryObjects(input *QueryObjectsInput) (output *QueryObjectsOutput, err error) {
	req, out := c.QueryObjectsRequest(input)
	output = out
	err = req.Send()
	return
}

var opQueryObjects *aws.Operation

// ReportTaskProgressRequest generates a request for the ReportTaskProgress operation.
func (c *DataPipeline) ReportTaskProgressRequest(input *ReportTaskProgressInput) (req *aws.Request, output *ReportTaskProgressOutput) {
	if opReportTaskProgress == nil {
		opReportTaskProgress = &aws.Operation{
			Name:       "ReportTaskProgress",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opReportTaskProgress, input, output)
	output = &ReportTaskProgressOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) ReportTaskProgress(input *ReportTaskProgressInput) (output *ReportTaskProgressOutput, err error) {
	req, out := c.ReportTaskProgressRequest(input)
	output = out
	err = req.Send()
	return
}

var opReportTaskProgress *aws.Operation

// ReportTaskRunnerHeartbeatRequest generates a request for the ReportTaskRunnerHeartbeat operation.
func (c *DataPipeline) ReportTaskRunnerHeartbeatRequest(input *ReportTaskRunnerHeartbeatInput) (req *aws.Request, output *ReportTaskRunnerHeartbeatOutput) {
	if opReportTaskRunnerHeartbeat == nil {
		opReportTaskRunnerHeartbeat = &aws.Operation{
			Name:       "ReportTaskRunnerHeartbeat",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opReportTaskRunnerHeartbeat, input, output)
	output = &ReportTaskRunnerHeartbeatOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) ReportTaskRunnerHeartbeat(input *ReportTaskRunnerHeartbeatInput) (output *ReportTaskRunnerHeartbeatOutput, err error) {
	req, out := c.ReportTaskRunnerHeartbeatRequest(input)
	output = out
	err = req.Send()
	return
}

var opReportTaskRunnerHeartbeat *aws.Operation

// SetStatusRequest generates a request for the SetStatus operation.
func (c *DataPipeline) SetStatusRequest(input *SetStatusInput) (req *aws.Request) {
	if opSetStatus == nil {
		opSetStatus = &aws.Operation{
			Name:       "SetStatus",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSetStatus, input, nil)

	return
}

func (c *DataPipeline) SetStatus(input *SetStatusInput) (err error) {
	req := c.SetStatusRequest(input)
	err = req.Send()
	return
}

var opSetStatus *aws.Operation

// SetTaskStatusRequest generates a request for the SetTaskStatus operation.
func (c *DataPipeline) SetTaskStatusRequest(input *SetTaskStatusInput) (req *aws.Request, output *SetTaskStatusOutput) {
	if opSetTaskStatus == nil {
		opSetTaskStatus = &aws.Operation{
			Name:       "SetTaskStatus",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSetTaskStatus, input, output)
	output = &SetTaskStatusOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) SetTaskStatus(input *SetTaskStatusInput) (output *SetTaskStatusOutput, err error) {
	req, out := c.SetTaskStatusRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetTaskStatus *aws.Operation

// ValidatePipelineDefinitionRequest generates a request for the ValidatePipelineDefinition operation.
func (c *DataPipeline) ValidatePipelineDefinitionRequest(input *ValidatePipelineDefinitionInput) (req *aws.Request, output *ValidatePipelineDefinitionOutput) {
	if opValidatePipelineDefinition == nil {
		opValidatePipelineDefinition = &aws.Operation{
			Name:       "ValidatePipelineDefinition",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opValidatePipelineDefinition, input, output)
	output = &ValidatePipelineDefinitionOutput{}
	req.Data = output
	return
}

func (c *DataPipeline) ValidatePipelineDefinition(input *ValidatePipelineDefinitionInput) (output *ValidatePipelineDefinitionOutput, err error) {
	req, out := c.ValidatePipelineDefinitionRequest(input)
	output = out
	err = req.Send()
	return
}

var opValidatePipelineDefinition *aws.Operation

type ActivatePipelineInput struct {
	ParameterValues []*ParameterValue `locationName:"parameterValues" type:"list" json:",omitempty"`
	PipelineID      *string           `locationName:"pipelineId" type:"string" json:",omitempty"`

	metadataActivatePipelineInput `json:"-", xml:"-"`
}

type metadataActivatePipelineInput struct {
	SDKShapeTraits bool `type:"structure" required:"pipelineId" json:",omitempty"`
}

type ActivatePipelineOutput struct {
	metadataActivatePipelineOutput `json:"-", xml:"-"`
}

type metadataActivatePipelineOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type CreatePipelineInput struct {
	Description *string `locationName:"description" type:"string" json:",omitempty"`
	Name        *string `locationName:"name" type:"string" json:",omitempty"`
	UniqueID    *string `locationName:"uniqueId" type:"string" json:",omitempty"`

	metadataCreatePipelineInput `json:"-", xml:"-"`
}

type metadataCreatePipelineInput struct {
	SDKShapeTraits bool `type:"structure" required:"name,uniqueId" json:",omitempty"`
}

type CreatePipelineOutput struct {
	PipelineID *string `locationName:"pipelineId" type:"string" json:",omitempty"`

	metadataCreatePipelineOutput `json:"-", xml:"-"`
}

type metadataCreatePipelineOutput struct {
	SDKShapeTraits bool `type:"structure" required:"pipelineId" json:",omitempty"`
}

type DeletePipelineInput struct {
	PipelineID *string `locationName:"pipelineId" type:"string" json:",omitempty"`

	metadataDeletePipelineInput `json:"-", xml:"-"`
}

type metadataDeletePipelineInput struct {
	SDKShapeTraits bool `type:"structure" required:"pipelineId" json:",omitempty"`
}

type DescribeObjectsInput struct {
	EvaluateExpressions *bool     `locationName:"evaluateExpressions" type:"boolean" json:",omitempty"`
	Marker              *string   `locationName:"marker" type:"string" json:",omitempty"`
	ObjectIDs           []*string `locationName:"objectIds" type:"list" json:",omitempty"`
	PipelineID          *string   `locationName:"pipelineId" type:"string" json:",omitempty"`

	metadataDescribeObjectsInput `json:"-", xml:"-"`
}

type metadataDescribeObjectsInput struct {
	SDKShapeTraits bool `type:"structure" required:"pipelineId,objectIds" json:",omitempty"`
}

type DescribeObjectsOutput struct {
	HasMoreResults  *bool             `locationName:"hasMoreResults" type:"boolean" json:",omitempty"`
	Marker          *string           `locationName:"marker" type:"string" json:",omitempty"`
	PipelineObjects []*PipelineObject `locationName:"pipelineObjects" type:"list" json:",omitempty"`

	metadataDescribeObjectsOutput `json:"-", xml:"-"`
}

type metadataDescribeObjectsOutput struct {
	SDKShapeTraits bool `type:"structure" required:"pipelineObjects" json:",omitempty"`
}

type DescribePipelinesInput struct {
	PipelineIDs []*string `locationName:"pipelineIds" type:"list" json:",omitempty"`

	metadataDescribePipelinesInput `json:"-", xml:"-"`
}

type metadataDescribePipelinesInput struct {
	SDKShapeTraits bool `type:"structure" required:"pipelineIds" json:",omitempty"`
}

type DescribePipelinesOutput struct {
	PipelineDescriptionList []*PipelineDescription `locationName:"pipelineDescriptionList" type:"list" json:",omitempty"`

	metadataDescribePipelinesOutput `json:"-", xml:"-"`
}

type metadataDescribePipelinesOutput struct {
	SDKShapeTraits bool `type:"structure" required:"pipelineDescriptionList" json:",omitempty"`
}

type EvaluateExpressionInput struct {
	Expression *string `locationName:"expression" type:"string" json:",omitempty"`
	ObjectID   *string `locationName:"objectId" type:"string" json:",omitempty"`
	PipelineID *string `locationName:"pipelineId" type:"string" json:",omitempty"`

	metadataEvaluateExpressionInput `json:"-", xml:"-"`
}

type metadataEvaluateExpressionInput struct {
	SDKShapeTraits bool `type:"structure" required:"pipelineId,objectId,expression" json:",omitempty"`
}

type EvaluateExpressionOutput struct {
	EvaluatedExpression *string `locationName:"evaluatedExpression" type:"string" json:",omitempty"`

	metadataEvaluateExpressionOutput `json:"-", xml:"-"`
}

type metadataEvaluateExpressionOutput struct {
	SDKShapeTraits bool `type:"structure" required:"evaluatedExpression" json:",omitempty"`
}

type Field struct {
	Key         *string `locationName:"key" type:"string" json:",omitempty"`
	RefValue    *string `locationName:"refValue" type:"string" json:",omitempty"`
	StringValue *string `locationName:"stringValue" type:"string" json:",omitempty"`

	metadataField `json:"-", xml:"-"`
}

type metadataField struct {
	SDKShapeTraits bool `type:"structure" required:"key" json:",omitempty"`
}

type GetPipelineDefinitionInput struct {
	PipelineID *string `locationName:"pipelineId" type:"string" json:",omitempty"`
	Version    *string `locationName:"version" type:"string" json:",omitempty"`

	metadataGetPipelineDefinitionInput `json:"-", xml:"-"`
}

type metadataGetPipelineDefinitionInput struct {
	SDKShapeTraits bool `type:"structure" required:"pipelineId" json:",omitempty"`
}

type GetPipelineDefinitionOutput struct {
	ParameterObjects []*ParameterObject `locationName:"parameterObjects" type:"list" json:",omitempty"`
	ParameterValues  []*ParameterValue  `locationName:"parameterValues" type:"list" json:",omitempty"`
	PipelineObjects  []*PipelineObject  `locationName:"pipelineObjects" type:"list" json:",omitempty"`

	metadataGetPipelineDefinitionOutput `json:"-", xml:"-"`
}

type metadataGetPipelineDefinitionOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InstanceIdentity struct {
	Document  *string `locationName:"document" type:"string" json:",omitempty"`
	Signature *string `locationName:"signature" type:"string" json:",omitempty"`

	metadataInstanceIdentity `json:"-", xml:"-"`
}

type metadataInstanceIdentity struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InternalServiceError struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataInternalServiceError `json:"-", xml:"-"`
}

type metadataInternalServiceError struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InvalidRequestException struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataInvalidRequestException `json:"-", xml:"-"`
}

type metadataInvalidRequestException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListPipelinesInput struct {
	Marker *string `locationName:"marker" type:"string" json:",omitempty"`

	metadataListPipelinesInput `json:"-", xml:"-"`
}

type metadataListPipelinesInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListPipelinesOutput struct {
	HasMoreResults *bool             `locationName:"hasMoreResults" type:"boolean" json:",omitempty"`
	Marker         *string           `locationName:"marker" type:"string" json:",omitempty"`
	PipelineIDList []*PipelineIDName `locationName:"pipelineIdList" type:"list" json:",omitempty"`

	metadataListPipelinesOutput `json:"-", xml:"-"`
}

type metadataListPipelinesOutput struct {
	SDKShapeTraits bool `type:"structure" required:"pipelineIdList" json:",omitempty"`
}

type Operator struct {
	Type   *string   `locationName:"type" type:"string" json:",omitempty"`
	Values []*string `locationName:"values" type:"list" json:",omitempty"`

	metadataOperator `json:"-", xml:"-"`
}

type metadataOperator struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ParameterAttribute struct {
	Key         *string `locationName:"key" type:"string" json:",omitempty"`
	StringValue *string `locationName:"stringValue" type:"string" json:",omitempty"`

	metadataParameterAttribute `json:"-", xml:"-"`
}

type metadataParameterAttribute struct {
	SDKShapeTraits bool `type:"structure" required:"key,stringValue" json:",omitempty"`
}

type ParameterObject struct {
	Attributes []*ParameterAttribute `locationName:"attributes" type:"list" json:",omitempty"`
	ID         *string               `locationName:"id" type:"string" json:",omitempty"`

	metadataParameterObject `json:"-", xml:"-"`
}

type metadataParameterObject struct {
	SDKShapeTraits bool `type:"structure" required:"id,attributes" json:",omitempty"`
}

type ParameterValue struct {
	ID          *string `locationName:"id" type:"string" json:",omitempty"`
	StringValue *string `locationName:"stringValue" type:"string" json:",omitempty"`

	metadataParameterValue `json:"-", xml:"-"`
}

type metadataParameterValue struct {
	SDKShapeTraits bool `type:"structure" required:"id,stringValue" json:",omitempty"`
}

type PipelineDeletedException struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataPipelineDeletedException `json:"-", xml:"-"`
}

type metadataPipelineDeletedException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type PipelineDescription struct {
	Description *string  `locationName:"description" type:"string" json:",omitempty"`
	Fields      []*Field `locationName:"fields" type:"list" json:",omitempty"`
	Name        *string  `locationName:"name" type:"string" json:",omitempty"`
	PipelineID  *string  `locationName:"pipelineId" type:"string" json:",omitempty"`

	metadataPipelineDescription `json:"-", xml:"-"`
}

type metadataPipelineDescription struct {
	SDKShapeTraits bool `type:"structure" required:"pipelineId,name,fields" json:",omitempty"`
}

type PipelineIDName struct {
	ID   *string `locationName:"id" type:"string" json:",omitempty"`
	Name *string `locationName:"name" type:"string" json:",omitempty"`

	metadataPipelineIDName `json:"-", xml:"-"`
}

type metadataPipelineIDName struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type PipelineNotFoundException struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataPipelineNotFoundException `json:"-", xml:"-"`
}

type metadataPipelineNotFoundException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type PipelineObject struct {
	Fields []*Field `locationName:"fields" type:"list" json:",omitempty"`
	ID     *string  `locationName:"id" type:"string" json:",omitempty"`
	Name   *string  `locationName:"name" type:"string" json:",omitempty"`

	metadataPipelineObject `json:"-", xml:"-"`
}

type metadataPipelineObject struct {
	SDKShapeTraits bool `type:"structure" required:"id,name,fields" json:",omitempty"`
}

type PollForTaskInput struct {
	Hostname         *string           `locationName:"hostname" type:"string" json:",omitempty"`
	InstanceIdentity *InstanceIdentity `locationName:"instanceIdentity" type:"structure" json:",omitempty"`
	WorkerGroup      *string           `locationName:"workerGroup" type:"string" json:",omitempty"`

	metadataPollForTaskInput `json:"-", xml:"-"`
}

type metadataPollForTaskInput struct {
	SDKShapeTraits bool `type:"structure" required:"workerGroup" json:",omitempty"`
}

type PollForTaskOutput struct {
	TaskObject *TaskObject `locationName:"taskObject" type:"structure" json:",omitempty"`

	metadataPollForTaskOutput `json:"-", xml:"-"`
}

type metadataPollForTaskOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type PutPipelineDefinitionInput struct {
	ParameterObjects []*ParameterObject `locationName:"parameterObjects" type:"list" json:",omitempty"`
	ParameterValues  []*ParameterValue  `locationName:"parameterValues" type:"list" json:",omitempty"`
	PipelineID       *string            `locationName:"pipelineId" type:"string" json:",omitempty"`
	PipelineObjects  []*PipelineObject  `locationName:"pipelineObjects" type:"list" json:",omitempty"`

	metadataPutPipelineDefinitionInput `json:"-", xml:"-"`
}

type metadataPutPipelineDefinitionInput struct {
	SDKShapeTraits bool `type:"structure" required:"pipelineId,pipelineObjects" json:",omitempty"`
}

type PutPipelineDefinitionOutput struct {
	Errored            *bool                `locationName:"errored" type:"boolean" json:",omitempty"`
	ValidationErrors   []*ValidationError   `locationName:"validationErrors" type:"list" json:",omitempty"`
	ValidationWarnings []*ValidationWarning `locationName:"validationWarnings" type:"list" json:",omitempty"`

	metadataPutPipelineDefinitionOutput `json:"-", xml:"-"`
}

type metadataPutPipelineDefinitionOutput struct {
	SDKShapeTraits bool `type:"structure" required:"errored" json:",omitempty"`
}

type Query struct {
	Selectors []*Selector `locationName:"selectors" type:"list" json:",omitempty"`

	metadataQuery `json:"-", xml:"-"`
}

type metadataQuery struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type QueryObjectsInput struct {
	Limit      *int    `locationName:"limit" type:"integer" json:",omitempty"`
	Marker     *string `locationName:"marker" type:"string" json:",omitempty"`
	PipelineID *string `locationName:"pipelineId" type:"string" json:",omitempty"`
	Query      *Query  `locationName:"query" type:"structure" json:",omitempty"`
	Sphere     *string `locationName:"sphere" type:"string" json:",omitempty"`

	metadataQueryObjectsInput `json:"-", xml:"-"`
}

type metadataQueryObjectsInput struct {
	SDKShapeTraits bool `type:"structure" required:"pipelineId,sphere" json:",omitempty"`
}

type QueryObjectsOutput struct {
	HasMoreResults *bool     `locationName:"hasMoreResults" type:"boolean" json:",omitempty"`
	IDs            []*string `locationName:"ids" type:"list" json:",omitempty"`
	Marker         *string   `locationName:"marker" type:"string" json:",omitempty"`

	metadataQueryObjectsOutput `json:"-", xml:"-"`
}

type metadataQueryObjectsOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ReportTaskProgressInput struct {
	Fields []*Field `locationName:"fields" type:"list" json:",omitempty"`
	TaskID *string  `locationName:"taskId" type:"string" json:",omitempty"`

	metadataReportTaskProgressInput `json:"-", xml:"-"`
}

type metadataReportTaskProgressInput struct {
	SDKShapeTraits bool `type:"structure" required:"taskId" json:",omitempty"`
}

type ReportTaskProgressOutput struct {
	Canceled *bool `locationName:"canceled" type:"boolean" json:",omitempty"`

	metadataReportTaskProgressOutput `json:"-", xml:"-"`
}

type metadataReportTaskProgressOutput struct {
	SDKShapeTraits bool `type:"structure" required:"canceled" json:",omitempty"`
}

type ReportTaskRunnerHeartbeatInput struct {
	Hostname     *string `locationName:"hostname" type:"string" json:",omitempty"`
	TaskRunnerID *string `locationName:"taskrunnerId" type:"string" json:",omitempty"`
	WorkerGroup  *string `locationName:"workerGroup" type:"string" json:",omitempty"`

	metadataReportTaskRunnerHeartbeatInput `json:"-", xml:"-"`
}

type metadataReportTaskRunnerHeartbeatInput struct {
	SDKShapeTraits bool `type:"structure" required:"taskrunnerId" json:",omitempty"`
}

type ReportTaskRunnerHeartbeatOutput struct {
	Terminate *bool `locationName:"terminate" type:"boolean" json:",omitempty"`

	metadataReportTaskRunnerHeartbeatOutput `json:"-", xml:"-"`
}

type metadataReportTaskRunnerHeartbeatOutput struct {
	SDKShapeTraits bool `type:"structure" required:"terminate" json:",omitempty"`
}

type Selector struct {
	FieldName *string   `locationName:"fieldName" type:"string" json:",omitempty"`
	Operator  *Operator `locationName:"operator" type:"structure" json:",omitempty"`

	metadataSelector `json:"-", xml:"-"`
}

type metadataSelector struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type SetStatusInput struct {
	ObjectIDs  []*string `locationName:"objectIds" type:"list" json:",omitempty"`
	PipelineID *string   `locationName:"pipelineId" type:"string" json:",omitempty"`
	Status     *string   `locationName:"status" type:"string" json:",omitempty"`

	metadataSetStatusInput `json:"-", xml:"-"`
}

type metadataSetStatusInput struct {
	SDKShapeTraits bool `type:"structure" required:"pipelineId,objectIds,status" json:",omitempty"`
}

type SetTaskStatusInput struct {
	ErrorID         *string `locationName:"errorId" type:"string" json:",omitempty"`
	ErrorMessage    *string `locationName:"errorMessage" type:"string" json:",omitempty"`
	ErrorStackTrace *string `locationName:"errorStackTrace" type:"string" json:",omitempty"`
	TaskID          *string `locationName:"taskId" type:"string" json:",omitempty"`
	TaskStatus      *string `locationName:"taskStatus" type:"string" json:",omitempty"`

	metadataSetTaskStatusInput `json:"-", xml:"-"`
}

type metadataSetTaskStatusInput struct {
	SDKShapeTraits bool `type:"structure" required:"taskId,taskStatus" json:",omitempty"`
}

type SetTaskStatusOutput struct {
	metadataSetTaskStatusOutput `json:"-", xml:"-"`
}

type metadataSetTaskStatusOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type TaskNotFoundException struct {
	Message *string `locationName:"message" type:"string" json:",omitempty"`

	metadataTaskNotFoundException `json:"-", xml:"-"`
}

type metadataTaskNotFoundException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type TaskObject struct {
	AttemptID  *string                     `locationName:"attemptId" type:"string" json:",omitempty"`
	Objects    *map[string]*PipelineObject `locationName:"objects" type:"map" json:",omitempty"`
	PipelineID *string                     `locationName:"pipelineId" type:"string" json:",omitempty"`
	TaskID     *string                     `locationName:"taskId" type:"string" json:",omitempty"`

	metadataTaskObject `json:"-", xml:"-"`
}

type metadataTaskObject struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ValidatePipelineDefinitionInput struct {
	ParameterObjects []*ParameterObject `locationName:"parameterObjects" type:"list" json:",omitempty"`
	ParameterValues  []*ParameterValue  `locationName:"parameterValues" type:"list" json:",omitempty"`
	PipelineID       *string            `locationName:"pipelineId" type:"string" json:",omitempty"`
	PipelineObjects  []*PipelineObject  `locationName:"pipelineObjects" type:"list" json:",omitempty"`

	metadataValidatePipelineDefinitionInput `json:"-", xml:"-"`
}

type metadataValidatePipelineDefinitionInput struct {
	SDKShapeTraits bool `type:"structure" required:"pipelineId,pipelineObjects" json:",omitempty"`
}

type ValidatePipelineDefinitionOutput struct {
	Errored            *bool                `locationName:"errored" type:"boolean" json:",omitempty"`
	ValidationErrors   []*ValidationError   `locationName:"validationErrors" type:"list" json:",omitempty"`
	ValidationWarnings []*ValidationWarning `locationName:"validationWarnings" type:"list" json:",omitempty"`

	metadataValidatePipelineDefinitionOutput `json:"-", xml:"-"`
}

type metadataValidatePipelineDefinitionOutput struct {
	SDKShapeTraits bool `type:"structure" required:"errored" json:",omitempty"`
}

type ValidationError struct {
	Errors []*string `locationName:"errors" type:"list" json:",omitempty"`
	ID     *string   `locationName:"id" type:"string" json:",omitempty"`

	metadataValidationError `json:"-", xml:"-"`
}

type metadataValidationError struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ValidationWarning struct {
	ID       *string   `locationName:"id" type:"string" json:",omitempty"`
	Warnings []*string `locationName:"warnings" type:"list" json:",omitempty"`

	metadataValidationWarning `json:"-", xml:"-"`
}

type metadataValidationWarning struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}