package external

import (
	"github.com/calebtracey/models/pkg/response"
	"github.com/jackc/pgx/v5/pgconn"
)

type Response struct {
	Details []any            `json:"Details,omitempty"`
	Message response.Message `json:"Message,omitempty"`
}

type ExecResponse struct {
	Status pgconn.CommandTag `json:"Status"`
}
