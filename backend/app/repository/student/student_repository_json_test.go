package student

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"sample/app/model"

	"github.com/stretchr/testify/assert"
)

// createMockStudentJsonFile creates a mock student JSON file and returns its file path
func createMockStudentJsonFile(t *testing.T, students []model.Student) string {
	t.Helper() // Marks this as a test helper

	// Create a sample data at data/minidata.json
	dataDir := "data"
	filePath := filepath.Join(dataDir, "minidata.json") // filePath = data/minidata.json

	err := os.MkdirAll(dataDir, os.ModePerm) // create data directory with permission
	assert.NoError(t, err)

	file, err := os.Create(filePath) // create file
	assert.NoError(t, err)

	err = json.NewEncoder(file).Encode(students)
	assert.NoError(t, err)

	file.Close()

	return filePath
}

// cleanUpMockDataDir deletes the mock data directory after testing
func cleanUpMockDataDir(t *testing.T, dataDir string) {
	t.Helper()
	err := os.RemoveAll(dataDir)
	assert.NoError(t, err)
}

// ✅ Happy path test: file exists and has valid student data
func TestJsonStudentRepository_GetStudents_Success(t *testing.T) {
	mockStudents := []model.Student{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Dob: "2000-01-01", Gpa: 3.8},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Dob: "2000-02-02", Gpa: 3.9},
	}

	dataDir := "data"
	createMockStudentJsonFile(t, mockStudents)
	defer cleanUpMockDataDir(t, dataDir)

	repo := NewJsonStudentRepository()
	students, err := repo.GetStudents()

	assert.NoError(t, err)
	assert.Len(t, students, 2)
	assert.Equal(t, "Alice", students[0].Name)
	assert.Equal(t, "Bob", students[1].Name)
}

// ❌ Failure case test: file does not exist
func TestJsonStudentRepository_GetStudents_FileNotFound(t *testing.T) {
	// Remove file to simulate "not found"
	filePath := "data/minidata.json"
	_ = os.Remove(filePath) // ignore error if file doesn't exist

	repo := NewJsonStudentRepository()
	students, err := repo.GetStudents()

	assert.Error(t, err)
	assert.Empty(t, students)
}
