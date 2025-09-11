package models

import (
	"encoding/json"
	"io"
)

type InputFileUpload struct {
	Filename string
	Data     io.Reader
}

func (*InputFileUpload) isInputFile() {}

func (i *InputFileUpload) MarshalJSON() ([]byte, error) {
	return []byte(`"@` + i.Filename + `"`), nil
}

type InputFileString struct {
	Data string
}

func (*InputFileString) isInputFile() {}

func (i *InputFileString) MarshalJSON() ([]byte, error) {
	return []byte(`"` + i.Data + `"`), nil
}

func (i *InputFileString) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &i.Data)
}
