package fibercaptcha

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssoda/captcha"
)

type retrieveCaptchaIDOutput struct {
	CaptchaID string `json:"captcha_id"`
}

func New(config ...*Config) fiber.Handler {
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {

		if c.Path() == cfg.RetrieveCaptchaIDPath {
			return retrieveCaptchaID(c, cfg.DefaultLen)
		}
		if c.Path() == cfg.ResolveCaptchaPath {
			return resolveCaptcha(c)
		}

		return c.Next()
	}
}

func retrieveCaptchaID(c *fiber.Ctx, captchaLen int) error {
	captchID := captcha.NewLen(captchaLen)
	return c.JSON(retrieveCaptchaIDOutput{
		CaptchaID: captchID,
	})
}

func resolveCaptcha(c *fiber.Ctx) error {
	return c.Next()
}
