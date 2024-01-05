package rate_limiter

type Request struct {
	RequestType RequestType

	IP         string
	Identifier string
}

type RequestType string

//nolint:staticcheck
const (
	CommonType       RequestType = "common"
	AuthType                     = "auth"
	RegistrationType             = "registration"
)
