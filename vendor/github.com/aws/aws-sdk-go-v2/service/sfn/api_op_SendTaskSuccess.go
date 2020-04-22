// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfn

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SendTaskSuccessInput struct {
	_ struct{} `type:"structure"`

	// The JSON output of the task.
	//
	// Output is a required field
	Output *string `locationName:"output" type:"string" required:"true" sensitive:"true"`

	// The token that represents this task. Task tokens are generated by Step Functions
	// when tasks are assigned to a worker, or in the context object (https://docs.aws.amazon.com/step-functions/latest/dg/input-output-contextobject.html)
	// when a workflow enters a task state. See GetActivityTaskOutput$taskToken.
	//
	// TaskToken is a required field
	TaskToken *string `locationName:"taskToken" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s SendTaskSuccessInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendTaskSuccessInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SendTaskSuccessInput"}

	if s.Output == nil {
		invalidParams.Add(aws.NewErrParamRequired("Output"))
	}

	if s.TaskToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskToken"))
	}
	if s.TaskToken != nil && len(*s.TaskToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TaskToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SendTaskSuccessOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SendTaskSuccessOutput) String() string {
	return awsutil.Prettify(s)
}

const opSendTaskSuccess = "SendTaskSuccess"

// SendTaskSuccessRequest returns a request value for making API operation for
// AWS Step Functions.
//
// Used by activity workers and task states using the callback (https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token)
// pattern to report that the task identified by the taskToken completed successfully.
//
//    // Example sending a request using SendTaskSuccessRequest.
//    req := client.SendTaskSuccessRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/SendTaskSuccess
func (c *Client) SendTaskSuccessRequest(input *SendTaskSuccessInput) SendTaskSuccessRequest {
	op := &aws.Operation{
		Name:       opSendTaskSuccess,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SendTaskSuccessInput{}
	}

	req := c.newRequest(op, input, &SendTaskSuccessOutput{})
	return SendTaskSuccessRequest{Request: req, Input: input, Copy: c.SendTaskSuccessRequest}
}

// SendTaskSuccessRequest is the request type for the
// SendTaskSuccess API operation.
type SendTaskSuccessRequest struct {
	*aws.Request
	Input *SendTaskSuccessInput
	Copy  func(*SendTaskSuccessInput) SendTaskSuccessRequest
}

// Send marshals and sends the SendTaskSuccess API request.
func (r SendTaskSuccessRequest) Send(ctx context.Context) (*SendTaskSuccessResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SendTaskSuccessResponse{
		SendTaskSuccessOutput: r.Request.Data.(*SendTaskSuccessOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SendTaskSuccessResponse is the response type for the
// SendTaskSuccess API operation.
type SendTaskSuccessResponse struct {
	*SendTaskSuccessOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SendTaskSuccess request.
func (r *SendTaskSuccessResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}