package usecase

import (
	"log/slog"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/raphaelbh/devex-api/internal/commons/encryption"
	"github.com/segmentio/ksuid"
)

type TriggerPipelineCommand struct {
	PipelineID string
	Input      map[string]string
}

func TriggerPipeline(command TriggerPipelineCommand) (string, error) {
	var executionID = ksuid.New().String()
	var pipeline, _ = pipelineRepository.FindById(command.PipelineID)
	logger.Info("Pipeline recovered")

	if err := createContainer(executionID, pipeline.Definition.Image); err != nil {
		logger.Error("Unable to create container", slog.String("error", err.Error()))
	}
	defer removeContainer(executionID)

	var secretsCache = map[string]string{}
	for _, step := range pipeline.Definition.Steps {

		var commandExec = step.Command
		logger.Info("Executing command", slog.String("command", commandExec))

		re := regexp.MustCompile(`\${{\s*.*?\s*}}`)
		var matches = re.FindAllString(step.Command, -1)
		for _, variable := range matches {

			reg := regexp.MustCompile(`\${{\s*(.+?)\s*}}`)
			matchess := reg.FindStringSubmatch(variable)

			var varType = strings.Split(matchess[1], ".")[0]
			var varKey = strings.Split(matchess[1], ".")[1]

			var toReplace = ""
			if varType == "input" {
				toReplace = command.Input[varKey]
			}

			if varType == "secrets" {
				if _, exists := secretsCache[varKey]; !exists {
					secretsCache[varKey] = getSecretValue(varKey)
				}
				toReplace = secretsCache[varKey]
			}

			commandExec = strings.ReplaceAll(commandExec, variable, toReplace)
			logger.Info("Command updated")
		}

		if err := executeCommand(executionID, commandExec); err != nil {
			logger.Error("Unable to execute command in container", slog.String("error", err.Error()))
		}
	}

	return executionID, nil
}

func createContainer(containerName, image string) error {
	cmd := exec.Command("docker", "run", "-d",
		"--name", containerName,
		"-v", "/var/run/docker.sock:/var/run/docker.sock",
		"-w", "/workspace",
		image, "sleep", "infinity")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func executeCommand(containerName, command string) error {
	cmd := exec.Command("docker", "exec", containerName, "sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func removeContainer(containerName string) {
	exec.Command("docker", "rm", "-f", containerName).Run()
	logger.Info("Container removed")
}

func getSecretValue(key string) string {
	var secret, _ = secretRepository.FindByKey(key)
	var secretValue, err = encryption.Decrypt(secret.Value)
	if err != nil {
		logger.Error("Unable to decrypt secret", slog.String("error", err.Error()))
	}
	return secretValue
}
