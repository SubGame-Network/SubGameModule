package domain

type ErrorFormat struct {
	Code    int
	Message string
}

var (
	ErrorUnauthorized     = ErrorFormat{Code: 401, Message: "Unauthorized"}
	ErrorForbidden        = ErrorFormat{Code: 403, Message: "Forbidden"}
	ErrorAuthTokenExpired = ErrorFormat{Code: 4011, Message: "auth token expired"}
	ErrorBadRequest       = ErrorFormat{Code: 400, Message: "bad request"}
	ErrorServer           = ErrorFormat{Code: 500, Message: "Server Error"}

	ErrorPasswordRules            = ErrorFormat{Code: 4001, Message: "password rules do not match"}
	ErrorNotSubGameAddress        = ErrorFormat{Code: 4002, Message: "bad request address"}
	ErrorEmailSendingLimit        = ErrorFormat{Code: 4003, Message: "Email Sending over Today Limit"}
	ErrorEmailSendingin60secLimit = ErrorFormat{Code: 4004, Message: "Email Sending in 60 sec Limit"}
	ErrorNotFound                 = ErrorFormat{Code: 4005, Message: "data not found"}
	ErrorEmailExpired             = ErrorFormat{Code: 4006, Message: "error email expired"}
	ErrorEmailFormat              = ErrorFormat{Code: 4007, Message: "error email format"}

	// User
	ErrorUserAlreadyVerify    = ErrorFormat{Code: 5001, Message: "User is already verify"}
	ErrorUserNotMatch         = ErrorFormat{Code: 5002, Message: "input user data is not match"}
	ErrorAlreadyExistsDB      = ErrorFormat{Code: 5003, Message: "input user data already exists in table"}
	ErrorVerifyCode           = ErrorFormat{Code: 5004, Message: "Error verify code"}
	ErrorUserNotFound         = ErrorFormat{Code: 5005, Message: "user not found"}
	ErrorBalanceNotEnough     = ErrorFormat{Code: 5006, Message: "balance not enough"}
	ErrorUserPermissionDenied = ErrorFormat{Code: 5007, Message: "user permission denied"}
	ErrorUserEmailExist       = ErrorFormat{Code: 5008, Message: "user input email exist"}

	// Event 監聽錯誤
	ErrorEventSubGameWalletNotActive = ErrorFormat{Code: 1000, Message: "SubGame wallet not active"}
)
