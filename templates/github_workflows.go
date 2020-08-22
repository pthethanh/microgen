package templates

const (
	// GithubWorkFlows github workflows template.
	GithubWorkFlows = `name: Go

on:
  push:
    branches: [ master ]
  pull_request:
    branches: [ master ]

jobs:

  build:
    name: Build
    runs-on: ubuntu-latest
    steps:

    - name: Set up Go {{.go_version}}
      uses: actions/setup-go@v1
      with:
        go-version: {{.go_version}}
      id: go

    - name: Check out code into the Go module directory
      uses: actions/checkout@v2

    - name: Build
      run: make build_test
`
)
