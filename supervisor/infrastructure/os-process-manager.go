package infrastructure

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"

	"github.com/touline-p/task-master/supervisor/domain/models"
	"github.com/touline-p/task-master/supervisor/domain/services"
)

type OSProcessManager struct {
	mutex     sync.RWMutex
	processes map[models.JobId]*exec.Cmd
}

func NewOSProcessManager() services.IProcessManager {
	return &OSProcessManager{
		processes: make(map[models.JobId]*exec.Cmd),
	}
}

func (pm *OSProcessManager) Launch(ctx context.Context, job *models.Job) (int, error) {
	cmdParts := strings.Fields(job.Config.Command)
	if len(cmdParts) == 0 {
		return 0, fmt.Errorf("empty command")
	}

	cmd := exec.CommandContext(ctx, cmdParts[0], cmdParts[1:]...)

	if job.Config.WorkingDir != "" {
		cmd.Dir = job.Config.WorkingDir
	}

	cmd.Env = os.Environ()
	for key, value := range job.Config.Environment {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
	}

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	if job.Config.Stdout != "" {
		stdoutFile, err := os.OpenFile(job.Config.Stdout, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return 0, fmt.Errorf("failed to open stdout file: %w", err)
		}
		cmd.Stdout = stdoutFile
	}

	if job.Config.Stderr != "" {
		if job.Config.Stderr == job.Config.Stdout {
			cmd.Stderr = cmd.Stdout
		} else {
			stderrFile, err := os.OpenFile(job.Config.Stderr, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				return 0, fmt.Errorf("failed to open stderr file: %w", err)
			}
			cmd.Stderr = stderrFile
		}
	}

	if err := cmd.Start(); err != nil {
		return 0, fmt.Errorf("failed to start process: %w", err)
	}

	pm.mutex.Lock()
	pm.processes[job.Id] = cmd
	pm.mutex.Unlock()

	return cmd.Process.Pid, nil
}

func (pm *OSProcessManager) Signal(job *models.Job, sig os.Signal) error {
	process, err := os.FindProcess(job.Pid)
	if err != nil {
		return err
	}
	return process.Signal(sig)
}

func (pm *OSProcessManager) Terminate(job *models.Job) error {
	return pm.Signal(job, syscall.SIGTERM)
}

func (pm *OSProcessManager) Kill(job *models.Job) error {
	process, err := os.FindProcess(job.Pid)
	if err != nil {
		return err
	}
	return process.Kill()
}

func (pm *OSProcessManager) WaitForExit(jobId models.JobId) (int, error, bool) {
	pm.mutex.RLock()
	cmd, exists := pm.processes[jobId]
	pm.mutex.RUnlock()

	if !exists {
		return 0, fmt.Errorf("process not found for job %s", jobId), false
	}

	err := cmd.Wait()

	pm.mutex.Lock()
	delete(pm.processes, jobId)
	pm.mutex.Unlock()

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus(), nil, true
			}
		}
		return 0, err, false
	}

	return 0, nil, true
}
