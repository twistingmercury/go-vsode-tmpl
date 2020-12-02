package httpd

import (
	"fmt"
	"rest_api/config"
	"rest_api/handler"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	ddlog "github.com/twistingmercury/go-datadog"
	"github.com/twistingmercury/go-healthcheck"
	"github.com/twistingmercury/go-id4client"

	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
)

// ListenAndServe starts the HTTP server.
func ListenAndServe() error {
	ddcfg := ddlog.DDConfig{
		APIKey:          config.MonitorAPIKey(),
		ServiceName:     config.ServiceName(),
		ServiceVersion:  config.ServiceVer(),
		Environment:     config.Environment(),
		LogIntakePort:   config.LoggingIntakePort(),
		LogIntakeURL:    config.LoggingIntakeURL(),
		TraceIntakeHost: config.TracingIntakeURL(),
		TraceIntakePort: config.TracingIntakePort(),
		GlobalTags: map[string]interface{}{
			"commit":        config.GitHash(),
			"environment":   config.Environment(),
			"version":       config.ServiceVer(),
			"service":       config.ServiceName(),
			"source":        "go",
			"my-custom-tag": "hello, there!",
		},
	}
	ddlog.Initialize(ddcfg)
	defer ddlog.Stop()

	go ListenAndServeHeartbeat(config.Dependencies()...)
	return ListenAndServeAPI()
}

// ListenAndServeHeartbeat starts the heartbeat endpoint.
func ListenAndServeHeartbeat(cfgd ...config.Dependency) {
	deps := make([]healthcheck.DependencyDescriptor, 0)

	for _, c := range cfgd {
		deps = append(deps, healthcheck.DependencyDescriptor{
			Name:       c.Name,
			Connection: c.Connection,
		})
	}

	r := gin.Default()
	r.GET("/ops/heartbeat", healthcheck.Handler(deps...))
	r.Run(fmt.Sprintf(":%d", config.HealthcheckPort()))
}

// ListenAndServeAPI starts the API endpoint(s).
func ListenAndServeAPI() error {
	r := newRouter()
	r.GET("/api/greeting", handler.Get())
	return r.Run(fmt.Sprintf(":%d", config.APIPort()))
}

func newRouter() (r *gin.Engine) {
	c := id4client.IdentityConfig{
		BaseURL:        config.OAuthProvider().BaseURL,
		ID:             config.OAuthProvider().ID,
		Secret:         config.OAuthProvider().Secret,
		IntrospectPath: config.OAuthProvider().IntrospectPath,
		TokenPath:      config.OAuthProvider().TokenPath,
		ServiceName:    config.ServiceName(),
		ServiceVersion: config.ServiceVer(),
		CommitHash:     config.GitHash(),
	}

	if err := id4client.Initialize(c); err != nil {
		logrus.Fatal(err.Error())
	}

	r = gin.New()
	r.Use(gin.Recovery())
	r.Use(ddlog.Monitor())
	r.Use(id4client.Authenticate())
	r.Use(gintrace.Middleware("rest-api"))
	return
}
