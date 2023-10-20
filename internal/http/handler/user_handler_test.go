// BEGIN: 8f7e2b3c5d4a
package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUser(t *testing.T) {
	app := fiber.New()
	handler := NewUserHandler()
	handler.Register(app)

	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, World!", string(body))
}

// END: 8f7e2b3c5d4a
