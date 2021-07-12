package session

type Config interface {
	AuthKey() string
	CryptoKey() string
	TimeToLive() int
}
