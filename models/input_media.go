package models

import (
	"encoding/json"
	"io"
)

func (ima *InputMediaAnimation) GetData() io.Reader {
	return ima.Data
}

func (ima *InputMediaAnimation) GetMedia() string {
	return ima.Media
}

func (ima *InputMediaAnimation) MarshalInputMedia() ([]byte, error) {
	mim := struct {
		Type string `json:"type"`
		*InputMediaAnimation
	}{
		Type:                "animation",
		InputMediaAnimation: ima,
	}

	return json.Marshal(mim)
}

func (InputMediaAnimation) isInputMedia() {}

func (ima *InputMediaDocument) GetData() io.Reader {
	return ima.Data
}

func (ima *InputMediaDocument) GetMedia() string {
	return ima.Media
}

func (imd *InputMediaDocument) MarshalInputMedia() ([]byte, error) {
	mim := struct {
		Type string `json:"type"`
		*InputMediaDocument
	}{
		Type:               "document",
		InputMediaDocument: imd,
	}
	return json.Marshal(mim)
}

func (InputMediaDocument) isInputMedia() {}

func (ima *InputMediaAudio) GetData() io.Reader {
	return ima.Data
}

func (ima *InputMediaAudio) GetMedia() string {
	return ima.Media
}

func (ima *InputMediaAudio) MarshalInputMedia() ([]byte, error) {
	mim := struct {
		Type string `json:"type"`
		*InputMediaAudio
	}{
		Type:            "audio",
		InputMediaAudio: ima,
	}
	return json.Marshal(mim)
}

func (InputMediaAudio) isInputMedia() {}

func (ima *InputMediaPhoto) GetData() io.Reader {
	return ima.Data
}

func (ima *InputMediaPhoto) GetMedia() string {
	return ima.Media
}

func (imp *InputMediaPhoto) MarshalInputMedia() ([]byte, error) {
	mim := struct {
		Type string `json:"type"`
		*InputMediaPhoto
	}{
		Type:            "photo",
		InputMediaPhoto: imp,
	}
	return json.Marshal(mim)
}

func (InputMediaPhoto) isInputMedia() {}

func (ima *InputMediaVideo) GetData() io.Reader {
	return ima.Data
}

func (ima *InputMediaVideo) GetMedia() string {
	return ima.Media
}

func (imv *InputMediaVideo) MarshalInputMedia() ([]byte, error) {
	mim := struct {
		Type string `json:"type"`
		*InputMediaVideo
	}{
		Type:            "video",
		InputMediaVideo: imv,
	}
	return json.Marshal(mim)
}

func (InputMediaVideo) isInputMedia() {}
