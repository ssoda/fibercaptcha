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
	Lang      string `query:"lang"`
}

func New(config ...*Config) fiber.Handler {
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {

		if c.Path() == cfg.RetrieveCaptchaIDPath {
			return retrieveCaptchaID(c, cfg.DefaultLen)
		}

		if c.Path() == cfg.ResolveCaptchaImagePath {
			return resolveCaptchaImage(c, cfg.StdWidth, cfg.StdHeight)
		}

		if c.Path() == cfg.ResolveCaptchaAudioPath {
			return resolveCaptchaAudio(c)
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

func resolveCaptchaImage(c *fiber.Ctx, captchaWidth int, captchaHeight int) error {
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

func resolveCaptchaAudio(c *fiber.Ctx) error {
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

	err := captcha.WriteAudio(c.Response().BodyWriter(), input.CaptchaID, input.Lang)
	if err != nil {
		c.SendString(fmt.Sprintf("write captcha image failed. err: %v", err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Set(fiber.HeaderCacheControl, "no-cache, no-store, must-revalidate")
	c.Set(fiber.HeaderPragma, "no-cache")
	c.Set(fiber.HeaderExpires, "0")
	c.Set(fiber.HeaderContentType, "audio/x-wav")

	return c.SendStatus(fiber.StatusOK)
}

// VerifyString accepts a string of digits.  It removes
// spaces and commas from the string, but any other characters, apart from
// digits and listed above, will cause the function to return false.
func VerifyString(id string, digits string) bool {
	return captcha.VerifyString(id, digits)
}
