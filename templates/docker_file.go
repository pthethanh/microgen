package templates

const (
	// Dockerfile docker file for service.
	Dockerfile = `FROM alpine:3.8

WORKDIR /home/
COPY {{.project_name}}.bin .
RUN chmod +x {{.project_name}}.bin
{{if .web}}COPY web/* ./web{{end}}

EXPOSE {{.port}}
CMD ["./{{.project_name}}.bin"]
`
)
