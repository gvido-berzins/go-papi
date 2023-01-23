package service

import "fmt"

type WindowsService struct{}

func (s WindowsService) RunShell(cmd string) string {
	fmt.Printf("Running command with bash: %s", cmd)
	return "result"
}

func (s WindowsService) Drives() []string {
	fmt.Println("listing windows drives")
	return []string{"drive1", "drive2"}
}
