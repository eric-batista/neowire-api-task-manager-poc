package models_test

import (
	"testing"

	"github.com/eric-batista/neowire-api-task-manager-poc/domain/models"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateNewTask(t *testing.T) {
	title := "Task 0001"
	description := "Implementar task"
	status := "Pending"
	clientId := uuid.NewV4().String()
	task, err := models.NewTask(title, description, status, clientId)

	require.Nil(t, err)
	require.Equal(t, task.Title, title)
	require.Equal(t, task.Description, description)
	require.Equal(t, task.Status, status)

	task, err = models.NewTask("", "", status, clientId)
	require.Nil(t, task)
	require.NotNil(t, err)

	task, err = models.NewTask("", "", "", "")
	require.Nil(t, task)
	require.NotNil(t, err)
}
