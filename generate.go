package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

const baseDir = "./"
const dockerDirPerm = os.FileMode(0755)

var configPath string = filepath.Join(baseDir, "config")
var versionFilesPattern string = filepath.Join(configPath, "*.json")
var dockerFileDir string = filepath.Join(baseDir, "versions")

var dockerFileTemplate = template.Must(
	template.New("dockerfile_template").
		Funcs(template.FuncMap{"StringsJoin": strings.Join}).
		ParseFiles(filepath.Join(baseDir, "dockerfile_template")))

type VersionConfig struct {
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	Jdk       string         `json:"jdk"`
	JvmOpts   string         `json:"jvmOpts"`
	Ports     map[string]int `json:"ports"`
	ExtraDeps []string       `json:"extraDeps"`
}

func main() {
	versionFiles, err := filepath.Glob(versionFilesPattern)
	if err != nil {
		log.Fatal(err)
	}

	for _, versionConfigFile := range versionFiles {
		log.Println("Processing: ", versionConfigFile)
		file, err := os.Open(versionConfigFile)
		if err != nil {
			log.Println(err)
			log.Println("Continuing")
			continue
		}

		var versionConfig VersionConfig
		err = json.NewDecoder(file).Decode(&versionConfig)
		if err != nil {
			log.Println(err)
			log.Println("Continuing")
		}

		processVersion(&versionConfig)
		file.Close()
	}
}

func processVersion(versionConfig *VersionConfig) {
	versionPath := filepath.Join(dockerFileDir, versionConfig.Name)

	versionConfig.Path = versionPath
	log.Println("Generating: ", versionPath)

	createDockerFile(versionConfig)
}

func ensureDirExists(dir string) {
	os.MkdirAll(dir, dockerDirPerm)
}

func createDockerFile(config *VersionConfig) {
	ensureDirExists(config.Path)
	contents := bytes.NewBuffer(nil)
	err := dockerFileTemplate.ExecuteTemplate(contents, "dockerfile_template", config)
	if err != nil {
		log.Println(err)
		log.Println("Continuing")
	}

	// cleanup white space
	regex := regexp.MustCompile("\n\n\n+")
	contentsBytes := regex.ReplaceAll(contents.Bytes(), []byte("\n\n"))

	// save file
	dockerfile, err := os.Create(filepath.Join(config.Path, "Dockerfile"))
	if err != nil {
		log.Println(err)
		log.Println("Continuing")
	}
	dockerfile.Write(contentsBytes)
	dockerfile.Close()
}
