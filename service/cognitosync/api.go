package cognitosync

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

// DeleteDatasetRequest generates a request for the DeleteDataset operation.
func (c *CognitoSync) DeleteDatasetRequest(input *DeleteDatasetInput) (req *aws.Request, output *DeleteDatasetOutput) {
	if opDeleteDataset == nil {
		opDeleteDataset = &aws.Operation{
			Name:       "DeleteDataset",
			HTTPMethod: "DELETE",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteDataset, input, output)
	output = &DeleteDatasetOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) DeleteDataset(input *DeleteDatasetInput) (output *DeleteDatasetOutput, err error) {
	req, out := c.DeleteDatasetRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteDataset *aws.Operation

// DescribeDatasetRequest generates a request for the DescribeDataset operation.
func (c *CognitoSync) DescribeDatasetRequest(input *DescribeDatasetInput) (req *aws.Request, output *DescribeDatasetOutput) {
	if opDescribeDataset == nil {
		opDescribeDataset = &aws.Operation{
			Name:       "DescribeDataset",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeDataset, input, output)
	output = &DescribeDatasetOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) DescribeDataset(input *DescribeDatasetInput) (output *DescribeDatasetOutput, err error) {
	req, out := c.DescribeDatasetRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeDataset *aws.Operation

// DescribeIdentityPoolUsageRequest generates a request for the DescribeIdentityPoolUsage operation.
func (c *CognitoSync) DescribeIdentityPoolUsageRequest(input *DescribeIdentityPoolUsageInput) (req *aws.Request, output *DescribeIdentityPoolUsageOutput) {
	if opDescribeIdentityPoolUsage == nil {
		opDescribeIdentityPoolUsage = &aws.Operation{
			Name:       "DescribeIdentityPoolUsage",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeIdentityPoolUsage, input, output)
	output = &DescribeIdentityPoolUsageOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) DescribeIdentityPoolUsage(input *DescribeIdentityPoolUsageInput) (output *DescribeIdentityPoolUsageOutput, err error) {
	req, out := c.DescribeIdentityPoolUsageRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeIdentityPoolUsage *aws.Operation

// DescribeIdentityUsageRequest generates a request for the DescribeIdentityUsage operation.
func (c *CognitoSync) DescribeIdentityUsageRequest(input *DescribeIdentityUsageInput) (req *aws.Request, output *DescribeIdentityUsageOutput) {
	if opDescribeIdentityUsage == nil {
		opDescribeIdentityUsage = &aws.Operation{
			Name:       "DescribeIdentityUsage",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeIdentityUsage, input, output)
	output = &DescribeIdentityUsageOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) DescribeIdentityUsage(input *DescribeIdentityUsageInput) (output *DescribeIdentityUsageOutput, err error) {
	req, out := c.DescribeIdentityUsageRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeIdentityUsage *aws.Operation

// GetIdentityPoolConfigurationRequest generates a request for the GetIdentityPoolConfiguration operation.
func (c *CognitoSync) GetIdentityPoolConfigurationRequest(input *GetIdentityPoolConfigurationInput) (req *aws.Request, output *GetIdentityPoolConfigurationOutput) {
	if opGetIdentityPoolConfiguration == nil {
		opGetIdentityPoolConfiguration = &aws.Operation{
			Name:       "GetIdentityPoolConfiguration",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}/configuration",
		}
	}

	req = aws.NewRequest(c.Service, opGetIdentityPoolConfiguration, input, output)
	output = &GetIdentityPoolConfigurationOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) GetIdentityPoolConfiguration(input *GetIdentityPoolConfigurationInput) (output *GetIdentityPoolConfigurationOutput, err error) {
	req, out := c.GetIdentityPoolConfigurationRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetIdentityPoolConfiguration *aws.Operation

// ListDatasetsRequest generates a request for the ListDatasets operation.
func (c *CognitoSync) ListDatasetsRequest(input *ListDatasetsInput) (req *aws.Request, output *ListDatasetsOutput) {
	if opListDatasets == nil {
		opListDatasets = &aws.Operation{
			Name:       "ListDatasets",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets",
		}
	}

	req = aws.NewRequest(c.Service, opListDatasets, input, output)
	output = &ListDatasetsOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) ListDatasets(input *ListDatasetsInput) (output *ListDatasetsOutput, err error) {
	req, out := c.ListDatasetsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListDatasets *aws.Operation

// ListIdentityPoolUsageRequest generates a request for the ListIdentityPoolUsage operation.
func (c *CognitoSync) ListIdentityPoolUsageRequest(input *ListIdentityPoolUsageInput) (req *aws.Request, output *ListIdentityPoolUsageOutput) {
	if opListIdentityPoolUsage == nil {
		opListIdentityPoolUsage = &aws.Operation{
			Name:       "ListIdentityPoolUsage",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools",
		}
	}

	req = aws.NewRequest(c.Service, opListIdentityPoolUsage, input, output)
	output = &ListIdentityPoolUsageOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) ListIdentityPoolUsage(input *ListIdentityPoolUsageInput) (output *ListIdentityPoolUsageOutput, err error) {
	req, out := c.ListIdentityPoolUsageRequest(input)
	output = out
	err = req.Send()
	return
}

var opListIdentityPoolUsage *aws.Operation

// ListRecordsRequest generates a request for the ListRecords operation.
func (c *CognitoSync) ListRecordsRequest(input *ListRecordsInput) (req *aws.Request, output *ListRecordsOutput) {
	if opListRecords == nil {
		opListRecords = &aws.Operation{
			Name:       "ListRecords",
			HTTPMethod: "GET",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}/records",
		}
	}

	req = aws.NewRequest(c.Service, opListRecords, input, output)
	output = &ListRecordsOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) ListRecords(input *ListRecordsInput) (output *ListRecordsOutput, err error) {
	req, out := c.ListRecordsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListRecords *aws.Operation

// RegisterDeviceRequest generates a request for the RegisterDevice operation.
func (c *CognitoSync) RegisterDeviceRequest(input *RegisterDeviceInput) (req *aws.Request, output *RegisterDeviceOutput) {
	if opRegisterDevice == nil {
		opRegisterDevice = &aws.Operation{
			Name:       "RegisterDevice",
			HTTPMethod: "POST",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identity/{IdentityId}/device",
		}
	}

	req = aws.NewRequest(c.Service, opRegisterDevice, input, output)
	output = &RegisterDeviceOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) RegisterDevice(input *RegisterDeviceInput) (output *RegisterDeviceOutput, err error) {
	req, out := c.RegisterDeviceRequest(input)
	output = out
	err = req.Send()
	return
}

var opRegisterDevice *aws.Operation

// SetIdentityPoolConfigurationRequest generates a request for the SetIdentityPoolConfiguration operation.
func (c *CognitoSync) SetIdentityPoolConfigurationRequest(input *SetIdentityPoolConfigurationInput) (req *aws.Request, output *SetIdentityPoolConfigurationOutput) {
	if opSetIdentityPoolConfiguration == nil {
		opSetIdentityPoolConfiguration = &aws.Operation{
			Name:       "SetIdentityPoolConfiguration",
			HTTPMethod: "POST",
			HTTPPath:   "/identitypools/{IdentityPoolId}/configuration",
		}
	}

	req = aws.NewRequest(c.Service, opSetIdentityPoolConfiguration, input, output)
	output = &SetIdentityPoolConfigurationOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) SetIdentityPoolConfiguration(input *SetIdentityPoolConfigurationInput) (output *SetIdentityPoolConfigurationOutput, err error) {
	req, out := c.SetIdentityPoolConfigurationRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetIdentityPoolConfiguration *aws.Operation

// SubscribeToDatasetRequest generates a request for the SubscribeToDataset operation.
func (c *CognitoSync) SubscribeToDatasetRequest(input *SubscribeToDatasetInput) (req *aws.Request, output *SubscribeToDatasetOutput) {
	if opSubscribeToDataset == nil {
		opSubscribeToDataset = &aws.Operation{
			Name:       "SubscribeToDataset",
			HTTPMethod: "POST",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}/subscriptions/{DeviceId}",
		}
	}

	req = aws.NewRequest(c.Service, opSubscribeToDataset, input, output)
	output = &SubscribeToDatasetOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) SubscribeToDataset(input *SubscribeToDatasetInput) (output *SubscribeToDatasetOutput, err error) {
	req, out := c.SubscribeToDatasetRequest(input)
	output = out
	err = req.Send()
	return
}

var opSubscribeToDataset *aws.Operation

// UnsubscribeFromDatasetRequest generates a request for the UnsubscribeFromDataset operation.
func (c *CognitoSync) UnsubscribeFromDatasetRequest(input *UnsubscribeFromDatasetInput) (req *aws.Request, output *UnsubscribeFromDatasetOutput) {
	if opUnsubscribeFromDataset == nil {
		opUnsubscribeFromDataset = &aws.Operation{
			Name:       "UnsubscribeFromDataset",
			HTTPMethod: "DELETE",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}/subscriptions/{DeviceId}",
		}
	}

	req = aws.NewRequest(c.Service, opUnsubscribeFromDataset, input, output)
	output = &UnsubscribeFromDatasetOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) UnsubscribeFromDataset(input *UnsubscribeFromDatasetInput) (output *UnsubscribeFromDatasetOutput, err error) {
	req, out := c.UnsubscribeFromDatasetRequest(input)
	output = out
	err = req.Send()
	return
}

var opUnsubscribeFromDataset *aws.Operation

// UpdateRecordsRequest generates a request for the UpdateRecords operation.
func (c *CognitoSync) UpdateRecordsRequest(input *UpdateRecordsInput) (req *aws.Request, output *UpdateRecordsOutput) {
	if opUpdateRecords == nil {
		opUpdateRecords = &aws.Operation{
			Name:       "UpdateRecords",
			HTTPMethod: "POST",
			HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateRecords, input, output)
	output = &UpdateRecordsOutput{}
	req.Data = output
	return
}

func (c *CognitoSync) UpdateRecords(input *UpdateRecordsInput) (output *UpdateRecordsOutput, err error) {
	req, out := c.UpdateRecordsRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateRecords *aws.Operation

type Dataset struct {
	CreationDate     *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	DataStorage      *int64     `type:"long" json:",omitempty"`
	DatasetName      *string    `type:"string" json:",omitempty"`
	IdentityID       *string    `locationName:"IdentityId" type:"string" json:"IdentityId,omitempty"`
	LastModifiedBy   *string    `type:"string" json:",omitempty"`
	LastModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	NumRecords       *int64     `type:"long" json:",omitempty"`

	metadataDataset `json:"-", xml:"-"`
}

type metadataDataset struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DeleteDatasetInput struct {
	DatasetName    *string `location:"uri" locationName:"DatasetName" type:"string" required:"true"json:"-" xml:"-"`
	IdentityID     *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"json:"-" xml:"-"`
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"json:"-" xml:"-"`

	metadataDeleteDatasetInput `json:"-", xml:"-"`
}

type metadataDeleteDatasetInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DeleteDatasetOutput struct {
	Dataset *Dataset `type:"structure" json:",omitempty"`

	metadataDeleteDatasetOutput `json:"-", xml:"-"`
}

type metadataDeleteDatasetOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeDatasetInput struct {
	DatasetName    *string `location:"uri" locationName:"DatasetName" type:"string" required:"true"json:"-" xml:"-"`
	IdentityID     *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"json:"-" xml:"-"`
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"json:"-" xml:"-"`

	metadataDescribeDatasetInput `json:"-", xml:"-"`
}

type metadataDescribeDatasetInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeDatasetOutput struct {
	Dataset *Dataset `type:"structure" json:",omitempty"`

	metadataDescribeDatasetOutput `json:"-", xml:"-"`
}

type metadataDescribeDatasetOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeIdentityPoolUsageInput struct {
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"json:"-" xml:"-"`

	metadataDescribeIdentityPoolUsageInput `json:"-", xml:"-"`
}

type metadataDescribeIdentityPoolUsageInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeIdentityPoolUsageOutput struct {
	IdentityPoolUsage *IdentityPoolUsage `type:"structure" json:",omitempty"`

	metadataDescribeIdentityPoolUsageOutput `json:"-", xml:"-"`
}

type metadataDescribeIdentityPoolUsageOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeIdentityUsageInput struct {
	IdentityID     *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"json:"-" xml:"-"`
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"json:"-" xml:"-"`

	metadataDescribeIdentityUsageInput `json:"-", xml:"-"`
}

type metadataDescribeIdentityUsageInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeIdentityUsageOutput struct {
	IdentityUsage *IdentityUsage `type:"structure" json:",omitempty"`

	metadataDescribeIdentityUsageOutput `json:"-", xml:"-"`
}

type metadataDescribeIdentityUsageOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetIdentityPoolConfigurationInput struct {
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"json:"-" xml:"-"`

	metadataGetIdentityPoolConfigurationInput `json:"-", xml:"-"`
}

type metadataGetIdentityPoolConfigurationInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetIdentityPoolConfigurationOutput struct {
	IdentityPoolID *string   `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`
	PushSync       *PushSync `type:"structure" json:",omitempty"`

	metadataGetIdentityPoolConfigurationOutput `json:"-", xml:"-"`
}

type metadataGetIdentityPoolConfigurationOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type IdentityPoolUsage struct {
	DataStorage       *int64     `type:"long" json:",omitempty"`
	IdentityPoolID    *string    `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`
	LastModifiedDate  *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	SyncSessionsCount *int64     `type:"long" json:",omitempty"`

	metadataIdentityPoolUsage `json:"-", xml:"-"`
}

type metadataIdentityPoolUsage struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type IdentityUsage struct {
	DataStorage      *int64     `type:"long" json:",omitempty"`
	DatasetCount     *int64     `type:"integer" json:",omitempty"`
	IdentityID       *string    `locationName:"IdentityId" type:"string" json:"IdentityId,omitempty"`
	IdentityPoolID   *string    `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`
	LastModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`

	metadataIdentityUsage `json:"-", xml:"-"`
}

type metadataIdentityUsage struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListDatasetsInput struct {
	IdentityID     *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"json:"-" xml:"-"`
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"json:"-" xml:"-"`
	MaxResults     *int64  `location:"querystring" locationName:"maxResults" type:"integer" json:"-" xml:"-"`
	NextToken      *string `location:"querystring" locationName:"nextToken" type:"string" json:"-" xml:"-"`

	metadataListDatasetsInput `json:"-", xml:"-"`
}

type metadataListDatasetsInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListDatasetsOutput struct {
	Count     *int64     `type:"integer" json:",omitempty"`
	Datasets  []*Dataset `type:"list" json:",omitempty"`
	NextToken *string    `type:"string" json:",omitempty"`

	metadataListDatasetsOutput `json:"-", xml:"-"`
}

type metadataListDatasetsOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListIdentityPoolUsageInput struct {
	MaxResults *int64  `location:"querystring" locationName:"maxResults" type:"integer" json:"-" xml:"-"`
	NextToken  *string `location:"querystring" locationName:"nextToken" type:"string" json:"-" xml:"-"`

	metadataListIdentityPoolUsageInput `json:"-", xml:"-"`
}

type metadataListIdentityPoolUsageInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListIdentityPoolUsageOutput struct {
	Count              *int64               `type:"integer" json:",omitempty"`
	IdentityPoolUsages []*IdentityPoolUsage `type:"list" json:",omitempty"`
	MaxResults         *int64               `type:"integer" json:",omitempty"`
	NextToken          *string              `type:"string" json:",omitempty"`

	metadataListIdentityPoolUsageOutput `json:"-", xml:"-"`
}

type metadataListIdentityPoolUsageOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListRecordsInput struct {
	DatasetName      *string `location:"uri" locationName:"DatasetName" type:"string" required:"true"json:"-" xml:"-"`
	IdentityID       *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"json:"-" xml:"-"`
	IdentityPoolID   *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"json:"-" xml:"-"`
	LastSyncCount    *int64  `location:"querystring" locationName:"lastSyncCount" type:"long" json:"-" xml:"-"`
	MaxResults       *int64  `location:"querystring" locationName:"maxResults" type:"integer" json:"-" xml:"-"`
	NextToken        *string `location:"querystring" locationName:"nextToken" type:"string" json:"-" xml:"-"`
	SyncSessionToken *string `location:"querystring" locationName:"syncSessionToken" type:"string" json:"-" xml:"-"`

	metadataListRecordsInput `json:"-", xml:"-"`
}

type metadataListRecordsInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListRecordsOutput struct {
	Count                                 *int64    `type:"integer" json:",omitempty"`
	DatasetDeletedAfterRequestedSyncCount *bool     `type:"boolean" json:",omitempty"`
	DatasetExists                         *bool     `type:"boolean" json:",omitempty"`
	DatasetSyncCount                      *int64    `type:"long" json:",omitempty"`
	LastModifiedBy                        *string   `type:"string" json:",omitempty"`
	MergedDatasetNames                    []*string `type:"list" json:",omitempty"`
	NextToken                             *string   `type:"string" json:",omitempty"`
	Records                               []*Record `type:"list" json:",omitempty"`
	SyncSessionToken                      *string   `type:"string" json:",omitempty"`

	metadataListRecordsOutput `json:"-", xml:"-"`
}

type metadataListRecordsOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type PushSync struct {
	ApplicationARNs []*string `locationName:"ApplicationArns" type:"list" json:"ApplicationArns,omitempty"`
	RoleARN         *string   `locationName:"RoleArn" type:"string" json:"RoleArn,omitempty"`

	metadataPushSync `json:"-", xml:"-"`
}

type metadataPushSync struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type Record struct {
	DeviceLastModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	Key                    *string    `type:"string" json:",omitempty"`
	LastModifiedBy         *string    `type:"string" json:",omitempty"`
	LastModifiedDate       *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	SyncCount              *int64     `type:"long" json:",omitempty"`
	Value                  *string    `type:"string" json:",omitempty"`

	metadataRecord `json:"-", xml:"-"`
}

type metadataRecord struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type RecordPatch struct {
	DeviceLastModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	Key                    *string    `type:"string" required:"true"json:",omitempty"`
	Op                     *string    `type:"string" required:"true"json:",omitempty"`
	SyncCount              *int64     `type:"long" required:"true"json:",omitempty"`
	Value                  *string    `type:"string" json:",omitempty"`

	metadataRecordPatch `json:"-", xml:"-"`
}

type metadataRecordPatch struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type RegisterDeviceInput struct {
	IdentityID     *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"json:"-" xml:"-"`
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"json:"-" xml:"-"`
	Platform       *string `type:"string" required:"true"json:",omitempty"`
	Token          *string `type:"string" required:"true"json:",omitempty"`

	metadataRegisterDeviceInput `json:"-", xml:"-"`
}

type metadataRegisterDeviceInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type RegisterDeviceOutput struct {
	DeviceID *string `locationName:"DeviceId" type:"string" json:"DeviceId,omitempty"`

	metadataRegisterDeviceOutput `json:"-", xml:"-"`
}

type metadataRegisterDeviceOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type SetIdentityPoolConfigurationInput struct {
	IdentityPoolID *string   `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"json:"-" xml:"-"`
	PushSync       *PushSync `type:"structure" json:",omitempty"`

	metadataSetIdentityPoolConfigurationInput `json:"-", xml:"-"`
}

type metadataSetIdentityPoolConfigurationInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type SetIdentityPoolConfigurationOutput struct {
	IdentityPoolID *string   `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`
	PushSync       *PushSync `type:"structure" json:",omitempty"`

	metadataSetIdentityPoolConfigurationOutput `json:"-", xml:"-"`
}

type metadataSetIdentityPoolConfigurationOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type SubscribeToDatasetInput struct {
	DatasetName    *string `location:"uri" locationName:"DatasetName" type:"string" required:"true"json:"-" xml:"-"`
	DeviceID       *string `location:"uri" locationName:"DeviceId" type:"string" required:"true"json:"-" xml:"-"`
	IdentityID     *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"json:"-" xml:"-"`
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"json:"-" xml:"-"`

	metadataSubscribeToDatasetInput `json:"-", xml:"-"`
}

type metadataSubscribeToDatasetInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type SubscribeToDatasetOutput struct {
	metadataSubscribeToDatasetOutput `json:"-", xml:"-"`
}

type metadataSubscribeToDatasetOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type UnsubscribeFromDatasetInput struct {
	DatasetName    *string `location:"uri" locationName:"DatasetName" type:"string" required:"true"json:"-" xml:"-"`
	DeviceID       *string `location:"uri" locationName:"DeviceId" type:"string" required:"true"json:"-" xml:"-"`
	IdentityID     *string `location:"uri" locationName:"IdentityId" type:"string" required:"true"json:"-" xml:"-"`
	IdentityPoolID *string `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"json:"-" xml:"-"`

	metadataUnsubscribeFromDatasetInput `json:"-", xml:"-"`
}

type metadataUnsubscribeFromDatasetInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type UnsubscribeFromDatasetOutput struct {
	metadataUnsubscribeFromDatasetOutput `json:"-", xml:"-"`
}

type metadataUnsubscribeFromDatasetOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type UpdateRecordsInput struct {
	ClientContext    *string        `location:"header" locationName:"x-amz-Client-Context" type:"string" json:"-" xml:"-"`
	DatasetName      *string        `location:"uri" locationName:"DatasetName" type:"string" required:"true"json:"-" xml:"-"`
	DeviceID         *string        `locationName:"DeviceId" type:"string" json:"DeviceId,omitempty"`
	IdentityID       *string        `location:"uri" locationName:"IdentityId" type:"string" required:"true"json:"-" xml:"-"`
	IdentityPoolID   *string        `location:"uri" locationName:"IdentityPoolId" type:"string" required:"true"json:"-" xml:"-"`
	RecordPatches    []*RecordPatch `type:"list" json:",omitempty"`
	SyncSessionToken *string        `type:"string" required:"true"json:",omitempty"`

	metadataUpdateRecordsInput `json:"-", xml:"-"`
}

type metadataUpdateRecordsInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type UpdateRecordsOutput struct {
	Records []*Record `type:"list" json:",omitempty"`

	metadataUpdateRecordsOutput `json:"-", xml:"-"`
}

type metadataUpdateRecordsOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}