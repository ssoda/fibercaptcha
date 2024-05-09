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
	// API path for get captcha id
	GetCaptchaIDPath string
	// API path for resolve captcha
	ResolveCaptchaPath string
	// API path for reload captcha
	ReloadCaptchaPath string
}

var ConfigDefault = Config{
	DefaultLen: 6,
	CollectNum: 100,
	Expiration: 10 * time.Minute,
	StdWidth:   240,
	StdHeight:  80,
}
