package constants

type ResponseMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

const (
	MESSAGE_SUCCESS              = "Success"
	MESSAGE_FAIL                 = "Fail"
	MESSAGE_INVALID              = "Invalid"
	MESSAGE_NOTFOUND             = "Not Found"
	MESSAGE_DUPLICATE            = "Duplicate"
	STATUS_OK                    = 200
	STATUS_CREATED               = 201
	STATUS_BAD_REQUEST           = 400
	STATUS_UNAUTHORIZED          = 401
	STATUS_FORBIDDEN             = 403
	STATUS_NOT_FOUND             = 404
	STATUS_INTERNAL_SERVER_ERROR = 500
)

var (
	SUCCESS = ResponseMessage{
		Code:    "0000",
		Message: MESSAGE_SUCCESS,
	}
	FAILED = ResponseMessage{
		Code:    "0001",
		Message: MESSAGE_FAIL,
	}
	NOT_FOUND = ResponseMessage{
		Code:    "0002",
		Message: MESSAGE_NOTFOUND,
	}
	DUPLICATE = ResponseMessage{
		Code:    "0003",
		Message: MESSAGE_DUPLICATE,
	}
	INVALID = ResponseMessage{
		Code:    "0004",
		Message: MESSAGE_INVALID,
	}
)
