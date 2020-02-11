// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateRoleRequest
type UpdateRoleInput struct {
	_ struct{} `type:"structure"`

	// The new description that you want to apply to the specified role.
	Description *string `type:"string"`

	// The maximum session duration (in seconds) that you want to set for the specified
	// role. If you do not specify a value for this setting, the default maximum
	// of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	//
	// Anyone who assumes the role from the AWS CLI or API can use the DurationSeconds
	// API parameter or the duration-seconds CLI parameter to request a longer session.
	// The MaxSessionDuration setting determines the maximum duration that can be
	// requested using the DurationSeconds parameter. If users don't specify a value
	// for the DurationSeconds parameter, their security credentials are valid for
	// one hour by default. This applies when you use the AssumeRole* API operations
	// or the assume-role* CLI operations but does not apply when you use those
	// operations to create a console URL. For more information, see Using IAM Roles
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html) in the
	// IAM User Guide.
	MaxSessionDuration *int64 `min:"3600" type:"integer"`

	// The name of the role that you want to modify.
	//
	// RoleName is a required field
	RoleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateRoleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRoleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRoleInput"}
	if s.MaxSessionDuration != nil && *s.MaxSessionDuration < 3600 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxSessionDuration", 3600))
	}

	if s.RoleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleName"))
	}
	if s.RoleName != nil && len(*s.RoleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateRoleResponse
type UpdateRoleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateRoleOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateRole = "UpdateRole"

// UpdateRoleRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Updates the description or maximum session duration setting of a role.
//
//    // Example sending a request using UpdateRoleRequest.
//    req := client.UpdateRoleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateRole
func (c *Client) UpdateRoleRequest(input *UpdateRoleInput) UpdateRoleRequest {
	op := &aws.Operation{
		Name:       opUpdateRole,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRoleInput{}
	}

	req := c.newRequest(op, input, &UpdateRoleOutput{})
	return UpdateRoleRequest{Request: req, Input: input, Copy: c.UpdateRoleRequest}
}

// UpdateRoleRequest is the request type for the
// UpdateRole API operation.
type UpdateRoleRequest struct {
	*aws.Request
	Input *UpdateRoleInput
	Copy  func(*UpdateRoleInput) UpdateRoleRequest
}

// Send marshals and sends the UpdateRole API request.
func (r UpdateRoleRequest) Send(ctx context.Context) (*UpdateRoleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRoleResponse{
		UpdateRoleOutput: r.Request.Data.(*UpdateRoleOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRoleResponse is the response type for the
// UpdateRole API operation.
type UpdateRoleResponse struct {
	*UpdateRoleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRole request.
func (r *UpdateRoleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
