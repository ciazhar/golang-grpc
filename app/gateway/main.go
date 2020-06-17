package main

import (
	"flag"
	"fmt"
	"github.com/ciazhar/golang-grpc/grpc/generated/golang"
	"net/http"
	"strings"

	"google.golang.org/grpc"

	"github.com/dgrijalva/jwt-go"
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
	err := golang.RegisterRecipesServiceHandlerFromEndpoint(ctx, mux, *serverEndpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
	return
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func checkHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Authorization"); auth != "" {
			if authHeader := strings.Split(auth, " "); len(authHeader) == 2 {

				//get claims
				claims := jwt.MapClaims{}
				if _, err := jwt.ParseWithClaims(authHeader[1], claims, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}
					return []byte("ThisIsSecretForJWTHS512SignatureAlgorithmThatMUSTHave512bitsKeySize"), nil
				}); err != nil {
					return
				}

				//get authorities
				authorities := make([]interface{}, 0)
				if claims["authorities"] != nil {
					authorities = claims["authorities"].([]interface{})
				}

				//get
				if strings.HasPrefix(r.URL.String(), "/server") {
					if exist := contains(authorities, "ROLE_USER"); exist {

					}
				}
			}
		}
		h.ServeHTTP(w, r)
	})
}

func contains(s []interface{}, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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

	return http.ListenAndServe(address, checkHeader(allowCORS(mux)))

}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := Run(":8080"); err != nil {
		glog.Fatal(err)
	}
}
