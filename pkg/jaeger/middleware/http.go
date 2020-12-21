package middleware

import (
	"context"
	"net/http"

	"github.com/jaegertracing/jaeger-client-go"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/pengsrc/go-shared/buffer"

	"example/constants"
)

// TraceSpan is a middleware that initialize a tracing span and injects span
// context to r.Context(). In one word, this middleware kept an eye on the
// whole HTTP request that the server receives.
func TraceSpan(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tracer := opentracing.GlobalTracer()
		if tracer == nil {
			// Tracer not found, just skip.
			next.ServeHTTP(w, r)
		}

		buf := buffer.GlobalBytesPool().Get()
		buf.AppendString("HTTP ")
		buf.AppendString(r.Method)

		// Start span.
		span := opentracing.StartSpan(buf.String())
		rc := opentracing.ContextWithSpan(r.Context(), span)

		// Set request ID for context.
		if sc, ok := span.Context().(jaeger.SpanContext); ok {
			rc = context.WithValue(rc, constants.RequestID, sc.TraceID().String())
		}

		next.ServeHTTP(w, r.WithContext(rc))

		// Finish span.
		wrapper, ok := w.(WrapResponseWriter)
		if ok {
			ext.HTTPStatusCode.Set(span, uint16(wrapper.Status()))
		}
		span.Finish()
	}
	return http.HandlerFunc(fn)
}
