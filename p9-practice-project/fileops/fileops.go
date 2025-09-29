package fileops

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type Fileop struct {
	InputFilePath string
	OutputFilePath string
}

func New(inpath, outpath string) Fileop {
	return Fileop{
		InputFilePath: inpath,
		OutputFilePath: outpath,
	}
}

func (fo Fileop) ReadLines() ([]string, error) {
	file, err := os.Open(fo.InputFilePath)
	if err != nil {
		return nil, errors.New("ERROR DURING FILE OPENING")
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("ERROR DURING FILE READING")
	}

	file.Close()
	return lines, nil
}

func (fo Fileop) WriteResult(data any) error {
	file, err := os.Create(fo.OutputFilePath)
	if err != nil {
		return errors.New("ERROR DURING FILE CREATING")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("ERROR DURING ENCODING")
	}

	file.Close()
	return nil

}
