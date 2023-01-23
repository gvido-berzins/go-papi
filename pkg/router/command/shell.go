package command

import (
	"go-papi/pkg/config"
	"go-papi/pkg/storage"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// CommandRequest represents a request to run a command sent by a user.
type CommandRequest struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}

func createShellHandler(cfg *config.Conf) func(*gin.Context) {
	return func(c *gin.Context) {
		log.Debug().Msg("called shell handler")
		var body CommandRequest
		if err := c.BindJSON(&body); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		sessionID := uuid.New().String()
		go handleRunningCommand(cfg, body, sessionID)
		c.JSON(http.StatusCreated, gin.H{"session_id": sessionID})
	}
}

func handleRunningCommand(cfg *config.Conf, req CommandRequest, sessionID string) {
	command := &storage.Command{
		Program:   req.Name,
		Args:      strings.Join(req.Args, " "),
		SessionID: sessionID,
		Status:    storage.StatusRunning,
	}

	cmd := exec.Command(req.Name, req.Args...)
	stdin, _ := cmd.StdinPipe()
	stdin.Close()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		command.Status = storage.StatusFailed
		command.TimeFinished = time.Now()
		cfg.DB.Create(command)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		command.Status = storage.StatusFailed
		command.TimeFinished = time.Now()
		cfg.DB.Create(command)
		return
	}

	if err := cmd.Start(); err != nil {
		command.Status = storage.StatusFailed
		command.TimeFinished = time.Now()
		cfg.DB.Create(command)
		return
	}

	command.Pid = cmd.Process.Pid
	cfg.DB.Create(command)
	outBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		cfg.DB.Model(command).Updates(storage.Command{Status: storage.StatusFailed, TimeFinished: time.Now()})
		return
	}
	errBytes, err := ioutil.ReadAll(stderr)
	if err != nil {
		cfg.DB.Model(command).Updates(storage.Command{Status: storage.StatusFailed, TimeFinished: time.Now()})
		return
	}

	if err := cmd.Wait(); err != nil {
		cfg.DB.Model(command).Updates(storage.Command{
			Status:       storage.StatusFailed,
			TimeFinished: time.Now(),
			Stdout:       string(outBytes),
			Stderr:       string(errBytes)})
		return
	}

	exitCode := cmd.ProcessState.ExitCode()
	var status string
	if exitCode != 0 {
		status = storage.StatusFailed
	} else {
		status = storage.StatusSuccess
	}
	cfg.DB.Model(command).Updates(storage.Command{
		Status:       status,
		ExitCode:     exitCode,
		TimeFinished: time.Now(),
		Stdout:       string(outBytes),
		Stderr:       string(errBytes)})
}

func createShellSessionGetHandler(cfg *config.Conf) func(*gin.Context) {
	return func(c *gin.Context) {
		log.Debug().Msg("called shell session get handler")
		sessionID := c.Param("sessionId")
		var command storage.Command
		if result := cfg.DB.First(&command, "session_id = ?", sessionID); result.Error != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusOK, command)
	}
}
