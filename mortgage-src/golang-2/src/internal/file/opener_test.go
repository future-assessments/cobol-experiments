package file_test

import (
	"os"
	"testing"

	"mortgage/internal/file"
	utils_test "mortgage/tests/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpener(t *testing.T) {
	t.Parallel()

	t.Run("should open file in a given file system", func(t *testing.T) {
		// Arrange
		filePath := "examples/two-examples.txt"
		opener := file.NewFileSytemOpener(utils_test.ExampleFilesFS)
		// Act
		fileOpened, err := opener.Open(filePath)
		// Assert
		require.NoError(t, err)
		err = fileOpened.Close()
		require.NoError(t, err)
	})

	t.Run("should open file using os.OpenFile when no file system is given", func(t *testing.T) {
		// Arrange
		filePath := "opener_test.go"
		opener := file.NewOsOpener()
		// Act
		fileOpened, err := opener.Open(filePath)
		// Assert
		require.NoError(t, err)
		err = fileOpened.Close()
		require.NoError(t, err)
	})

	t.Run("should open file given an absolute path", func(t *testing.T) {
		// Arrange
		testFolder, _ := os.Getwd()
		filePath := testFolder + "/opener_test.go"
		opener := file.NewOsOpener()
		// Act
		fileOpened, err := opener.Open(filePath)
		// Assert
		require.NoError(t, err)
		err = fileOpened.Close()
		require.NoError(t, err)
	})

	t.Run("should fail when file does not exist", func(t *testing.T) {
		// Arrange
		filePath := "does-not-exist.txt"
		opener := file.NewOsOpener()
		// Act
		fileOpened, err := opener.Open(filePath)
		// Assert
		require.Error(t, err)
		assert.Nil(t, fileOpened)
	})
}
