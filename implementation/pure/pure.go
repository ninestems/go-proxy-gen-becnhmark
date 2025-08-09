package pure

import (
	"context"
	"strconv"
)

type Pure struct{}

func New() *Pure {
	return &Pure{}
}

func (p *Pure) Convert(_ context.Context, id string) (int64, error) {
	return strconv.ParseInt(id, 10, 64)
}
