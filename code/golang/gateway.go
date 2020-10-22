package golang

const (
	// about Unauthorized
	NoAuthorizationInHeader = -101
	InvalidAuthorizationFormat = -102
	InvalidSignatureOfJWT = -103
	ExpiredJWTToken = -104
	InvalidClaimsOfJWT = -105
)
