# fibercaptcha

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Gihub Action: testing](https://github.com/ssoda/fibercaptcha/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/ssoda/fibercaptcha/actions)

captcha middleware for [Fiber](https://github.com/gofiber/fiber).

idea from [dchest/captcha](https://github.com/dchest/captcha) and the forked repo [LyricTian/captcha](https://github.com/LyricTian/captcha)

## Signatures
```
func New(config ...*fibercaptcha.Config) fiber.Handler
```

## Config

| Property                | Type                   | Description                                                                            | Default                                 |
|:------------------------|:-----------------------|:---------------------------------------------------------------------------------------|:----------------------------------------|
| DefaultLen              | `int`                  | Default number of digits in captcha solution.                                          | `6`                                     |
| CollectNum              | `int`                  | The number of captchas created that triggers garbage collection used by default store. | `100`                                   |
| Expiration              | `time.Duration`        | Expiration time of captchas used by default store.                                     | `10 * time.Minute`                      |
| StdWidth                | `int`                  | Standard width of a captcha image.                                                     | `240`                                   |
| StdHeight               | `int`                  | Standard height of a captcha image.                                                    | `80`                                    |
| RetrieveCaptchaIDPath   | `string`               | API path for retrieve captcha id.                                                      | `/api/captcha/retrieve-id`              |
| ResolveCaptchaImagePath | `string`               | API path for resolve captcha image.                                                    | `/api/captcha/resolve-image`            |
| ResolveCaptchaAudioPath | `string`               | API path for resolve captcha audio.                                                    | `/api/captcha/resolve-audio`            |
| Logger                  | `*log.Logger`          | logger                                                                                 | `log.New(os.Stderr, "", log.LstdFlags)` |
| RedisClient             | `*redis.Client`        | redis client                                                                           |                                         |
| RedisClusterClient      | `*redis.ClusterClient` | redis cluster client                                                                   |                                         |
| RedisCaptchaPrefix      | `string`               | redis captcha key prefix                                                               | `captcha`                               |

## Examples

Installation
```
go get -u github.com/ssoda/fibercaptcha
```

Import package
```
import (
    github.com/ssoda/fibercaptcha
)
```

Use default config
```
app.Use(fibercaptcha.New(&fibercaptcha.Config{}))
```

Call the default retrieve captcha id route to get captcha id
```
/api/captcha/retrieve-id
```

Call the resolve captcha route with captcha id (assume that the captcha id is `CAPTCHA_ID`)
```
/api/captcha/resolve-image?captcha_id=CAPTCHA_ID
```

When you need to reload the image
```
/api/captcha/resolve-image?captcha_id=CAPTCHA_ID&reload=true
```

Verify the captcha digit in your login API or something auth function (assume that digit is `123456`)
```
isValid := fibercaptcha.VerifyString("CAPTCHA_ID", "123456")
if !isValid {
    return ...
}
```
