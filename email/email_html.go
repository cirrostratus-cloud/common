package email

import (
	"bytes"
	"html/template"
)

func CreateEmailHtml(data interface{}, htmlTemplate string) (string, error) {
	t, err := template.New("email_template").Parse(htmlTemplate)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	if err != nil {
		return "", err
	}
	return tpl.String(), nil
}
