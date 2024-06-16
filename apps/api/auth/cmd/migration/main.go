package main

import (
	"context"
	"log"
	"os"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/kohTkd/experimental_learning/ent/migrate"
	"github.com/kohTkd/experimental_learning/internal/config"
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
	dsn := config.Conf.Database.Writer.MigrationDsn()
	ctx := context.Background()

	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("atlas/migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                          // provide migration directory
		schema.WithMigrationMode(schema.ModeInspect), // provide migration mode
		schema.WithDialect(dialect.MySQL),            // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).

	if err = migrate.NamedDiff(ctx, dsn, os.Args[1], opts...); err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
