module github.com/aniketpuro/devsecops-secure-nginx-platform/microservice

go 1.23

require (
    github.com/gin-gonic/gin v1.10.0
    github.com/golang-jwt/jwt/v5 v5.2.1
    github.com/prometheus/client_golang v1.20.5
    go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.57.0
    go.opentelemetry.io/otel v1.32.0
    go.opentelemetry.io/otel/exporters/prometheus v0.54.0
    go.opentelemetry.io/otel/metric v1.32.0
    go.opentelemetry.io/otel/sdk/metric v1.32.0
)