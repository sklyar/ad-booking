// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: booking/contact_person.proto

package bookingconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	booking "github.com/sklyar/ad-booking/backend/api/gen/booking"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ContactPersonServiceName is the fully-qualified name of the ContactPersonService service.
	ContactPersonServiceName = "booking.ContactPersonService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ContactPersonServiceCreateProcedure is the fully-qualified name of the ContactPersonService's
	// Create RPC.
	ContactPersonServiceCreateProcedure = "/booking.ContactPersonService/Create"
	// ContactPersonServiceListProcedure is the fully-qualified name of the ContactPersonService's List
	// RPC.
	ContactPersonServiceListProcedure = "/booking.ContactPersonService/List"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	contactPersonServiceServiceDescriptor      = booking.File_booking_contact_person_proto.Services().ByName("ContactPersonService")
	contactPersonServiceCreateMethodDescriptor = contactPersonServiceServiceDescriptor.Methods().ByName("Create")
	contactPersonServiceListMethodDescriptor   = contactPersonServiceServiceDescriptor.Methods().ByName("List")
)

// ContactPersonServiceClient is a client for the booking.ContactPersonService service.
type ContactPersonServiceClient interface {
	Create(context.Context, *connect.Request[booking.CreatePersonRequest]) (*connect.Response[booking.CreatePersonResponse], error)
	List(context.Context, *connect.Request[booking.ListPersonRequest]) (*connect.Response[booking.ListPersonResponse], error)
}

// NewContactPersonServiceClient constructs a client for the booking.ContactPersonService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewContactPersonServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ContactPersonServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &contactPersonServiceClient{
		create: connect.NewClient[booking.CreatePersonRequest, booking.CreatePersonResponse](
			httpClient,
			baseURL+ContactPersonServiceCreateProcedure,
			connect.WithSchema(contactPersonServiceCreateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		list: connect.NewClient[booking.ListPersonRequest, booking.ListPersonResponse](
			httpClient,
			baseURL+ContactPersonServiceListProcedure,
			connect.WithSchema(contactPersonServiceListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// contactPersonServiceClient implements ContactPersonServiceClient.
type contactPersonServiceClient struct {
	create *connect.Client[booking.CreatePersonRequest, booking.CreatePersonResponse]
	list   *connect.Client[booking.ListPersonRequest, booking.ListPersonResponse]
}

// Create calls booking.ContactPersonService.Create.
func (c *contactPersonServiceClient) Create(ctx context.Context, req *connect.Request[booking.CreatePersonRequest]) (*connect.Response[booking.CreatePersonResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// List calls booking.ContactPersonService.List.
func (c *contactPersonServiceClient) List(ctx context.Context, req *connect.Request[booking.ListPersonRequest]) (*connect.Response[booking.ListPersonResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// ContactPersonServiceHandler is an implementation of the booking.ContactPersonService service.
type ContactPersonServiceHandler interface {
	Create(context.Context, *connect.Request[booking.CreatePersonRequest]) (*connect.Response[booking.CreatePersonResponse], error)
	List(context.Context, *connect.Request[booking.ListPersonRequest]) (*connect.Response[booking.ListPersonResponse], error)
}

// NewContactPersonServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewContactPersonServiceHandler(svc ContactPersonServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	contactPersonServiceCreateHandler := connect.NewUnaryHandler(
		ContactPersonServiceCreateProcedure,
		svc.Create,
		connect.WithSchema(contactPersonServiceCreateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	contactPersonServiceListHandler := connect.NewUnaryHandler(
		ContactPersonServiceListProcedure,
		svc.List,
		connect.WithSchema(contactPersonServiceListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/booking.ContactPersonService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ContactPersonServiceCreateProcedure:
			contactPersonServiceCreateHandler.ServeHTTP(w, r)
		case ContactPersonServiceListProcedure:
			contactPersonServiceListHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedContactPersonServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedContactPersonServiceHandler struct{}

func (UnimplementedContactPersonServiceHandler) Create(context.Context, *connect.Request[booking.CreatePersonRequest]) (*connect.Response[booking.CreatePersonResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("booking.ContactPersonService.Create is not implemented"))
}

func (UnimplementedContactPersonServiceHandler) List(context.Context, *connect.Request[booking.ListPersonRequest]) (*connect.Response[booking.ListPersonResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("booking.ContactPersonService.List is not implemented"))
}
