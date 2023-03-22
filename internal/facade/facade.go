package facade

import (
	"context"
	"github.com/calebtracey/mind-your-business-api/external"
	"github.com/calebtracey/mind-your-business-api/internal/dao/psql"
	"github.com/calebtracey/models/pkg/response"
	log "github.com/sirupsen/logrus"
)

type ServiceI interface {
	NewUser(ctx context.Context, params any) (resp *external.Response)
}

type Service struct {
	PsqlDAO    psql.DAOI
	PsqlMapper psql.MapperI
}

func (s Service) NewUser(ctx context.Context, params any) (resp *external.Response) {
	resp = new(external.Response)
	// TODO add request validation
	// TODO parse params and map request query

	if pgResp, pgErr := s.PsqlDAO.ExecContext(ctx, s.PsqlMapper.PostgresExec()); pgErr != nil {
		resp.Message.ErrorLog = errorLog(pgErr, "NewUser")
	} else {
		log.Infoln(pgResp.Status)
	}

	// TODO add response mapping

	return resp
}

func errorLog(err error, trace string) response.ErrorLogs {
	return response.ErrorLogs{
		{
			RootCause:  err.Error(),
			Trace:      trace,
			StatusCode: "500",
		},
	}
}
