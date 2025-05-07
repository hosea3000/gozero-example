package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"gozero_example/server/internal/config"
	"gozero_example/server/internal/handler"
	"gozero_example/server/internal/svc"
	"gozero_example/server/response"
	"gozero_example/server/response/http_error"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/server-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(ErrorHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func ErrorHandler(err error) (int, any) {
	logx.Errorf("handle error: %s", err)

	var e *http_error.APIError
	if !errors.As(err, &e) {
		e = http_error.ErrInternalServiceError // default error
	}

	return e.HTTPStatus, response.ErrorFormat(e)
}
