package openuao

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"openuao/internal/orchestrator"
	"os"
	"time"

	"github.com/sethvargo/go-envconfig"
)


func NewOrchestrator(opts ...*OrchestratorOptions) *Orchestrator {
	return &Orchestrator{}
}

type Orchestrator struct {
	Config orchestrator.Config
}

func (o *Orchestrator) Run(i *int) {
	for {
		fmt.Println("Running...")
		*i += 1
		time.Sleep(5 * time.Second)
	}
}

type OrchestratorOptions struct {
	Database struct {
		Path 	 string `yaml:"path" json:"path" env:"OPENUAO-DB-PATH"`
		Port 	 string `yaml:"port" json:"port" env:"OPENUAO-DB-PORT"`
		Username string `yaml:"username" json:"username" env:"OPENUAO-DB-USERNAME"`
		Password string `yaml:"password" json:"password" env:"OPENUAO-DB-PASSWORD"`
	}
	Cache struct {
		Path 	 string `yaml:"path" json:"path" env:"OPENUAO-CACHE-PATH"`
		Port 	 string `yaml:"port" json:"port" env:"OPENUAO-CACHE-PORT"`
		Username string `yaml:"username" json:"username" env:"OPENUAO-CACHE-USERNAME"`
		Password string `yaml:"password" json:"password" env:"OPENUAO-CACHE-PASSWORD"`
	}
}

func WithJsonConfig(path string) *OrchestratorOptions {
	var c OrchestratorOptions

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &c
}

func WithYamlConfig(path string) *OrchestratorOptions {
	return nil
}

func WithEnvVars() *OrchestratorOptions {
	var c OrchestratorOptions
	ctx := context.Background()

	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
		return nil
	}

	return &c
}