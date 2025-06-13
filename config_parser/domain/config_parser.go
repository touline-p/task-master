package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	config_mod "github.com/touline-p/task-master/core/settings"
	supervisor_models "github.com/touline-p/task-master/supervisor/domain/models"
)

type JobConfigAsJson struct {
	Name          string            `json:"name"`
	Command       string            `json:"cmd"`
	NumProcs      int               `json:"numprocs"`
	Umask         string            `json:"umask"`
	WorkingDir    string            `json:"workingdir"`
	AutoStart     bool              `json:"autostart"`
	AutoRestart   string            `json:"autorestart"`
	ExitCodes     []int             `json:"exitcodes"`
	StartRetries  int               `json:"startretries"`
	StartDuration int               `json:"starttime"`
	StopSignal    string            `json:"stopsignal"`
	StopDuration  int               `json:"stoptime"`
	Stdout        string            `json:"stdout"`
	Stderr        string            `json:"stderr"`
	Environment   map[string]string `json:"env"`
	ConfigValues  []string          `json:"ConfigValues"`
}

type ConfigAsJson struct {
	Jobs []JobConfigAsJson `json:"programs"`
}

func LoadConfig() (*[]supervisor_models.JobConfig, error) {
	var jsonConfig ConfigAsJson
	filename := config_mod.CONFIG_PATH
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file '%s': %w", filename, err)
	}
	err = json.Unmarshal(data, &jsonConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON from '%s': %w", filename, err)
	}
	print_config(jsonConfig)

	return translateConfig(jsonConfig)
}

func translateConfig(jsonconfig ConfigAsJson) (*[]supervisor_models.JobConfig, error) {
	array := make([]supervisor_models.JobConfig, len(jsonconfig.Jobs))
	for index, Job := range jsonconfig.Jobs {
		value, err := translateIndividualConfig(Job)
		if err != nil {
			return nil, err
		}
		array[index] = *value
	}
	return &array, nil
}

func translateIndividualConfig(job JobConfigAsJson) (*supervisor_models.JobConfig, error) {
	newJobConfigValues := make([]supervisor_models.JobConfigValue, 0, 1)

	parsed_int, err := strconv.ParseInt(job.Umask, 8, 32)
	if err != nil {
		return nil, err
	}
	umask := os.FileMode(parsed_int)
	value, err := strToRestartPolicy(job.AutoRestart)
	if err != nil {
		return nil, err
	}
	restartPolicy := value
	startDuration := time.Duration(job.StartDuration)
	stopDuration := time.Duration(job.StartDuration)

	return &supervisor_models.JobConfig{
		Name:          job.Name,
		Command:       job.Command,
		NumProcs:      job.NumProcs,
		Umask:         umask,
		WorkingDir:    job.WorkingDir,
		AutoStart:     job.AutoStart,
		AutoRestart:   restartPolicy,
		ExitCodes:     job.ExitCodes,
		StartRetries:  job.StartRetries,
		StartDuration: startDuration,
		StopSignal:    job.StopSignal,
		StopDuration:  stopDuration,
		Stdout:        job.Stdout,
		Stderr:        job.Stderr,
		Environment:   job.Environment,
		ConfigValues:  newJobConfigValues,
	}, nil
}

func strToRestartPolicy(str string) (supervisor_models.RestartPolicy, error) {
	switch str {
	case string(supervisor_models.RestartAlways):
		return supervisor_models.RestartAlways, nil
	case string(supervisor_models.RestartNever):
		return supervisor_models.RestartNever, nil
	case string(supervisor_models.RestartUnexpected):
		return supervisor_models.RestartUnexpected, nil
	default:
		return "badly", errors.New("str is not a restart policy")

	}
}

func print_config(config any) {
	fmt.Println("=== Configuration ===")

	if cfg, ok := config.(*ConfigAsJson); ok {
		fmt.Printf("Total programs: %d\n\n", len(cfg.Jobs))

		for i, job := range cfg.Jobs {
			fmt.Printf("Program #%d:\n", i+1)
			fmt.Printf("  Name: %s\n", job.Name)
			fmt.Printf("  Command: %s\n", job.Command)
			fmt.Printf("  Number of processes: %d\n", job.NumProcs)
			fmt.Printf("  Umask: %s\n", job.Umask)
			fmt.Printf("  Working directory: %s\n", job.WorkingDir)
			fmt.Printf("  Auto start: %t\n", job.AutoStart)
			fmt.Printf("  Auto restart: %s\n", job.AutoRestart)
			fmt.Printf("  Exit codes: %v\n", job.ExitCodes)
			fmt.Printf("  Start retries: %d\n", job.StartRetries)
			fmt.Printf("  Start duration: %d\n", job.StartDuration)
			fmt.Printf("  Stop signal: %s\n", job.StopSignal)
			fmt.Printf("  Stop duration: %d\n", job.StopDuration)
			fmt.Printf("  Stdout: %s\n", job.Stdout)
			fmt.Printf("  Stderr: %s\n", job.Stderr)

			if len(job.Environment) > 0 {
				fmt.Printf("  Environment variables:\n")
				for key, value := range job.Environment {
					fmt.Printf("    %s=%s\n", key, value)
				}
			} else {
				fmt.Printf("  Environment variables: none\n")
			}

			if len(job.ConfigValues) > 0 {
				fmt.Printf("  Config values: %v\n", job.ConfigValues)
			} else {
				fmt.Printf("  Config values: none\n")
			}

			fmt.Println()
		}
	} else {
		fmt.Println("Configuration (JSON format):")
		if jsonData, err := json.MarshalIndent(config, "", "  "); err == nil {
			fmt.Println(string(jsonData))
		} else {
			fmt.Printf("Error formatting config: %v\n", err)
		}
	}
}
