package custom_error

const (
	SuccessMsg             = "Success"
	UnauthorizedMsg        = "Unauthorized"
	ServerErrMsg           = "Internal Server Error"
	ValidationErrMsg       = "Validation Error"
	MissingFieldErrMsg     = "Some field is missing"
	PermissionDeniedErrMsg = "Permission Denied"
	ParsingErrMsg          = "Error while parsing"
)

const (
	MissingField = "%v is missing"
	InvalidField = "%v is invalid"
)

const (
	ErrorWhileReadingFile = "err %v while reading a file from path %v"
)

const (
	InvalidAcquirerName = "%v as acquirer name is not identified"
	InvalidDumpType     = "%v as dump type is not identified"
)

const (
	SuccessStatus = 0
	ErrStatus     = 1
)

const (
	GET_REQ_URL_MISSING  = "URL"
	POST_REQ_URL_MISSING = "URL"
)
