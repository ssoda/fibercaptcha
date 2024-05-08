package fibercaptcha

import "time"

type Config struct {
	// Default number of digits in captcha solution.
	DefaultLen int
	// The number of captchas created that triggers garbage collection used
	// by default store.
	CollectNum int
	// Expiration time of captchas used by default store.
	Expiration time.Duration
	// Standard width of a captcha image.
	StdWidth int
	// Standard height of a captcha image.
	StdHeight int
}

const (
	// Default number of digits in captcha solution.
	DefaultLen = 6
	// The number of captchas created that triggers garbage collection used
	// by default store.
	CollectNum = 100
	// Expiration time of captchas used by default store.
	Expiration = 10 * time.Minute
	// Standard width and height of a captcha image.
	StdWidth  = 240
	StdHeight = 80
)
