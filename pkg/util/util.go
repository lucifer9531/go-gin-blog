package util

import "github.com/ruixijiejie/go-gin-blog/pkg/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
