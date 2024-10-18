// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Unassigns secondary private IPv4 addresses from a private NAT gateway. You
// cannot unassign your primary private IP. For more information, see [Edit secondary IP address associations]in the
// Amazon VPC User Guide.
//
// While unassigning is in progress, you cannot assign/unassign additional IP
// addresses while the connections are being drained. You are, however, allowed to
// delete the NAT gateway.
//
// A private IP address will only be released at the end of
// MaxDrainDurationSeconds. The private IP addresses stay associated and support
// the existing connections, but do not support any new connections (new
// connections are distributed across the remaining assigned private IP address).
// After the existing connections drain out, the private IP addresses are released.
//
// [Edit secondary IP address associations]: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-edit-secondary
func (c *Client) UnassignPrivateNatGatewayAddress(ctx context.Context, params *UnassignPrivateNatGatewayAddressInput, optFns ...func(*Options)) (*UnassignPrivateNatGatewayAddressOutput, error) {
	if params == nil {
		params = &UnassignPrivateNatGatewayAddressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UnassignPrivateNatGatewayAddress", params, optFns, c.addOperationUnassignPrivateNatGatewayAddressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UnassignPrivateNatGatewayAddressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UnassignPrivateNatGatewayAddressInput struct {

	// The ID of the NAT gateway.
	//
	// This member is required.
	NatGatewayId *string

	// The private IPv4 addresses you want to unassign.
	//
	// This member is required.
	PrivateIpAddresses []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The maximum amount of time to wait (in seconds) before forcibly releasing the
	// IP addresses if connections are still in progress. Default value is 350 seconds.
	MaxDrainDurationSeconds *int32

	noSmithyDocumentSerde
}

type UnassignPrivateNatGatewayAddressOutput struct {

	// Information about the NAT gateway IP addresses.
	NatGatewayAddresses []types.NatGatewayAddress

	// The ID of the NAT gateway.
	NatGatewayId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUnassignPrivateNatGatewayAddressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpUnassignPrivateNatGatewayAddress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpUnassignPrivateNatGatewayAddress{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UnassignPrivateNatGatewayAddress"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpUnassignPrivateNatGatewayAddressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUnassignPrivateNatGatewayAddress(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUnassignPrivateNatGatewayAddress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UnassignPrivateNatGatewayAddress",
	}
}