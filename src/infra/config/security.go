package config

// SecurityConfig security config data
type SecurityConfig struct {
	Pepper  string        `json:"pepper"`
	Session SessionConfig `json:"session"`
}

type SessionConfig struct {
	AuthKey   string `json:"auth_key"`
	CryptoKey string `json:"crypto_key"`
	TTL       int    `json:"ttl"`
}
