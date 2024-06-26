package fibercaptcha

import (
	"encoding/json"
	"io"
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

		body1, _ := io.ReadAll(resp1.Body)
		var retrieveCaptchaIDOutput retrieveCaptchaIDOutput
		unmarshalErr := json.Unmarshal(body1, &retrieveCaptchaIDOutput)
		if unmarshalErr != nil {
			t.Fatal("cannot unmarshal retrieve captcha id output", unmarshalErr)
		}

		r2 := httptest.NewRequest(fiber.MethodGet, ConfigDefault.ResolveCaptchaImagePath+"?captcha_id="+retrieveCaptchaIDOutput.CaptchaID, nil)
		resp2, _ := app.Test(r2, -1)
		require.Equal(t, fiber.StatusOK, resp2.StatusCode)

		r3 := httptest.NewRequest(fiber.MethodGet, ConfigDefault.ResolveCaptchaAudioPath+"?captcha_id="+retrieveCaptchaIDOutput.CaptchaID, nil)
		resp3, _ := app.Test(r3, -1)
		require.Equal(t, fiber.StatusOK, resp3.StatusCode)
	})

	t.Run("Check resolve captcha id required", func(t *testing.T) {
		app := fiber.New()

		cfg := Config{}
		app.Use(New(&cfg))

		r1 := httptest.NewRequest(fiber.MethodGet, ConfigDefault.RetrieveCaptchaIDPath, nil)
		resp1, _ := app.Test(r1, -1)
		require.Equal(t, fiber.StatusOK, resp1.StatusCode)

		body1, _ := io.ReadAll(resp1.Body)
		var retrieveCaptchaIDOutput retrieveCaptchaIDOutput
		unmarshalErr := json.Unmarshal(body1, &retrieveCaptchaIDOutput)
		if unmarshalErr != nil {
			t.Fatal("cannot unmarshal retrieve captcha id output", unmarshalErr)
		}

		r2 := httptest.NewRequest(fiber.MethodGet, ConfigDefault.ResolveCaptchaImagePath, nil)
		resp2, _ := app.Test(r2, -1)
		require.Equal(t, fiber.StatusBadRequest, resp2.StatusCode)

		r3 := httptest.NewRequest(fiber.MethodGet, ConfigDefault.ResolveCaptchaAudioPath, nil)
		resp3, _ := app.Test(r3, -1)
		require.Equal(t, fiber.StatusBadRequest, resp3.StatusCode)
	})
}
