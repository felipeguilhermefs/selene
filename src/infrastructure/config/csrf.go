package config

import "github.com/felipeguilhermefs/selene/infrastructure/middleware/csrf"

//SELENE_CSRF_SECRET secret to secure cookie
func loadCSRFConfig() csrf.Config {
	secret := getEnv("SELENE_CSRF_SECRET", "SecretWith32Chars...............")

	return csrf.Config{
		Secret: secret,
	}
}
