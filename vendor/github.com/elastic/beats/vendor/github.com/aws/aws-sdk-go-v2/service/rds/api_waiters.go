// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilDBInstanceAvailable uses the Amazon RDS API operation
// DescribeDBInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBInstanceAvailable(ctx context.Context, input *DescribeDBInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBInstanceAvailable",
		MaxAttempts: 60,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "available",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "deleted",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "deleting",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "failed",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "incompatible-restore",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "incompatible-parameters",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBInstanceDeleted uses the Amazon RDS API operation
// DescribeDBInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBInstanceDeleted(ctx context.Context, input *DescribeDBInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBInstanceDeleted",
		MaxAttempts: 60,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "deleted",
			},
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "DBInstanceNotFound",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "creating",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "modifying",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "rebooting",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "resetting-master-credentials",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBSnapshotAvailable uses the Amazon RDS API operation
// DescribeDBSnapshots to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBSnapshotAvailable(ctx context.Context, input *DescribeDBSnapshotsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBSnapshotAvailable",
		MaxAttempts: 60,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "available",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "deleted",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "deleting",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "failed",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "incompatible-restore",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "incompatible-parameters",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBSnapshotsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBSnapshotsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBSnapshotDeleted uses the Amazon RDS API operation
// DescribeDBSnapshots to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBSnapshotDeleted(ctx context.Context, input *DescribeDBSnapshotsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBSnapshotDeleted",
		MaxAttempts: 60,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "deleted",
			},
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "DBSnapshotNotFound",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "creating",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "modifying",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "rebooting",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "DBSnapshots[].Status",
				Expected: "resetting-master-credentials",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBSnapshotsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBSnapshotsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
