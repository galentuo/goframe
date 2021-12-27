package health

import (
	"time"

	"github.com/galentuo/goframe"
)

type health struct {
	goframe.HTTPService
}

func HealthService(appName string) *health {
	srv := goframe.NewHTTPService("")
	srv.Route("/", "GET", func(ctx goframe.ServerContext) error {
		data := make(map[string]interface{})
		data["app"] = appName
		data["time"] = time.Now().Unix()
		return ctx.Response().SuccessJSON(200, data, "Up & Running")
	})
	return &health{srv}
}
