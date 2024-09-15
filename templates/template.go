package templates

import (
	"bytes"
	"html/template"
)

type TemplatePayload struct {
	Title      string
	Message    string
	TraceId    string
	StatusCode int
}

func (t *TemplatePayload) ToBytes(tmpl *template.Template) ([]byte, error) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, t); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
