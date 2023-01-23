package service

// OSService represents an OS specific interface for running commands
// and getting other info based on how it's done on the OS.
type OSService interface {
	RunShell(cmd string) string
	Drives() []string
}
