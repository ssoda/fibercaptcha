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
	// API path for retrieve captcha id
	RetrieveCaptchaIDPath string
	// API path for resolve captcha
	ResolveCaptchaPath string
}

var ConfigDefault = &Config{
	DefaultLen:            6,
	CollectNum:            100,
	Expiration:            10 * time.Minute,
	StdWidth:              240,
	StdHeight:             80,
	RetrieveCaptchaIDPath: "/api/captcha/retrieve-id",
	ResolveCaptchaPath:    "/api/captcha/resolve",
}

func configDefault(config ...*Config) *Config {
	// Return default config if nothing provided
	if len(config) == 0 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	if cfg.DefaultLen == 0 {
		cfg.DefaultLen = ConfigDefault.DefaultLen
	}

	if cfg.CollectNum == 0 {
		cfg.CollectNum = ConfigDefault.CollectNum
	}

	if cfg.Expiration.Minutes() == 0 {
		cfg.Expiration = ConfigDefault.Expiration
	}

	if cfg.StdHeight == 0 {
		cfg.StdHeight = ConfigDefault.StdHeight
	}

	if cfg.StdWidth == 0 {
		cfg.StdWidth = ConfigDefault.StdWidth
	}

	if cfg.RetrieveCaptchaIDPath == "" {
		cfg.RetrieveCaptchaIDPath = ConfigDefault.RetrieveCaptchaIDPath
	}

	if cfg.ResolveCaptchaPath == "" {
		cfg.ResolveCaptchaPath = ConfigDefault.ResolveCaptchaPath
	}

	return cfg
}
