module github.com/alexeykirinyuk/go_grpc_workshop/product_service

go 1.19

require (
	github.com/alexeykirinyuk/go_grpc_workshop/product_service/pkg/product_service v0.0.0-00010101000000-000000000000
	github.com/golang/mock v1.6.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.28.0
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.49.0
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.6.7 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sys v0.0.0-20220610221304-9f5ed59c137d // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace github.com/alexeykirinyuk/go_grpc_workshop/product_service/pkg/product_service => ./pkg/product_service
