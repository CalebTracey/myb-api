package main

import (
	config "github.com/calebtracey/config-yaml"
	"github.com/calebtracey/mind-your-business-api/internal/dao/psql"
	"github.com/calebtracey/mind-your-business-api/internal/facade"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Infoln("=== MIND YOUR BUSINESS API ===")
}

type InitializerI interface {
	Database(cfg *config.Config, svc *facade.Service) error
}

type Initializer struct{}

func (i *Initializer) Database(cfg *config.Config, svc *facade.Service) error {
	if psqlService, err := cfg.Database(PostgresDB); err != nil {
		return err
	} else {
		svc.PsqlDAO = psql.DAO{Pool: psqlService.Pool}
		svc.PsqlMapper = psql.Mapper{}
	}
	return nil
}

const PostgresDB = "PSQL"
