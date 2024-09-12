// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: allprotos/externalGateway.proto

package externalgatewaypb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ExternalGatewayService_OhlcvsWithCalculations_FullMethodName = "/externalgateway.ExternalGatewayService/OhlcvsWithCalculations"
	ExternalGatewayService_Ohlcvs_FullMethodName                 = "/externalgateway.ExternalGatewayService/Ohlcvs"
	ExternalGatewayService_OhlcvsStream_FullMethodName           = "/externalgateway.ExternalGatewayService/OhlcvsStream"
	ExternalGatewayService_Exchanges_FullMethodName              = "/externalgateway.ExternalGatewayService/Exchanges"
	ExternalGatewayService_Exchange_FullMethodName               = "/externalgateway.ExternalGatewayService/Exchange"
	ExternalGatewayService_OhlcvsMessage_FullMethodName          = "/externalgateway.ExternalGatewayService/OhlcvsMessage"
	ExternalGatewayService_Balance_FullMethodName                = "/externalgateway.ExternalGatewayService/Balance"
	ExternalGatewayService_FetchRealtimeValues_FullMethodName    = "/externalgateway.ExternalGatewayService/FetchRealtimeValues"
	ExternalGatewayService_ValidateAPIKey_FullMethodName         = "/externalgateway.ExternalGatewayService/ValidateAPIKey"
	ExternalGatewayService_SendTelegramMessage_FullMethodName    = "/externalgateway.ExternalGatewayService/SendTelegramMessage"
	ExternalGatewayService_ReportUnexpectedEvent_FullMethodName  = "/externalgateway.ExternalGatewayService/ReportUnexpectedEvent"
	ExternalGatewayService_SendNotification_FullMethodName       = "/externalgateway.ExternalGatewayService/SendNotification"
	ExternalGatewayService_ImportUSDTPairs_FullMethodName        = "/externalgateway.ExternalGatewayService/ImportUSDTPairs"
	ExternalGatewayService_CreatePayment_FullMethodName          = "/externalgateway.ExternalGatewayService/CreatePayment"
	ExternalGatewayService_VerifyPayment_FullMethodName          = "/externalgateway.ExternalGatewayService/VerifyPayment"
	ExternalGatewayService_Markets_FullMethodName                = "/externalgateway.ExternalGatewayService/Markets"
)

// ExternalGatewayServiceClient is the client API for ExternalGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExternalGatewayServiceClient interface {
	OhlcvsWithCalculations(ctx context.Context, in *OhlcvWithCalculationsRequest, opts ...grpc.CallOption) (*OhlcvsWithCalculationsResponse, error)
	Ohlcvs(ctx context.Context, in *OhlcvsRequest, opts ...grpc.CallOption) (*OhlcvsResponse, error)
	OhlcvsStream(ctx context.Context, in *OhlcvsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OhlcvsResponse], error)
	Exchanges(ctx context.Context, in *ExchangesRequest, opts ...grpc.CallOption) (*ExchangesResponse, error)
	Exchange(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*ExchangeResponse, error)
	OhlcvsMessage(ctx context.Context, in *OhlcvsMessageRequest, opts ...grpc.CallOption) (*OhlcvsMessageResponse, error)
	Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
	FetchRealtimeValues(ctx context.Context, in *FetchRealtimeValuesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[FetchRealtimeValuesResponse], error)
	ValidateAPIKey(ctx context.Context, in *ValidateAPIKeyRequest, opts ...grpc.CallOption) (*ValidateAPIKeyResponse, error)
	SendTelegramMessage(ctx context.Context, in *SendTelegramMessageRequest, opts ...grpc.CallOption) (*SendTelegramMessageResponse, error)
	ReportUnexpectedEvent(ctx context.Context, in *ReportUnexpectedEventRequest, opts ...grpc.CallOption) (*ReportUnexpectedEventResponse, error)
	SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*SendNotificationResponse, error)
	ImportUSDTPairs(ctx context.Context, in *ImportUSDTPairsRequest, opts ...grpc.CallOption) (*ImportUSDTPairsResponse, error)
	CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
	VerifyPayment(ctx context.Context, in *VerifyPaymentRequest, opts ...grpc.CallOption) (*VerifyPaymentResponse, error)
	Markets(ctx context.Context, in *MarketsRequest, opts ...grpc.CallOption) (*MarketsResponse, error)
}

type externalGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalGatewayServiceClient(cc grpc.ClientConnInterface) ExternalGatewayServiceClient {
	return &externalGatewayServiceClient{cc}
}

func (c *externalGatewayServiceClient) OhlcvsWithCalculations(ctx context.Context, in *OhlcvWithCalculationsRequest, opts ...grpc.CallOption) (*OhlcvsWithCalculationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OhlcvsWithCalculationsResponse)
	err := c.cc.Invoke(ctx, ExternalGatewayService_OhlcvsWithCalculations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalGatewayServiceClient) Ohlcvs(ctx context.Context, in *OhlcvsRequest, opts ...grpc.CallOption) (*OhlcvsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OhlcvsResponse)
	err := c.cc.Invoke(ctx, ExternalGatewayService_Ohlcvs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalGatewayServiceClient) OhlcvsStream(ctx context.Context, in *OhlcvsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OhlcvsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExternalGatewayService_ServiceDesc.Streams[0], ExternalGatewayService_OhlcvsStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[OhlcvsRequest, OhlcvsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExternalGatewayService_OhlcvsStreamClient = grpc.ServerStreamingClient[OhlcvsResponse]

func (c *externalGatewayServiceClient) Exchanges(ctx context.Context, in *ExchangesRequest, opts ...grpc.CallOption) (*ExchangesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExchangesResponse)
	err := c.cc.Invoke(ctx, ExternalGatewayService_Exchanges_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalGatewayServiceClient) Exchange(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*ExchangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExchangeResponse)
	err := c.cc.Invoke(ctx, ExternalGatewayService_Exchange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalGatewayServiceClient) OhlcvsMessage(ctx context.Context, in *OhlcvsMessageRequest, opts ...grpc.CallOption) (*OhlcvsMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OhlcvsMessageResponse)
	err := c.cc.Invoke(ctx, ExternalGatewayService_OhlcvsMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalGatewayServiceClient) Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, ExternalGatewayService_Balance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalGatewayServiceClient) FetchRealtimeValues(ctx context.Context, in *FetchRealtimeValuesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[FetchRealtimeValuesResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExternalGatewayService_ServiceDesc.Streams[1], ExternalGatewayService_FetchRealtimeValues_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FetchRealtimeValuesRequest, FetchRealtimeValuesResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExternalGatewayService_FetchRealtimeValuesClient = grpc.ServerStreamingClient[FetchRealtimeValuesResponse]

func (c *externalGatewayServiceClient) ValidateAPIKey(ctx context.Context, in *ValidateAPIKeyRequest, opts ...grpc.CallOption) (*ValidateAPIKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateAPIKeyResponse)
	err := c.cc.Invoke(ctx, ExternalGatewayService_ValidateAPIKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalGatewayServiceClient) SendTelegramMessage(ctx context.Context, in *SendTelegramMessageRequest, opts ...grpc.CallOption) (*SendTelegramMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendTelegramMessageResponse)
	err := c.cc.Invoke(ctx, ExternalGatewayService_SendTelegramMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalGatewayServiceClient) ReportUnexpectedEvent(ctx context.Context, in *ReportUnexpectedEventRequest, opts ...grpc.CallOption) (*ReportUnexpectedEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReportUnexpectedEventResponse)
	err := c.cc.Invoke(ctx, ExternalGatewayService_ReportUnexpectedEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalGatewayServiceClient) SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*SendNotificationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendNotificationResponse)
	err := c.cc.Invoke(ctx, ExternalGatewayService_SendNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalGatewayServiceClient) ImportUSDTPairs(ctx context.Context, in *ImportUSDTPairsRequest, opts ...grpc.CallOption) (*ImportUSDTPairsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ImportUSDTPairsResponse)
	err := c.cc.Invoke(ctx, ExternalGatewayService_ImportUSDTPairs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalGatewayServiceClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePaymentResponse)
	err := c.cc.Invoke(ctx, ExternalGatewayService_CreatePayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalGatewayServiceClient) VerifyPayment(ctx context.Context, in *VerifyPaymentRequest, opts ...grpc.CallOption) (*VerifyPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyPaymentResponse)
	err := c.cc.Invoke(ctx, ExternalGatewayService_VerifyPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalGatewayServiceClient) Markets(ctx context.Context, in *MarketsRequest, opts ...grpc.CallOption) (*MarketsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MarketsResponse)
	err := c.cc.Invoke(ctx, ExternalGatewayService_Markets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalGatewayServiceServer is the server API for ExternalGatewayService service.
// All implementations must embed UnimplementedExternalGatewayServiceServer
// for forward compatibility.
type ExternalGatewayServiceServer interface {
	OhlcvsWithCalculations(context.Context, *OhlcvWithCalculationsRequest) (*OhlcvsWithCalculationsResponse, error)
	Ohlcvs(context.Context, *OhlcvsRequest) (*OhlcvsResponse, error)
	OhlcvsStream(*OhlcvsRequest, grpc.ServerStreamingServer[OhlcvsResponse]) error
	Exchanges(context.Context, *ExchangesRequest) (*ExchangesResponse, error)
	Exchange(context.Context, *ExchangeRequest) (*ExchangeResponse, error)
	OhlcvsMessage(context.Context, *OhlcvsMessageRequest) (*OhlcvsMessageResponse, error)
	Balance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	FetchRealtimeValues(*FetchRealtimeValuesRequest, grpc.ServerStreamingServer[FetchRealtimeValuesResponse]) error
	ValidateAPIKey(context.Context, *ValidateAPIKeyRequest) (*ValidateAPIKeyResponse, error)
	SendTelegramMessage(context.Context, *SendTelegramMessageRequest) (*SendTelegramMessageResponse, error)
	ReportUnexpectedEvent(context.Context, *ReportUnexpectedEventRequest) (*ReportUnexpectedEventResponse, error)
	SendNotification(context.Context, *SendNotificationRequest) (*SendNotificationResponse, error)
	ImportUSDTPairs(context.Context, *ImportUSDTPairsRequest) (*ImportUSDTPairsResponse, error)
	CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error)
	VerifyPayment(context.Context, *VerifyPaymentRequest) (*VerifyPaymentResponse, error)
	Markets(context.Context, *MarketsRequest) (*MarketsResponse, error)
	mustEmbedUnimplementedExternalGatewayServiceServer()
}

// UnimplementedExternalGatewayServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExternalGatewayServiceServer struct{}

func (UnimplementedExternalGatewayServiceServer) OhlcvsWithCalculations(context.Context, *OhlcvWithCalculationsRequest) (*OhlcvsWithCalculationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OhlcvsWithCalculations not implemented")
}
func (UnimplementedExternalGatewayServiceServer) Ohlcvs(context.Context, *OhlcvsRequest) (*OhlcvsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ohlcvs not implemented")
}
func (UnimplementedExternalGatewayServiceServer) OhlcvsStream(*OhlcvsRequest, grpc.ServerStreamingServer[OhlcvsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method OhlcvsStream not implemented")
}
func (UnimplementedExternalGatewayServiceServer) Exchanges(context.Context, *ExchangesRequest) (*ExchangesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exchanges not implemented")
}
func (UnimplementedExternalGatewayServiceServer) Exchange(context.Context, *ExchangeRequest) (*ExchangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exchange not implemented")
}
func (UnimplementedExternalGatewayServiceServer) OhlcvsMessage(context.Context, *OhlcvsMessageRequest) (*OhlcvsMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OhlcvsMessage not implemented")
}
func (UnimplementedExternalGatewayServiceServer) Balance(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedExternalGatewayServiceServer) FetchRealtimeValues(*FetchRealtimeValuesRequest, grpc.ServerStreamingServer[FetchRealtimeValuesResponse]) error {
	return status.Errorf(codes.Unimplemented, "method FetchRealtimeValues not implemented")
}
func (UnimplementedExternalGatewayServiceServer) ValidateAPIKey(context.Context, *ValidateAPIKeyRequest) (*ValidateAPIKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAPIKey not implemented")
}
func (UnimplementedExternalGatewayServiceServer) SendTelegramMessage(context.Context, *SendTelegramMessageRequest) (*SendTelegramMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTelegramMessage not implemented")
}
func (UnimplementedExternalGatewayServiceServer) ReportUnexpectedEvent(context.Context, *ReportUnexpectedEventRequest) (*ReportUnexpectedEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportUnexpectedEvent not implemented")
}
func (UnimplementedExternalGatewayServiceServer) SendNotification(context.Context, *SendNotificationRequest) (*SendNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotification not implemented")
}
func (UnimplementedExternalGatewayServiceServer) ImportUSDTPairs(context.Context, *ImportUSDTPairsRequest) (*ImportUSDTPairsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportUSDTPairs not implemented")
}
func (UnimplementedExternalGatewayServiceServer) CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedExternalGatewayServiceServer) VerifyPayment(context.Context, *VerifyPaymentRequest) (*VerifyPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPayment not implemented")
}
func (UnimplementedExternalGatewayServiceServer) Markets(context.Context, *MarketsRequest) (*MarketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Markets not implemented")
}
func (UnimplementedExternalGatewayServiceServer) mustEmbedUnimplementedExternalGatewayServiceServer() {
}
func (UnimplementedExternalGatewayServiceServer) testEmbeddedByValue() {}

// UnsafeExternalGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExternalGatewayServiceServer will
// result in compilation errors.
type UnsafeExternalGatewayServiceServer interface {
	mustEmbedUnimplementedExternalGatewayServiceServer()
}

func RegisterExternalGatewayServiceServer(s grpc.ServiceRegistrar, srv ExternalGatewayServiceServer) {
	// If the following call pancis, it indicates UnimplementedExternalGatewayServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExternalGatewayService_ServiceDesc, srv)
}

func _ExternalGatewayService_OhlcvsWithCalculations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OhlcvWithCalculationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalGatewayServiceServer).OhlcvsWithCalculations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalGatewayService_OhlcvsWithCalculations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalGatewayServiceServer).OhlcvsWithCalculations(ctx, req.(*OhlcvWithCalculationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalGatewayService_Ohlcvs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OhlcvsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalGatewayServiceServer).Ohlcvs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalGatewayService_Ohlcvs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalGatewayServiceServer).Ohlcvs(ctx, req.(*OhlcvsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalGatewayService_OhlcvsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OhlcvsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExternalGatewayServiceServer).OhlcvsStream(m, &grpc.GenericServerStream[OhlcvsRequest, OhlcvsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExternalGatewayService_OhlcvsStreamServer = grpc.ServerStreamingServer[OhlcvsResponse]

func _ExternalGatewayService_Exchanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalGatewayServiceServer).Exchanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalGatewayService_Exchanges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalGatewayServiceServer).Exchanges(ctx, req.(*ExchangesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalGatewayService_Exchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalGatewayServiceServer).Exchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalGatewayService_Exchange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalGatewayServiceServer).Exchange(ctx, req.(*ExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalGatewayService_OhlcvsMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OhlcvsMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalGatewayServiceServer).OhlcvsMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalGatewayService_OhlcvsMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalGatewayServiceServer).OhlcvsMessage(ctx, req.(*OhlcvsMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalGatewayService_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalGatewayServiceServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalGatewayService_Balance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalGatewayServiceServer).Balance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalGatewayService_FetchRealtimeValues_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FetchRealtimeValuesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExternalGatewayServiceServer).FetchRealtimeValues(m, &grpc.GenericServerStream[FetchRealtimeValuesRequest, FetchRealtimeValuesResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExternalGatewayService_FetchRealtimeValuesServer = grpc.ServerStreamingServer[FetchRealtimeValuesResponse]

func _ExternalGatewayService_ValidateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalGatewayServiceServer).ValidateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalGatewayService_ValidateAPIKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalGatewayServiceServer).ValidateAPIKey(ctx, req.(*ValidateAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalGatewayService_SendTelegramMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTelegramMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalGatewayServiceServer).SendTelegramMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalGatewayService_SendTelegramMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalGatewayServiceServer).SendTelegramMessage(ctx, req.(*SendTelegramMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalGatewayService_ReportUnexpectedEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportUnexpectedEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalGatewayServiceServer).ReportUnexpectedEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalGatewayService_ReportUnexpectedEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalGatewayServiceServer).ReportUnexpectedEvent(ctx, req.(*ReportUnexpectedEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalGatewayService_SendNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalGatewayServiceServer).SendNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalGatewayService_SendNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalGatewayServiceServer).SendNotification(ctx, req.(*SendNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalGatewayService_ImportUSDTPairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportUSDTPairsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalGatewayServiceServer).ImportUSDTPairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalGatewayService_ImportUSDTPairs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalGatewayServiceServer).ImportUSDTPairs(ctx, req.(*ImportUSDTPairsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalGatewayService_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalGatewayServiceServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalGatewayService_CreatePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalGatewayServiceServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalGatewayService_VerifyPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalGatewayServiceServer).VerifyPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalGatewayService_VerifyPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalGatewayServiceServer).VerifyPayment(ctx, req.(*VerifyPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalGatewayService_Markets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalGatewayServiceServer).Markets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalGatewayService_Markets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalGatewayServiceServer).Markets(ctx, req.(*MarketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExternalGatewayService_ServiceDesc is the grpc.ServiceDesc for ExternalGatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExternalGatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "externalgateway.ExternalGatewayService",
	HandlerType: (*ExternalGatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OhlcvsWithCalculations",
			Handler:    _ExternalGatewayService_OhlcvsWithCalculations_Handler,
		},
		{
			MethodName: "Ohlcvs",
			Handler:    _ExternalGatewayService_Ohlcvs_Handler,
		},
		{
			MethodName: "Exchanges",
			Handler:    _ExternalGatewayService_Exchanges_Handler,
		},
		{
			MethodName: "Exchange",
			Handler:    _ExternalGatewayService_Exchange_Handler,
		},
		{
			MethodName: "OhlcvsMessage",
			Handler:    _ExternalGatewayService_OhlcvsMessage_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _ExternalGatewayService_Balance_Handler,
		},
		{
			MethodName: "ValidateAPIKey",
			Handler:    _ExternalGatewayService_ValidateAPIKey_Handler,
		},
		{
			MethodName: "SendTelegramMessage",
			Handler:    _ExternalGatewayService_SendTelegramMessage_Handler,
		},
		{
			MethodName: "ReportUnexpectedEvent",
			Handler:    _ExternalGatewayService_ReportUnexpectedEvent_Handler,
		},
		{
			MethodName: "SendNotification",
			Handler:    _ExternalGatewayService_SendNotification_Handler,
		},
		{
			MethodName: "ImportUSDTPairs",
			Handler:    _ExternalGatewayService_ImportUSDTPairs_Handler,
		},
		{
			MethodName: "CreatePayment",
			Handler:    _ExternalGatewayService_CreatePayment_Handler,
		},
		{
			MethodName: "VerifyPayment",
			Handler:    _ExternalGatewayService_VerifyPayment_Handler,
		},
		{
			MethodName: "Markets",
			Handler:    _ExternalGatewayService_Markets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OhlcvsStream",
			Handler:       _ExternalGatewayService_OhlcvsStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FetchRealtimeValues",
			Handler:       _ExternalGatewayService_FetchRealtimeValues_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "allprotos/externalGateway.proto",
}
