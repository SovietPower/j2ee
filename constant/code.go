package constant

const (
	SUCCESS                 = 200
	UPDATE_PASSWORD_SUCCESS = 201
	NOT_EXISTED_IDENTIFIER  = 202
	ERROR                   = 500
	INVALID_PARAMS          = 400

	ERROR_EXISTED_USER         = 10001
	ERROR_NO_SUCH_USER         = 10002
	ERROR_WRONG_PASSWORD       = 10003
	ERROR_WRONG_CAPTCHA        = 10004
	ERROR_FAIL_GETTING_CAPTCHA = 10005
	ERROR_FAIL_ENCRYPTION      = 10006
	ERROR_NO_SUCH_ADDRESS      = 10007
	ERROR_NO_SUCH_ORDER        = 10008

	ERROR_AUTH_CHECK_TOKEN_FAIL       = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT    = 20002
	ERROR_AUTH_TOKEN                  = 20003
	ERROR_AUTH                        = 20004
	ERROR_AUTH_INSUFFICIENT_AUTHORITY = 20005
	ERROR_READ_FILE                   = 20006
	ERROR_SEND_EMAIL                  = 20007
	ERROR_CALL_API                    = 20008
	ERROR_UNMARSHAL_JSON              = 20009

	ERROR_DATABASE      = 30001
	ERROR_NO_PERMISSION = 30002
	ERROR_UPDATING_INFO = 30003

	ERROR_OSS = 40001
)
