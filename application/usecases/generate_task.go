package usecases

import "github.com/eric-batista/neowire-api-task-manager-poc/domain/models"

func GenerateTaskBasedOnPromt(payload models.GenerateTaskPrompt) models.Task {
	// Adicionar talker para o serviço em python (consumindo GPT-4 ou 3) para a geração de tarefas
	return models.Task{}
}
