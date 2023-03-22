package psql

import (
	"context"
	"github.com/calebtracey/mind-your-business-api/external"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"reflect"
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

type ArgMap map[string]any

func ParseStructToSlices(obj any) (tags []string, values []any) {
	obj = dereferencePointer(obj)

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	numFields := t.NumField()

	tags = make([]string, numFields)
	values = make([]any, numFields)

	for i := 0; i < numFields; i++ {
		tag := t.Field(i).Tag.Get(DatabaseStructTag)
		field := v.Field(i)
		if field.IsValid() {
			switch field.Kind() {
			case reflect.String:
				log.Infof("%s: %s", tag, field.String())
				values[i] = field.String()
				tags[i] = tag
			default:
				values[i] = field.Interface()
				tags[i] = tag
			}
		}
	}
	return tags, values
}

func dereferencePointer(obj any) any {
	if reflect.ValueOf(obj).Kind() == reflect.Pointer {
		obj = reflect.ValueOf(obj).Elem().Interface()
	}
	return obj
}

const DatabaseStructTag = "db"
