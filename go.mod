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
	github.com/google/uuid v1.6.0
	github.com/grafana/pyroscope-go v1.1.1
	github.com/joho/godotenv v1.5.1
	github.com/kevinconway/loadshed/v2 v2.0.0
	github.com/labstack/echo/v4 v4.12.0
	github.com/nats-io/nats.go v1.36.0
	github.com/open-feature/go-sdk v1.12.0
	github.com/open-feature/go-sdk-contrib/providers/flagd v0.2.2
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/stretchr/testify v1.9.0
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
	buf.build/gen/go/open-feature/flagd/connectrpc/go v1.16.1-20240215170432-1e611e2999cc.1 // indirect
	buf.build/gen/go/open-feature/flagd/grpc/go v1.3.0-20240215170432-1e611e2999cc.2 // indirect
	buf.build/gen/go/open-feature/flagd/protocolbuffers/go v1.33.0-20240215170432-1e611e2999cc.1 // indirect
	connectrpc.com/connect v1.16.1 // indirect
	connectrpc.com/otelconnect v0.7.0 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/barkimedes/go-deepcopy v0.0.0-20220514131651-17c30cfc62df // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/diegoholiveira/jsonlogic/v3 v3.5.2 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-logr/zapr v1.3.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/grafana/pyroscope-go/godeltaprof v0.1.6 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kevinconway/rolling/v3 v3.0.0 // indirect
	github.com/klauspost/compress v1.17.3 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/open-feature/flagd-schemas v0.2.9-0.20240408192555-ea4f119d2bd7 // indirect
	github.com/open-feature/flagd/core v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/swaggo/files/v2 v2.0.0 // indirect
	github.com/tklauser/go-sysconf v0.3.14 // indirect
	github.com/tklauser/numcpus v0.8.0 // indirect
	github.com/twmb/murmur3 v1.1.8 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240415180920-8c6c420018be // indirect
	google.golang.org/grpc v1.63.2 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apimachinery v0.29.3 // indirect
	k8s.io/klog/v2 v2.120.1 // indirect
	k8s.io/utils v0.0.0-20240310230437-4693a0247e57 // indirect
	sigs.k8s.io/controller-runtime v0.17.3 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
)
