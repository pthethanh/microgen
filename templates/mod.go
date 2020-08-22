package templates

const (
	// Mod go module template.
	Mod = `module {{.module_name}}

go {{.go_version}}

require (
	github.com/golang/protobuf {{.protobuf_version}}
	github.com/grpc-ecosystem/grpc-gateway {{.grpc_gw_version}}
	github.com/pthethanh/micro {{.micro_version}}
	google.golang.org/grpc {{.grpc_version}}
)
	`
)
