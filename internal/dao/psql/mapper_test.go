package psql

import (
	"github.com/calebtracey/mind-your-business-api/external"
	"github.com/calebtracey/mind-your-business-api/external/models"
	"github.com/jackc/pgx/v5/pgtype"
	"testing"
	"time"
)

func TestMapper_PostgresExec(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		request *external.ApiRequest
		want    string
	}{
		{
			name: "Happy Path",
			request: &external.ApiRequest{
				Request: external.Request{
					User: &models.User{
						ID:           pgtype.UUID{},
						FirstName:    "TEST",
						LastName:     "TEST",
						Email:        "TEST",
						Username:     "TEST",
						Password:     "TEST",
						Token:        "TEST",
						RefreshToken: "TEST",
						CreatedAt:    mockTime,
						UpdatedAt:    mockTime,
					},
				},
			},
			want: "insert into users (\"[id first_name last_name email username password token refresh_Token created_at updated_at]\") values ('[{\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000 %!s(bool=false)} TEST TEST TEST TEST TEST TEST TEST 2023-03-22 10:30:00 +0000 UTC 2023-03-22 10:30:00 +0000 UTC]')",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Mapper{}
			if got := m.PostgresExec(tt.request); got != tt.want {
				t.Errorf("PostgresExec() = %v, want %v", got, tt.want)
			}
		})
	}
}

var mockTime = getFakeTime()

func getFakeTime() time.Time {
	year, month, day := 2023, time.March, 22
	hour, min, sec := 10, 30, 0
	return time.Date(year, month, day, hour, min, sec, 0, time.UTC)
}
