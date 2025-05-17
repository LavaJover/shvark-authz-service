package rbac

import (
	"log"

	"github.com/LavaJover/shvark-authz-service/internal/config"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitEnforcer(cfg *config.AuthzConfig) *casbin.Enforcer {
	db, err := gorm.Open(postgres.Open(cfg.AuthzDB.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to db: %v\n", err)
	}

	// casbin postgres adapter
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("failed to create adapter: %v\n", err)
	}

	// loading models
	m, err := model.NewModelFromFile("./model.conf")
	if err != nil {
		log.Fatalf("failed to load model: %v\n", err)
	}

	// creating new enforcer
	e, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		log.Fatalf("failed to create enforcer: %v\n", err)
	}

	// loading policies from db
	if err := e.LoadPolicy(); err != nil {
		log.Fatalf("failed to load policy: %v\n", err)
	}

	return e
}