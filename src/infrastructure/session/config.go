package session

type Config struct {
	AuthenticationKey string
	EncryptionKey     string
	TimeToLive        int
}
