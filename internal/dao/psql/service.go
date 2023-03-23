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

func ParseStructToSlices(obj any) (tags []string, values []string) {

	obj = dereferencePointer(obj)
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	numFields := t.NumField()
	strCount := 0

	tags = make([]string, 0)
	values = make([]string, 0)

	for i := 0; i < numFields; i++ {

		tag := t.Field(i).Tag.Get(DatabaseStructTag)
		field := v.Field(i)

		if field.IsValid() {
			switch field.Kind() {
			case reflect.String:
				if str := field.String(); str != "" {
					log.Infof("%s: %s", tag, field.String())
					values = append(values, wrapInSingleQuotes(field.String()))
					tags = append(tags, tag)
					strCount++
				}
			default:
				//values[i] = ""
				//tags[i] = tag
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

func wrapInSingleQuotes(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "\\'") + "'"
}

const DatabaseStructTag = "db"
