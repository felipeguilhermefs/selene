package config

// SecurityConfig security config data
type SecurityConfig struct {
	CSRF     CSRFConfig     `json:"csrf"`
	Password PasswordConfig `json:"password"`
	Policy   PolicyConfig   `json:"policy"`
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

type PolicyConfig struct {
	Emb   string `json:"embedder"`
	Open  string `json:"opener"`
	Refer string `json:"referrer"`
	Res   string `json:"resource"`
}

func (c *PolicyConfig) Embedder() string {
	return c.Emb
}

func (c *PolicyConfig) Opener() string {
	return c.Open
}

func (c *PolicyConfig) Referrer() string {
	return c.Refer
}

func (c *PolicyConfig) Resource() string {
	return c.Res
}

type SessionConfig struct {
	AuthKey   string `json:"auth_key"`
	CryptoKey string `json:"crypto_key"`
	TTL       int    `json:"ttl"`
}
