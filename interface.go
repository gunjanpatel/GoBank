package bank

// Declare interface
type Bank interface {
	Statement() string
}

// Implement interface
func Statement(b Bank) string {
	return b.Statement()
}
