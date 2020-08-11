package templates

const (
	// ReadMe README.md template.
	ReadMe = `# {{.project_name}}

[![Actions Status](https://{{.module_name}}/workflows/Go/badge.svg)](https://{{.module_name}}/actions)
[![GoDoc](https://godoc.org/{{.module_name}}?status.svg)](https://pkg.go.dev/mod/{{.module_name}})
[![GoReportCard](https://goreportcard.com/badge/{{.module_name}})](https://goreportcard.com/report/{{.module_name}})

## Getting started

Run **make** to generate APIs, docs.

Run **go run main.go** to start the service.

`
)
