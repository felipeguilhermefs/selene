package session

type Config interface {
	AuthenticationKey() string
	EncryptionKey() string
	TimeToLive() int
}
