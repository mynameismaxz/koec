package main

import (
	"fmt"
	"html/template"
	"net/http"
	"slices"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
	"github.com/mynameismaxz/koec/templates"
)

const (
	Version  = "0.0.1"
	Priority = 1
)

func main() {
	server.StartServer(New, Version, Priority)
}

type Config struct {
	ResponseCode []int `json:"response_code"`
}

func New() interface{} {
	return &Config{}
}

const (
	ERROR_INDEX_NOT_FOUND = -1
)

func (c *Config) Access(kong *pdk.PDK) {
	// get response code from upstream
	respCode, err := kong.Response.GetStatus()
	if err != nil {
		kong.Log.Err(err.Error())
		return
	}
	traceId, err := kong.Response.GetHeader("X-Kong-Request-Id")
	if err != nil {
		kong.Log.Err(err.Error())
		return
	}

	idx := slices.IndexFunc(c.ResponseCode, func(i int) bool {
		return i == respCode
	})

	if idx != ERROR_INDEX_NOT_FOUND {
		tmpl, err := template.New("error").Parse(templates.ErrorPageLayout)
		if err != nil {
			kong.Log.Err(err.Error())
			return
		}

		bodyResp := &templates.TemplatePayload{
			Title:      http.StatusText(respCode),
			Message:    fmt.Sprintf("Oh no! Something went wrong. Error code: %d", respCode),
			TraceId:    traceId,
			StatusCode: respCode,
		}

		body, err := bodyResp.ToBytes(tmpl)
		if err != nil {
			kong.Log.Err(err.Error())
			return
		}

		kong.Response.ClearHeader("Content-Type")
		kong.Response.SetHeader("Content-Type", "text/html; charset=utf-8")
		kong.Response.Exit(respCode, body, nil)
	} else {
		kong.Log.Err(fmt.Sprintf("Response code %d is not in the list", respCode))
		return
	}
}
