module github.com/autoscalerhq/autoscaler

go 1.22

replace (
	github.com/autoscalerhq/autoscaler/lib/dkron => ./lib/dkron
	github.com/autoscalerhq/autoscaler/lib/math => ./lib/math
	github.com/autoscalerhq/autoscaler/services/api/middleware => ./services/api/middleware
	github.com/autoscalerhq/autoscaler/services/api/monitoring => ./services/api/monitoring
	github.com/autoscalerhq/autoscaler/services/api/routes => ./services/api/routes
)

require (
	github.com/asecurityteam/rolling v2.0.4+incompatible
	github.com/go-co-op/gocron v1.37.0
	github.com/grafana/pyroscope-go v1.1.1
	github.com/joho/godotenv v1.5.1
	github.com/kevinconway/loadshed/v2 v2.0.0
	github.com/labstack/echo/v4 v4.12.0
	github.com/nats-io/nats.go v1.36.0
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/swaggo/echo-swagger v1.4.1
	github.com/swaggo/swag v1.16.3
	go.opentelemetry.io/otel v1.28.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutlog v0.4.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.28.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.28.0
	go.opentelemetry.io/otel/log v0.4.0
	go.opentelemetry.io/otel/metric v1.28.0
	go.opentelemetry.io/otel/sdk v1.28.0
	go.opentelemetry.io/otel/sdk/log v0.4.0
	go.opentelemetry.io/otel/sdk/metric v1.28.0
	go.opentelemetry.io/otel/trace v1.28.0
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grafana/pyroscope-go/godeltaprof v0.1.6 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/kevinconway/rolling/v3 v3.0.0 // indirect
	github.com/klauspost/compress v1.17.3 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/open-feature/go-sdk v1.12.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/swaggo/files/v2 v2.0.0 // indirect
	github.com/tklauser/go-sysconf v0.3.14 // indirect
	github.com/tklauser/numcpus v0.8.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/exp v0.0.0-20240110193028-0dcbfd608b1e // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.16.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
