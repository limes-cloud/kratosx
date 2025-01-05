package ip

import "context"

type IPKey struct {
}

func ClientIP(ctx context.Context) string {
	ip, _ := ctx.Value(IPKey{}).(string)
	return ip
}
