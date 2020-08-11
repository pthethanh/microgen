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

    - name: Set up Go {{.version_go}}
      uses: actions/setup-go@v1
      with:
        go-version: {{.version_go}}
      id: go

    - name: Check out code into the Go module directory
      uses: actions/checkout@v2

    - name: Build
      run: make all
	`
)
