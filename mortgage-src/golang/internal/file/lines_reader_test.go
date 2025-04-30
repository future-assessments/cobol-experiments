package file_test

import (
	"testing"

	"mortgage/internal/file"
	utils_test "mortgage/tests/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLinesReader(t *testing.T) {
	t.Parallel()

	t.Run("should return error when file does not exist", func(t *testing.T) {
		// Arrange
		filePath := "examples/does-not-exist.txt"
		linesReader := file.NewLinesReaderFromFileSytem(utils_test.ExampleFilesFS)
		// Act
		_, err := linesReader.GetLinesFrom(filePath)
		// Assert
		assert.EqualError(
			t,
			err,
			"fail to open file: open examples/does-not-exist.txt: "+
				"file does not exist",
		)
	})

	t.Run("should not return error when file exists", func(t *testing.T) {
		// Arrange
		filePath := "examples/two-examples.txt"
		linesReader := file.NewLinesReaderFromFileSytem(utils_test.ExampleFilesFS)
		// Act
		_, err := linesReader.GetLinesFrom(filePath)
		// Assert
		require.NoError(t, err)
	})

	t.Run("should not return error when file exists and is empty", func(t *testing.T) {
		// Arrange
		filePath := "examples/empty.txt"
		linesReader := file.NewLinesReaderFromFileSytem(utils_test.ExampleFilesFS)
		// Act
		lines, err := linesReader.GetLinesFrom(filePath)
		// Assert
		require.NoError(t, err)
		assert.Empty(t, lines)
	})

	t.Run("should return list of strings when file exists and is not empty", func(t *testing.T) {
		// Arrange
		filePath := "examples/two-examples.txt"
		linesReader := file.NewLinesReaderFromFileSytem(utils_test.ExampleFilesFS)
		// Act
		lines, err := linesReader.GetLinesFrom(filePath)
		// Assert
		require.NoError(t, err)
		assert.Equal(t, 2, len(lines))
	})
}
