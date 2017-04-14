package protocol

type ActiveRequest struct {
	SerialNum string `json:serial_num`
}

type ActiveResponse struct {
	ErrorCode int `json:error_code`
}

type HeartbeatRequest struct {
	SerialNum string `json:serial_num`
}

type HeartbeatResponse struct {
	ErrorCode int `json:error_code`
}

const (
	OK          = 0
	ReadFailed  = 1
	ParseFailed = 2
	DbFailed    = 3
	Reject      = 4
)
