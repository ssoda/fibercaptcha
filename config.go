package fibercaptcha

import (
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ssoda/captcha"
	"github.com/ssoda/captcha/store"
)

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
	// API path for resolve captcha image
	ResolveCaptchaImagePath string
	// API path for resolve captcha audio
	ResolveCaptchaAudioPath string
	// logger
	Logger *log.Logger
	// redis client
	RedisClient *redis.Client
	// redis cluster client
	RedisClusterClient *redis.ClusterClient
	// redis captcha key prefix
	RedisCaptchaPrefix string
}

var ConfigDefault = &Config{
	DefaultLen:              6,
	CollectNum:              100,
	Expiration:              10 * time.Minute,
	StdWidth:                240,
	StdHeight:               80,
	RetrieveCaptchaIDPath:   "/api/captcha/retrieve-id",
	ResolveCaptchaImagePath: "/api/captcha/resolve-image",
	ResolveCaptchaAudioPath: "/api/captcha/resolve-audio",
	Logger:                  log.New(os.Stderr, "", log.LstdFlags),
	RedisCaptchaPrefix:      "captcha",
}

func configDefault(config ...*Config) *Config {
	// Return default config if nothing provided
	if len(config) == 0 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	if cfg.DefaultLen <= 0 {
		cfg.DefaultLen = ConfigDefault.DefaultLen
	}

	if cfg.CollectNum <= 0 {
		cfg.CollectNum = ConfigDefault.CollectNum
	}

	if cfg.Expiration.Minutes() == 0 {
		cfg.Expiration = ConfigDefault.Expiration
	}

	if cfg.StdHeight <= 0 {
		cfg.StdHeight = ConfigDefault.StdHeight
	}

	if cfg.StdWidth <= 0 {
		cfg.StdWidth = ConfigDefault.StdWidth
	}

	if cfg.RetrieveCaptchaIDPath == "" {
		cfg.RetrieveCaptchaIDPath = ConfigDefault.RetrieveCaptchaIDPath
	}

	if cfg.ResolveCaptchaImagePath == "" {
		cfg.ResolveCaptchaImagePath = ConfigDefault.ResolveCaptchaImagePath
	}

	if cfg.ResolveCaptchaAudioPath == "" {
		cfg.ResolveCaptchaAudioPath = ConfigDefault.ResolveCaptchaAudioPath
	}

	if cfg.Logger != nil {
		cfg.Logger = ConfigDefault.Logger
	}

	if cfg.RedisCaptchaPrefix == "" {
		cfg.RedisCaptchaPrefix = ConfigDefault.RedisCaptchaPrefix
	}

	if cfg.RedisClient != nil {
		captcha.SetCustomStore(store.NewRedisStoreWithCli(cfg.RedisClient, cfg.Expiration, cfg.Logger, cfg.RedisCaptchaPrefix))
	}

	if cfg.RedisClusterClient != nil {
		captcha.SetCustomStore(store.NewRedisClusterStoreWithCli(cfg.RedisClusterClient, cfg.Expiration, cfg.Logger, cfg.RedisCaptchaPrefix))
	}

	return cfg
}
