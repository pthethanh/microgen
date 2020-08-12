package templates

const (
	// DockerCompose docker compose template.
	DockerCompose = `version: '3'

services:
  {{.project_name}}:
    image: "{{.project_name}}:{{.version}}"
    hostname: {{.project_name}}
    ports:
      - "{{.port}}:{{.port}}"
	environment:
      # LOGGING
      - LOG_FORMAT=${LOG_FORMAT}
      # SERVER
      - ADDRESS=${ADDRESS}
      - JWT_SECRET=${JWT_SECRET}
	restart: always
`

	// DockerComposeEnv docker compose .env template.
	DockerComposeEnv = `# LOGGING
LOG_FORMAT=json

# SERVER
ADDRESS=:{{.port}}
`
)
