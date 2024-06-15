package traceid

import "github.com/google/uuid"

var TraceIdFieldName = "trace_id"

func New() string {
	return uuid.New().String()
}
