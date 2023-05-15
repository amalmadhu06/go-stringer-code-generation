package http

//go:generate stringer -type=StatusCode
type StatusCode int

const (
	Ok               StatusCode = 200
	Created          StatusCode = 201
	Accepted         StatusCode = 202
	MultipleChoice   StatusCode = 300
	MovedPermanently StatusCode = 301
	Found            StatusCode = 302
	BadRequest       StatusCode = 400
	Unauthorized     StatusCode = 401
	PaymentRequired  StatusCode = 402
)
