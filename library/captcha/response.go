package captcha

import "time"

type Response interface {
	ID() string
	Expire() time.Duration
	Base64String() string
}

type response struct {
	id     string
	base64 string
	expire time.Duration
}

func (r *response) ID() string {
	return r.id
}

func (r *response) Expire() time.Duration {
	return r.expire
}

func (r *response) Base64String() string {
	return r.base64
}
