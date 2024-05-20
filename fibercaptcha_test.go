package fibercaptcha

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("Endpoint check default path", func(t *testing.T) {
		app := fiber.New()

		cfg := Config{}
		app.Use(New(&cfg))

		r1 := httptest.NewRequest(fiber.MethodGet, ConfigDefault.RetrieveCaptchaIDPath, nil)
		resp1, _ := app.Test(r1, -1)
		require.Equal(t, fiber.StatusOK, resp1.StatusCode)

		r2 := httptest.NewRequest(fiber.MethodGet, ConfigDefault.ResolveCaptchaPath, nil)
		resp2, _ := app.Test(r2, -1)
		require.Equal(t, fiber.StatusBadRequest, resp2.StatusCode)
	})
}
