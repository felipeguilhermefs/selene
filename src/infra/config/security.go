package config

// SecurityConfig security config data
type SecurityConfig struct {
	CSRF     CSRFConfig     `json:"csrf"`
	Password PasswordConfig `json:"password"`
	Session  SessionConfig  `json:"session"`
}

type CSRFConfig struct {
	Sct string `json:"secret"`
}

func (c *CSRFConfig) Secret() string {
	return c.Sct
}

type PasswordConfig struct {
	MinLen int    `json:"min_length"`
	Pepper string `json:"pepper"`
}

type SessionConfig struct {
	AuthKey   string `json:"auth_key"`
	CryptoKey string `json:"crypto_key"`
	TTL       int    `json:"ttl"`
}
