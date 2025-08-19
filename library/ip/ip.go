package ip

import (
	"context"
)

type Key struct{}

func ClientIP(ctx context.Context) string {
	ip, _ := ctx.Value(Key{}).(string)
	return ip
}
