package curriculum

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	filePath := "../../examples/data.yaml"

	cv, err := ParseFile(filePath)

	assert.Nil(t, err)
	assert.Equal(t, "Software Engineer", cv.Job)
}

func TestParseNonExistingFile(t *testing.T) {
	filePath := "incorrect.yaml"

	cv, err := ParseFile(filePath)

	assert.Error(t, err)
	assert.Nil(t, cv)
}

func TestParseInvalidFile(t *testing.T) {
	filePath := os.TempDir() + "data.yaml"
	f, _ := os.Create(filePath)
	f.Write([]byte("invalid file"))

	cv, err := ParseFile(f.Name())

	assert.Error(t, err)
	assert.Nil(t, cv)
	os.Remove(filePath)
}

func TestParseFileWithInvalidExtension(t *testing.T) {
	filePath := os.TempDir() + "data.err"
	f, _ := os.Create(filePath)

	cv, err := ParseFile(f.Name())

	assert.Error(t, err)
	assert.Nil(t, cv)
	os.Remove(filePath)
}
