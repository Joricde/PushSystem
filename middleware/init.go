package middleware

import "awesomeProject/config"

var secret string

func init() {
	secret = config.Conf.JwtSecret

}
