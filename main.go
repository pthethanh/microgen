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
)

var (
	defaultConf = map[string]interface{}{
		"version":          "0.0.1",
		"version_go":       "1.14",
		"version_protoc":   "3.10.1",
		"version_grpc_gw":  "v1.14.6",
		"version_protobuf": "v1.4.2",
		"version_micro":    "v0.1.2",
		"version_grpc":     "v1.29.1",
		"web":              false,
		"heroku":           true,
		"port":             "8000",
	}
)

func main() {
	moduleName := flag.String("module", "", "module name (ex: github.com/pthethanh/user)")
	projectName := flag.String("name", "", "project name (ex: user)")
	herokuName := flag.String("heroku_app_name", "", "heroku app name")
	web := flag.Bool("web", false, "web app?")
	port := flag.String("port", "8000", "port")
	flag.Parse()
	if *moduleName == "" || *projectName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	goPath := os.Getenv("GO_PATH")
	if goPath == "" {
		p, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		goPath = filepath.Join(p, "go")
	}
	conf := defaultConf
	conf["module_name"] = *moduleName
	conf["project_name"] = *projectName
	conf["heroku_app_name"] = *herokuName
	conf["web"] = *web
	conf["port"] = *port
	conf["web_host"] = fmt.Sprintf("http://localhost:%s", *port)
	if *herokuName != "" {
		conf["web_host"] = fmt.Sprintf("https://%s.herokuapp.com", *herokuName)
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
	t, err := template.New("").Parse(templateStr)
	if err != nil {
		return err
	}
	if err := t.Execute(f, conf); err != nil {
		return err
	}
	return nil
}
