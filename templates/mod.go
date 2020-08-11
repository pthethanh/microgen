package templates

const (
	// Mod go module template.
	Mod = `module {{.module_name}}

go {{.version_go}}

require (
	github.com/golang/protobuf {{.version_protobuf}}
	github.com/grpc-ecosystem/grpc-gateway {{.version_grpc_gw}}
	github.com/pthethanh/micro {{.version_micro}}
	google.golang.org/grpc {{.version_grpc}}
)
	`
)
