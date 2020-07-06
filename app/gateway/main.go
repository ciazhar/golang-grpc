package main

import (
	"flag"
	util "github.com/ciazhar/golang-grpc/common/gateway"
	"github.com/ciazhar/golang-grpc/grpc/generated/golang"
	"google.golang.org/grpc"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
)

var (
	serverEndpoint = flag.String("/server", "localhost:50051", "endpoint of YourService")
)

func newGateway(ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	if err := golang.RegisterRecipesServiceHandlerFromEndpoint(ctx, mux, *serverEndpoint, dialOpts); err != nil {
		return nil, err
	}
	if err := golang.RegisterUserServiceHandlerFromEndpoint(ctx, mux, *serverEndpoint, dialOpts); err != nil {
		return nil, err
	}

	return mux, nil
}

func Run(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	gw, err := newGateway(ctx, opts...)
	if err != nil {
		return err
	}
	mux.Handle("/", gw)
	mux.Handle("/swagger/", http.StripPrefix("/swagger", http.FileServer(http.Dir("grpc/generated/swagger/grpc/proto"))))

	return http.ListenAndServe(address, util.CheckAuth(util.AllowCORS(mux)))

}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := Run(":8080"); err != nil {
		glog.Fatal(err)
	}
}
