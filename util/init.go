package util

import "PushSystem/config"

var secret string

func init() {
	secret = config.Conf.JwtSecret
}
