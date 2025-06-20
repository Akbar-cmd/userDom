package metric

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"time"
)

const (
	namespace = "user_space"
	appName   = "user_app"
)

type ClientMetrics struct {
	requestsSent      prometheus.Counter
	responsesReceived *prometheus.CounterVec
	requestDuration   *prometheus.HistogramVec
}

var clientMetrics *ClientMetrics

func InitClientMetrics() error {
	labels := []string{"service", "method", "status"}

	clientMetrics = &ClientMetrics{
		requestsSent: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "grpc_client",
			Name:      "requests_sent_total",
			Help:      "Total number of gRPC requests sent",
		}),

		responsesReceived: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "grpc_client",
			Name:      "responses_received_total",
			Help:      "Total number of gRPC responses received",
		}, labels),

		requestDuration: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: "grpc_client",
			Name:      "request_duration_seconds",
			Help:      "gRPC request latency distribution",
			Buckets:   prometheus.ExponentialBuckets(0.001, 2, 16),
		}, labels),
	}

	return nil
}

func ClientMetricsInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		clientMetrics.requestsSent.Inc()
		start := time.Now()

		err := invoker(ctx, method, req, reply, cc, opts...)
		duration := time.Since(start)

		grpcStatus, _ := status.FromError(err)
		statusCode := grpcStatus.Code().String()

		serviceName := "bike_service" // или извлекать из method

		clientMetrics.responsesReceived.WithLabelValues(serviceName, method, statusCode).Inc()
		clientMetrics.requestDuration.WithLabelValues(serviceName, method, statusCode).Observe(duration.Seconds())

		return err
	}
}
