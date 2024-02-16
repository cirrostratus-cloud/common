package email_test

import (
	"testing"

	"github.com/cirrostratus-cloud/common/email"
	"github.com/stretchr/testify/assert"
)

const (
	htmlResult = `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Test Email</title>
		</head>
		<body>
			<h1>Test Email</h1>
			<p>This is a test email</p>
		</body>
		</html>
	`
	htmlTemplate = `
		<!DOCTYPE html>
		<html>
		<head>
			<title>{{.Title}}</title>
		</head>
		<body>
			<h1>{{.Header}}</h1>
			<p>{{.Content}}</p>
		</body>
		</html>
	`
)

type EmailModel struct {
	Title   string
	Header  string
	Content string
}

func TestTemplateOk(t *testing.T) {
	assert := assert.New(t)
	html, err := email.CreateEmailHtml(EmailModel{
		Title:   "Test Email",
		Header:  "Test Email",
		Content: "This is a test email",
	}, htmlTemplate)
	assert.Nil(err)
	assert.Equal(htmlResult, html)
}
