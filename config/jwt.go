package config

import "github.com/kratos2377/clubhouse/utils"

type JwtConfig struct {
	Secret []byte
}

func NewJwtConfig() *JwtConfig {
	return &JwtConfig{
		Secret: []byte(utils.GetIni("jwt_secret", "JWT_SECRET", "ja9fahf0asdahdaofasf")),
	}
}
