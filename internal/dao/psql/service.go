package psql

import (
	"context"
	"github.com/calebtracey/mind-your-business-api/external"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

type DAOI interface {
	ExecContext(ctx context.Context, exec string) (resp *external.ExecResponse, err error)
}

type DAO struct {
	Pool *pgxpool.Pool
}

func (s DAO) ExecContext(ctx context.Context, exec string) (resp *external.ExecResponse, err error) {
	resp = new(external.ExecResponse)

	if resp.Status, err = s.Pool.Exec(ctx, exec); err != nil {
		log.Errorf("ExecContext: %v", err)
		return nil, err
	} else {
		return resp, nil
	}
}

func dereferencePointer(obj any) any {
	if reflect.ValueOf(obj).Kind() == reflect.Pointer {
		obj = reflect.ValueOf(obj).Elem().Interface()
	}
	return obj
}

func wrapInSingleQuotes(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "\\'") + "'"
}

const DatabaseStructTag = "db"
