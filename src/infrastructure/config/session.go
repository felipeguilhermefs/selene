package config

import "github.com/felipeguilhermefs/selene/infrastructure/session"

// SELENE_SESSION_AUTH_KEY: Session Authentication Key.
// SELENE_SESSION_CRYPTO_KEY: Session Encryption Key.
// SELENE_SESSION_TTL: Session max lifetime.
func loadSessionConfig() (session.Config, error) {
	authKey := getEnv("SELENE_SESSION_AUTH_KEY", "AuthKeyWith64Chars..............................................")

	cryptoKey := getEnv("SELENE_SESSION_CRYPTO_KEY", "CryptoKeyWith32Chars............")

	ttl, err := getEnvInt("SELENE_SESSION_TTL", 900)
	if err != nil {
		return session.Config{}, err
	}

	return session.Config{
		AuthenticationKey: authKey,
		EncryptionKey:     cryptoKey,
		TimeToLive:        ttl,
	}, nil
}
