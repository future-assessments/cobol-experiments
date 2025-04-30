package file

import (
	"bufio"
	"fmt"
	"io/fs"
)

type LinesReader struct {
	opener Opener
}

func (reader *LinesReader) GetLinesFrom(filePath string) ([]string, error) {
	textFile, err := reader.opener.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("fail to open file: %w", err)
	}
	defer textFile.Close()

	fileReader := bufio.NewScanner(textFile)

	lines, err := readLines(fileReader)
	if err != nil {
		return nil, fmt.Errorf("fail to read text file: %w", err)
	}

	return lines, nil
}

func readLines(fileReader *bufio.Scanner) ([]string, error) {
	lines := []string{}

	for fileReader.Scan() {
		line := fileReader.Text()
		if fileReader.Err() != nil {
			return nil, fmt.Errorf("fail to read line: %w", fileReader.Err())
		}

		lines = append(lines, line)
	}

	return lines, nil
}

func NewLinesReaderFromFileSytem(fileSystem fs.FS) *LinesReader {
	return &LinesReader{opener: NewFileSytemOpener(fileSystem)}
}

func NewLinesReader() *LinesReader {
	return &LinesReader{opener: NewOsOpener()}
}
