package config

// SecurityConfig security config data
type SecurityConfig struct {
	CSRF struct {
		Secret string `json:"secret"`
	} `json:"csrf"`
	HSTS struct {
		IncludeSubDomains bool `json:"includeSubDomains"`
		MaxAge            int  `json:"maxAge"`
		Preload           bool `json:"preload"`
	} `json:"hsts"`
	Password PasswordConfig `json:"password"`
	Policy   struct {
		Embedder string `json:"embedder"`
		Opener   string `json:"opener"`
		Referrer string `json:"referrer"`
		Resource string `json:"resource"`
	} `json:"policy"`
	Session SessionConfig `json:"session"`
}

func (c *SecurityConfig) Secret() string {
	return c.CSRF.Secret
}

func (c *SecurityConfig) IncludeSubDomains() bool {
	return c.HSTS.IncludeSubDomains
}

func (c *SecurityConfig) MaxAge() int {
	return c.HSTS.MaxAge
}

func (c *SecurityConfig) Preload() bool {
	return c.HSTS.Preload
}

func (c *SecurityConfig) Embedder() string {
	return c.Policy.Embedder
}

func (c *SecurityConfig) Opener() string {
	return c.Policy.Opener
}

func (c *SecurityConfig) Referrer() string {
	return c.Policy.Referrer
}

func (c *SecurityConfig) Resource() string {
	return c.Policy.Resource
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
