package main

import (
	"context"
	"log"
	"os"

	"github.com/kohTkd/experimental_learning/ent"
	"github.com/kohTkd/experimental_learning/ent/migrate"
	"github.com/kohTkd/experimental_learning/internal/config"
	"github.com/kohTkd/experimental_learning/internal/infrastructure/repository"
	"github.com/kohTkd/experimental_learning/internal/util/environment"
)

func init() {
	var envName string
	if envName = os.Getenv("ENV_NAME"); envName == "" {
		envName = "development"
	}
	var env *environment.EnvName
	if env = environment.Get(envName); env == nil {
		log.Panic("invalid environment name: " + envName)
	}
	config.Read(config.ReadConfigOption{Env: *env})
}

func main() {
	client, err := repository.NewClient()
	if err != nil {
		log.Fatalf("failed opening mysql client: %v", err)
	}
	defer client.Close()
	createDBSchema(client.Writer)
}

func createDBSchema(client *ent.Client) {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
