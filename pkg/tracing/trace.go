package tracing

import (
	"gin_init/pkg/setting"
	"io"
	"time"

	"github.com/gin-gonic/gin"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/zipkin"
)

// Init 初始化Jeager
func Init(service string) (opentracing.Tracer, io.Closer) {
	zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	injector := jaeger.TracerOptions.Injector(opentracing.HTTPHeaders, zipkinPropagator)
	extractor := jaeger.TracerOptions.Extractor(opentracing.HTTPHeaders, zipkinPropagator)

	// Zipkin shares span ID between client and server spans; it must be enabled via the following option.
	zipkinSharedRPCSpan := jaeger.TracerOptions.ZipkinSharedRPCSpan(true)

	sender, _ := jaeger.NewUDPTransport(setting.JaegerSetting.Host, 0)
	tracer, closer := jaeger.NewTracer(
		setting.JaegerSetting.Server,
		jaeger.NewConstSampler(true),
		jaeger.NewRemoteReporter(
			sender,
			jaeger.ReporterOptions.BufferFlushInterval(1*time.Second)),
		injector,
		extractor,
		zipkinSharedRPCSpan,
	)
	return tracer, closer
}

//StartSpan 开始一个上下文span,用于controllers
//c *gin.Context 上下文
//name span名字
//return 上下文的新span
func StartSpan(c *gin.Context, name string) opentracing.Span {
	span := opentracing.SpanFromContext(c.Request.Context())
	tracer := opentracing.GlobalTracer()
	parentSpan := tracer.StartSpan(name, opentracing.ChildOf(span.Context()))
	return parentSpan
}

//StartFromSpan 开始一个子span,用于连接redis,mysql等
//parentSpan 父类span名字
//return 子类span
func StartFromSpan(parentSpan opentracing.Span, name string) opentracing.Span {
	t := opentracing.GlobalTracer()
	childSpan := t.StartSpan(name, opentracing.ChildOf(parentSpan.Context()))
	return childSpan
}
