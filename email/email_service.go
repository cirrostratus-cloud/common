package email

type EmailService interface {
	SendEmail(from string, to string, subject string, body string) error
}
