package models

import (
	"encoding/json"
	"io"
)

func (imp *InputPaidMediaPhoto) GetData() io.Reader {
	return imp.Data
}

func (imp *InputPaidMediaPhoto) GetMedia() string {
	return imp.Media
}

func (imp *InputPaidMediaPhoto) MarshalInputMedia() ([]byte, error) {
	mim := struct {
		Type string `json:"type"`
		*InputPaidMediaPhoto
	}{
		Type:                "photo",
		InputPaidMediaPhoto: imp,
	}

	return json.Marshal(mim)
}

func (imp *InputPaidMediaPhoto) isInputPaidMedia() {}

func (ipv *InputPaidMediaVideo) GetData() io.Reader {
	return ipv.Data
}

func (ipv *InputPaidMediaVideo) GetMedia() string {
	return ipv.Media
}

func (ipv *InputPaidMediaVideo) MarshalInputMedia() ([]byte, error) {
	mim := struct {
		Type string `json:"type"`
		*InputPaidMediaVideo
	}{
		Type:                "video",
		InputPaidMediaVideo: ipv,
	}

	return json.Marshal(mim)
}

func (ipv *InputPaidMediaVideo) isInputPaidMedia() {}
