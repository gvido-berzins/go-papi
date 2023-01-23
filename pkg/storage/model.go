package storage

import (
	"time"

	"gorm.io/gorm"
)

var (
	StatusRunning = "running"
	StatusFailed  = "failed"
	StatusSuccess = "success"
)

// Command represents a shell command that had been run on the system
type Command struct {
	gorm.Model
	SessionID    string    `gorm:"column:session_id;index;not null"`
	Program      string    `gorm:"column:program;not null"`
	Args         string    `gorm:"column:args"`
	Pid          int       `gorm:"column:pid"`
	Status       string    `gorm:"column:status;not null"`
	ExitCode     int       `gorm:"column:exit_code"`
	Stdout       string    `gorm:"column:stdout"`
	Stderr       string    `gorm:"column:stderr"`
	TimeFinished time.Time `gorm:"column:time_finished"`
}
