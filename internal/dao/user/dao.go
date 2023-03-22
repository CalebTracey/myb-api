package user

import (
	"context"
	"github.com/calebtracey/mind-your-business-api/external"
	"github.com/calebtracey/mind-your-business-api/internal/dao/psql"
)

type DAOI interface {
	NewUser(ctx context.Context, exec string) (resp *external.ExecResponse, errs []error)
}

type DAO struct {
	PSQL psql.DAOI
}

// NewUser TODO probably dont need this, can just call what is in here inside of the facade function
func (s DAO) NewUser(ctx context.Context, exec string) (resp *external.ExecResponse, errs []error) {
	var err error
	if resp, err = s.PSQL.ExecContext(ctx, exec); err != nil {
		return resp, []error{err}
	}

	// TODO mapping stuff here

	return resp, nil
}
