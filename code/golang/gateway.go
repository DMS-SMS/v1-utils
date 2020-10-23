package golang

const (
	// about Unauthorized
	NoAuthorizationInHeader = -101
	InvalidFormatOfAuthorization = -102
	InvalidSignatureOfJWT = -103
	ExpiredJWTToken = -104
	InvalidClaimsOfJWT = -105
	UnsupportedAuthorization = -106

	// about NotFound
	UnsupportedContentType = -201
	FailToBindRequestToStruct = -202
	IntegrityInvalidRequest = -203

	// about ServiceUnavailable
	AvailableServiceNotExist = -301
	CircuitBreakerOpen = -302
)
