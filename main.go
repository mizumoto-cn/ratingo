package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mizumoto-cn/ratingo/analyzer"
	"github.com/mizumoto-cn/ratingo/collector"
	"github.com/mizumoto-cn/ratingo/pkg/data"
	"github.com/mizumoto-cn/ratingo/pkg/display"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/zipkin"
)

func main() {
	// Initialize Jaeger Tracer
	cfg := jaegercfg.Configuration{
		ServiceName: "myapp",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "localhost:14268",
		},
	}

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jaeger.StdLogger),
		jaegercfg.ZipkinSharedRPCSpan(true),
		jaegercfg.Injector(opentracing.HTTPHeaders, zipkin.Propagator{}),
		jaegercfg.Extractor(opentracing.HTTPHeaders, zipkin.Propagator{}),
	)
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	defer closer.Close()

	db, err := data.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	opentracing.SetGlobalTracer(tracer)

	// Set up Gin engine and routes
	r := gin.Default()
	r.Use(ginhttp.Middleware(tracer))

	v1 := r.Group("/v1")
	{
		v1.POST("/collect", collector.Collect(db))
		v1.POST("/analyze", analyzer.Analyze(db))
		v1.GET("/display", display.Display(db))
	}

	r.Run()
}
