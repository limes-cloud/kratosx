package ip

import "context"

func ClientIP(ctx context.Context) string {
	ip, _ := ctx.Value("ClientIP").(string)
	return ip
}
