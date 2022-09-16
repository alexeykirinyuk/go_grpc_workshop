module github.com/alexeykirinyuk/go_grpc_workshop/product_service

go 1.19

require (
	github.com/Masterminds/squirrel v1.5.3
	github.com/alexeykirinyuk/go_grpc_workshop/category-service/pkg/category-service v0.0.0-20220908222502-03bacbf99651
	github.com/alexeykirinyuk/go_grpc_workshop/product_service/pkg/product_service v0.0.0-00010101000000-000000000000
	github.com/golang/mock v1.6.0
	github.com/jackc/pgx/v4 v4.17.2
	github.com/jmoiron/sqlx v1.3.5
	github.com/pkg/errors v0.9.1
	github.com/pressly/goose/v3 v3.7.0
	github.com/rs/zerolog v1.28.0
	github.com/stretchr/testify v1.8.0
	google.golang.org/genproto v0.0.0-20220908141613-51c1cc9bc6d0
	google.golang.org/grpc v1.49.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.6.8 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/net v0.0.0-20220907135653-1e95f45603a7 // indirect
	golang.org/x/sys v0.0.0-20220908150016-7ac13a9a928d // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace github.com/alexeykirinyuk/go_grpc_workshop/product_service/pkg/product_service => ./pkg/product_service
