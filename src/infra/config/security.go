package config

// SecurityConfig security config data
type SecurityConfig struct {
	CSP struct {
		BaseURI         string   `json:"baseURI"`
		DefaultSrc      string   `json:"defaultSrc"`
		FormAction      string   `json:"formAction"`
		FrameAncestors  string   `json:"frameAncestors"`
		StyleSrc        []string `json:"styleSrc"`
		ScriptSrc       []string `json:"scriptSrc"`
		UpgradeInsecure bool     `json:"upgradeInsecure"`
	} `json:"csp"`
	CSRF     CSRFConfig     `json:"csrf"`
	HSTS     HSTSConfig     `json:"hsts"`
	Password PasswordConfig `json:"password"`
	Policy   PolicyConfig   `json:"policy"`
	Session  SessionConfig  `json:"session"`
}

func (c *SecurityConfig) BaseURI() string {
	return c.CSP.BaseURI
}

func (c *SecurityConfig) DefaultSrc() string {
	return c.CSP.DefaultSrc
}

func (c *SecurityConfig) FormAction() string {
	return c.CSP.FormAction
}

func (c *SecurityConfig) FrameAncestors() string {
	return c.CSP.FrameAncestors
}

func (c *SecurityConfig) StyleSrc() []string {
	return c.CSP.StyleSrc
}

func (c *SecurityConfig) ScriptSrc() []string {
	return c.CSP.ScriptSrc
}

func (c *SecurityConfig) UpgradeInsecure() bool {
	return c.CSP.UpgradeInsecure
}

type CSRFConfig struct {
	Sct string `json:"secret"`
}

func (c *CSRFConfig) Secret() string {
	return c.Sct
}

type HSTSConfig struct {
	IncSubDomain bool `json:"includeSubDomains"`
	MxAge        int  `json:"maxAge"`
	Preld        bool `json:"preload"`
}

func (c *HSTSConfig) IncludeSubDomains() bool {
	return c.IncSubDomain
}

func (c *HSTSConfig) MaxAge() int {
	return c.MxAge
}

func (c *HSTSConfig) Preload() bool {
	return c.Preld
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
