package fibercaptcha

import "github.com/gofiber/fiber/v2"

func New(config ...*Config) fiber.Handler {
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {

		if c.Path() == cfg.RetrieveCaptchaIDPath {
			return retrieveCaptchaID(c)
		}
		if c.Path() == cfg.ResolveCaptchaPath {
			return resolveCaptcha(c)
		}

		return c.Next()
	}
}

func retrieveCaptchaID(c *fiber.Ctx) error {
	return c.Next()
}

func resolveCaptcha(c *fiber.Ctx) error {
	return c.Next()
}
