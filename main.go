package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/pthethanh/microgen/templates"
	tt "github.com/pthethanh/template"
)

func main() {
	moduleName := flag.String("module", "", "module name (ex: github.com/pthethanh/user)")
	projectName := flag.String("name", "", "project name (ex: user)")
	herokuName := flag.String("heroku_app_name", "", "heroku app name")
	web := flag.Bool("web", false, "web app?")
	spa := flag.Bool("spa", false, "single page application?")
	port := flag.String("port", "8000", "port")
	version := flag.String("version", "0.0.1", "app servesion")
	goVersion := flag.String("go_version", "1.14", "golang version")
	protocVersion := flag.String("protoc_version", "3.10.1", "protoc version")
	grpcGWVersion := flag.String("grpc_gw_version", "v1.14.6", "gRPC Gateway version")
	protobufVersion := flag.String("protobuf_version", "v1.4.2", "protobuf version")
	microVersion := flag.String("micro_version", "v0.1.4", "micro version")
	grpcVersion := flag.String("grpc_version", "v1.30.0", "grpc version")

	flag.Parse()
	if *moduleName == "" || *projectName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		p, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		goPath = filepath.Join(p, "go")
	}
	conf := map[string]interface{}{
		"version":          *version,
		"go_version":       *goVersion,
		"protoc_version":   *protocVersion,
		"grpc_gw_version":  *grpcGWVersion,
		"protobuf_version": *protobufVersion,
		"micro_version":    *microVersion,
		"grpc_version":     *grpcVersion,
		"module_name":      *moduleName,
		"project_name":     *projectName,
		"heroku_app_name":  *herokuName,
		"heroku":           *herokuName != "",
		"web":              *web,
		"port":             *port,
		"web_host":         fmt.Sprintf("http://localhost:%s", *port),
		"spa":              *spa,
	}
	if *herokuName != "" {
		conf["web_host"] = fmt.Sprintf("https://%s.herokuapp.com", *herokuName)
	}
	if *spa {
		conf["web"] = true
	}

	rootDir := filepath.Join(goPath, "src", *moduleName)
	getPath := func(s string) string {
		return filepath.Join(rootDir, s)
	}
	tmpls := map[string]string{
		".github/workflows/go.yml":             templates.GithubWorkFlows,
		"api/user/rpc.proto":                   templates.UserProto,
		"api/user/event/rpc.proto":             templates.UserEventProto,
		"deployment/docker/Dockerfile":         templates.Dockerfile,
		"deployment/docker/.dockerignore":      templates.DockerIgnore,
		"deployment/docker/docker-compose.yml": templates.DockerCompose,
		"deployment/docker/.env":               templates.DockerComposeEnv,
		"internal/app/user/user.go":            templates.User,
		".gitignore":                           templates.GitIgnore,
		".gitattributes":                       templates.GitAttributes,
		"go.mod":                               templates.Mod,
		"main.go":                              templates.Main,
		"Makefile":                             templates.Makefile,
		"README.md":                            templates.ReadMe,
		"VERSION":                              templates.Version,
		"doc/":                                 "",
	}
	if *web {
		tmpls["web/"] = ""
		tmpls["web/index.html"] = templates.Web
	}
	for f, t := range tmpls {
		d := getPath(path.Dir(f))
		if _, err := os.Stat(d); err != nil {
			if !os.IsNotExist(err) {
				panic(err)
			}
			if err := os.MkdirAll(d, 0755); err != nil {
				panic(err)
			}
		}
		if t == "" {
			continue
		}
		if err := executeTemplate(getPath(f), t, conf); err != nil {
			panic(err)
		}
	}
	cmd := exec.Command("gofmt", "-w", "-l", ".")
	cmd.Dir = rootDir
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func executeTemplate(targetFile string, templateStr string, conf interface{}) error {
	f, err := os.Create(targetFile)
	t, err := template.New("").Funcs(tt.FuncMap()).Parse(templateStr)
	if err != nil {
		return err
	}
	if err := t.Execute(f, conf); err != nil {
		return err
	}
	return nil
}
