// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: contactperson/contactperson.proto

package contactpersonconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	contactperson "github.com/sklyar/ad-booking/backend/api/gen/contactperson"
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
	// ServiceName is the fully-qualified name of the Service service.
	ServiceName = "contactperson.Service"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ServiceCreateProcedure is the fully-qualified name of the Service's Create RPC.
	ServiceCreateProcedure = "/contactperson.Service/Create"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	serviceServiceDescriptor      = contactperson.File_contactperson_contactperson_proto.Services().ByName("Service")
	serviceCreateMethodDescriptor = serviceServiceDescriptor.Methods().ByName("Create")
)

// ServiceClient is a client for the contactperson.Service service.
type ServiceClient interface {
	Create(context.Context, *connect.Request[contactperson.CreateRequest]) (*connect.Response[contactperson.CreateResponse], error)
}

// NewServiceClient constructs a client for the contactperson.Service service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &serviceClient{
		create: connect.NewClient[contactperson.CreateRequest, contactperson.CreateResponse](
			httpClient,
			baseURL+ServiceCreateProcedure,
			connect.WithSchema(serviceCreateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// serviceClient implements ServiceClient.
type serviceClient struct {
	create *connect.Client[contactperson.CreateRequest, contactperson.CreateResponse]
}

// Create calls contactperson.Service.Create.
func (c *serviceClient) Create(ctx context.Context, req *connect.Request[contactperson.CreateRequest]) (*connect.Response[contactperson.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// ServiceHandler is an implementation of the contactperson.Service service.
type ServiceHandler interface {
	Create(context.Context, *connect.Request[contactperson.CreateRequest]) (*connect.Response[contactperson.CreateResponse], error)
}

// NewServiceHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewServiceHandler(svc ServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	serviceCreateHandler := connect.NewUnaryHandler(
		ServiceCreateProcedure,
		svc.Create,
		connect.WithSchema(serviceCreateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/contactperson.Service/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ServiceCreateProcedure:
			serviceCreateHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedServiceHandler struct{}

func (UnimplementedServiceHandler) Create(context.Context, *connect.Request[contactperson.CreateRequest]) (*connect.Response[contactperson.CreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("contactperson.Service.Create is not implemented"))
}