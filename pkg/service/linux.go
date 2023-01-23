package service

import "fmt"

type LinuxService struct{}

func (s LinuxService) RunShell(cmd string) string {
	fmt.Printf("Running command with bash: %s", cmd)
	return "result"
}

func (s LinuxService) Drives() []string {
	fmt.Println("listing linux drives")
	return []string{"drive1", "drive2"}
}
