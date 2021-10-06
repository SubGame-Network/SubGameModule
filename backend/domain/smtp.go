package domain

type Sender struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	ReplyTo string `json:"replyTo"`
}
type Email struct {
	Email string `json:"email"`
}
type Payload struct {
	Sender      *Sender `json:"sender"`
	To          []Email `json:"to"`
	RepleyTo    Email   `json:"replyTo"`
	HtmlContent string  `json:"htmlContent"`
	Subject     string  `json:"subject"`
}

const RedisStatusEmailVerify = "emailVerify"
const RedisStatusAdminEmailVerify = "emailAdminVerify"
const RedisStatusForgetPassword = "forgetPassword"
const RedisStatusResetEmail = "resetEmail"
