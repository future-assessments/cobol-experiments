package file

import (
	"io/fs"
	"os"
	"path/filepath"
)

type Opener interface {
	Open(filePath string) (fs.File, error)
}

type OpenWrapper struct {
	fileSystem fs.FS
}

func (opener *OpenWrapper) Open(filePath string) (fs.File, error) {
	cleanedFilePath := filepath.Clean(filePath)

	if opener.fileSystem == nil {
		const readOnlyFileMode fs.FileMode = 0o400

		return os.OpenFile(cleanedFilePath, os.O_RDONLY, readOnlyFileMode)
	}

	return opener.fileSystem.Open(cleanedFilePath)
}

func NewFileSytemOpener(fileSystem fs.FS) *OpenWrapper {
	return &OpenWrapper{fileSystem: fileSystem}
}

func NewOsOpener() *OpenWrapper {
	return &OpenWrapper{fileSystem: nil}
}
