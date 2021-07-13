package config

import (
	"log"

	"github.com/felipeguilhermefs/selene/services"
)

// SELENE_PW_MIN_LEN: Password min lenght.
// SELENE_PW_PEPPER: Password pepper.
func loadPasswordConfig() (services.PasswordConfig, error) {
	minLen, err := getEnvInt("SELENE_PW_MIN_LEN", 8)
	if err != nil {
		return services.PasswordConfig{}, err
	}

	pepper := getEnv("SELENE_PW_PEPPER", "PepperWith64Chars...............................................")

	log.Println(pepper)

	return services.PasswordConfig{
		Pepper: pepper,
		MinLen: minLen,
	}, nil
}
