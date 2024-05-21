package fibercaptcha

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssoda/captcha"
)

type retrieveCaptchaIDOutput struct {
	CaptchaID string `json:"captcha_id"`
}

type resolveCaptchInput struct {
	CaptchaID string `query:"captcha_id"`
	Reload    string `query:"reload"`
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
	r := new(resolveCaptchInput)
	if err := c.QueryParser(r); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Next()
}
