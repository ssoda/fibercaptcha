package fibercaptcha

import (
	"fmt"

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
			return resolveCaptcha(c, cfg.StdWidth, cfg.StdHeight)
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

func resolveCaptcha(c *fiber.Ctx, captchaWidth int, captchaHeight int) error {
	input := new(resolveCaptchInput)
	if err := c.QueryParser(input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if input.CaptchaID == "" {
		c.SendString("captcha id required.")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if input.Reload != "" && !captcha.Reload(input.CaptchaID) {
		c.SendString("invalid captcha id.")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := captcha.WriteImage(c.Response().BodyWriter(), input.CaptchaID, captchaWidth, captchaHeight)
	if err != nil {
		c.SendString(fmt.Sprintf("write captcha image failed. err: %v", err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Set(fiber.HeaderCacheControl, "no-cache, no-store, must-revalidate")
	c.Set(fiber.HeaderPragma, "no-cache")
	c.Set(fiber.HeaderExpires, "0")
	c.Set(fiber.HeaderContentType, "image/png")

	return c.SendStatus(fiber.StatusOK)
}
