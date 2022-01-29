package entity

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}
type ErrorResponse struct {
	httpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{
		httpSC: 400,
		Error:  Err{Error: "Request Body is not correct", ErrorCode: "001"},
	}
	ErrorNotAuthUser = ErrorResponse{
		httpSC: 401,
		Error:  Err{Error: "User Authentication failed", ErrorCode: "002"},
	}
)
