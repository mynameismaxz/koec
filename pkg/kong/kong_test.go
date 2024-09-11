package kong

import (
	"net/http"
	"testing"

	"github.com/Kong/go-pdk/test"
	"github.com/stretchr/testify/assert"
)

func TestRequestNormal(t *testing.T) {
	/*
		objective: test the plugin with normal case (response code 200)
		and will return 200 with same response header, body
	*/
	t.Run("WithExpectedResponseCode", func(t *testing.T) {
		env, err := test.New(t, test.Request{
			Method: http.MethodGet,
			Url:    "http://example.com/test",
		})
		assert.NoError(t, err)

		// mocking status will be 200
		env.ClientRes.Status = http.StatusOK
		env.DoAccess(&Config{
			ResponseCode: []int{http.StatusNotFound, http.StatusInternalServerError},
		})

		assert.Equal(t, http.StatusOK, env.ClientRes.Status)
	})

	t.Run("WithExpectSameHeader", func(t *testing.T) {
		env, err := test.New(t, test.Request{
			Method: http.MethodGet,
			Url:    "http://example.com/test",
		})
		assert.NoError(t, err)

		// mocking status will be 200
		env.ClientRes.Status = http.StatusOK

		// mocking header
		env.ClientRes.Headers.Add("X-Test", "test")
		env.DoAccess(&Config{
			ResponseCode: []int{http.StatusNotFound, http.StatusInternalServerError},
		})

		assert.Equal(t, http.StatusOK, env.ClientRes.Status)
		assert.Equal(t, "test", env.ClientRes.Headers.Get("X-Test"))
	})

	t.Run("WithExpectSameBody", func(t *testing.T) {
		env, err := test.New(t, test.Request{
			Method: http.MethodGet,
			Url:    "http://example.com/test",
		})
		assert.NoError(t, err)

		// mocking status will be 200
		env.ClientRes.Status = http.StatusOK
		env.ClientRes.Body = []byte("test")
		env.DoAccess(&Config{
			ResponseCode: []int{http.StatusNotFound, http.StatusInternalServerError},
		})

		assert.Equal(t, http.StatusOK, env.ClientRes.Status)
		assert.Equal(t, "test", string(env.ClientRes.Body))
	})
}

func TestRequestError(t *testing.T) {
	/*
		Objective: test the plugin with error case e.g. response code 404, 500
		and will return with error code page with the html style.
	*/
	var responseErrorCode = []int{http.StatusNotFound, http.StatusInternalServerError}
	t.Run("WithExpectedResponseCode", func(t *testing.T) {
		env, err := test.New(t, test.Request{
			Method: http.MethodGet,
		})
		assert.NoError(t, err)

		// mocking status
		env.ClientRes.Status = http.StatusNotFound
		env.DoAccess(&Config{
			ResponseCode: responseErrorCode,
		})

		assert.Equal(t, http.StatusNotFound, env.ClientRes.Status)
		assert.NotEqual(t, "", string(env.ClientRes.Body))
	})

	t.Run("WithExpectSameHeader", func(t *testing.T) {
		env, err := test.New(t, test.Request{
			Method: http.MethodGet,
		})
		assert.NoError(t, err)

		// mocking status
		env.ClientRes.Status = http.StatusNotFound
		env.ClientRes.Headers.Add("X-Test", "test")
		env.DoAccess(&Config{
			ResponseCode: responseErrorCode,
		})

		assert.Equal(t, http.StatusNotFound, env.ClientRes.Status)
		assert.Equal(t, "test", env.ClientRes.Headers.Get("X-Test"))
	})

	t.Run("WithExpectHTMLBody", func(t *testing.T) {
		env, err := test.New(t, test.Request{
			Method: http.MethodGet,
		})
		assert.NoError(t, err)

		// mocking status
		env.ClientRes.Status = http.StatusNotFound
		env.DoAccess(&Config{
			ResponseCode: responseErrorCode,
		})

		assert.Equal(t, http.StatusNotFound, env.ClientRes.Status)
		assert.Contains(t, string(env.ClientRes.Body), "<!DOCTYPE html>")
	})

	t.Run("WithContainsTraceID", func(t *testing.T) {
		env, err := test.New(t, test.Request{
			Method: http.MethodGet,
		})
		assert.NoError(t, err)

		// mocking status
		env.ClientRes.Status = http.StatusNotFound
		env.ClientRes.Headers.Add("X-Kong-Request-Id", "test")
		env.DoAccess(&Config{
			ResponseCode: responseErrorCode,
		})

		assert.Equal(t, http.StatusNotFound, env.ClientRes.Status)
		assert.Contains(t, string(env.ClientRes.Body), "RequestID: test")
	})

	t.Run("WithOutofResponseCodeRange", func(t *testing.T) {
		env, err := test.New(t, test.Request{
			Method: http.MethodGet,
		})
		assert.NoError(t, err)

		// mocking status
		env.ClientRes.Status = http.StatusAccepted
		env.DoAccess(&Config{
			ResponseCode: responseErrorCode,
		})

		assert.Equal(t, http.StatusAccepted, env.ClientRes.Status)
		assert.NotContains(t, string(env.ClientRes.Body), "<!DOCTYPE html>")
	})
}
