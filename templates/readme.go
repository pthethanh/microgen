package templates

const (
	// ReadMe README.md template.
	ReadMe = `# {{.project_name}}

[![Actions Status](https://{{.module_name}}/workflows/Go/badge.svg)](https://{{.module_name}}/actions)
[![GoDoc](https://godoc.org/{{.module_name}}?status.svg)](https://pkg.go.dev/mod/{{.module_name}})
[![GoReportCard](https://goreportcard.com/badge/{{.module_name}})](https://goreportcard.com/report/{{.module_name}})

## Getting started

` + "```go" + `
// Install Protocol Buffer if you have not done yet.
make install_protobuf

// Generate APIs, docs and verify the code.
make

// Start the service.
go run main.go

// Check APIs
curl http://localhost:8000/internal/readiness
curl http://localhost:8000/internal/liveness
curl http://localhost:8000/internal/metrics
curl http://localhost:8000/api/v1/users
` + "```" + `
`
)
