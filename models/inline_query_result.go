package models

import "encoding/json"

func (*InlineQueryResultCachedAudio) isInlineQueryResult() {}
func (m *InlineQueryResultCachedAudio) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultCachedAudio
	}{
		Type:                         "audio",
		InlineQueryResultCachedAudio: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultCachedDocument) isInlineQueryResult() {}
func (m *InlineQueryResultCachedDocument) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultCachedDocument
	}{
		Type:                            "document",
		InlineQueryResultCachedDocument: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultCachedGif) isInlineQueryResult() {}
func (m *InlineQueryResultCachedGif) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultCachedGif
	}{
		Type:                       "gif",
		InlineQueryResultCachedGif: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultCachedMpeg4Gif) isInlineQueryResult() {}
func (m *InlineQueryResultCachedMpeg4Gif) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultCachedMpeg4Gif
	}{
		Type:                            "mpeg4_gif",
		InlineQueryResultCachedMpeg4Gif: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultCachedPhoto) isInlineQueryResult() {}
func (m *InlineQueryResultCachedPhoto) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultCachedPhoto
	}{
		Type:                         "photo",
		InlineQueryResultCachedPhoto: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultCachedSticker) isInlineQueryResult() {}
func (m *InlineQueryResultCachedSticker) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultCachedSticker
	}{
		Type:                           "sticker",
		InlineQueryResultCachedSticker: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultCachedVideo) isInlineQueryResult() {}
func (m *InlineQueryResultCachedVideo) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultCachedVideo
	}{
		Type:                         "video",
		InlineQueryResultCachedVideo: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultCachedVoice) isInlineQueryResult() {}
func (m *InlineQueryResultCachedVoice) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultCachedVoice
	}{
		Type:                         "voice",
		InlineQueryResultCachedVoice: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultArticle) isInlineQueryResult() {}
func (m *InlineQueryResultArticle) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultArticle
	}{
		Type:                     "article",
		InlineQueryResultArticle: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultAudio) isInlineQueryResult() {}
func (m *InlineQueryResultAudio) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultAudio
	}{
		Type:                   "audio",
		InlineQueryResultAudio: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultContact) isInlineQueryResult() {}
func (m *InlineQueryResultContact) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultContact
	}{
		Type:                     "contact",
		InlineQueryResultContact: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultGame) isInlineQueryResult() {}
func (m *InlineQueryResultGame) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultGame
	}{
		Type:                  "game",
		InlineQueryResultGame: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultDocument) isInlineQueryResult() {}
func (m *InlineQueryResultDocument) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultDocument
	}{
		Type:                      "document",
		InlineQueryResultDocument: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultGif) isInlineQueryResult() {}
func (m *InlineQueryResultGif) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultGif
	}{
		Type:                 "gif",
		InlineQueryResultGif: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultLocation) isInlineQueryResult() {}
func (m *InlineQueryResultLocation) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultLocation
	}{
		Type:                      "location",
		InlineQueryResultLocation: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultMpeg4Gif) isInlineQueryResult() {}
func (m *InlineQueryResultMpeg4Gif) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultMpeg4Gif
	}{
		Type:                      "mpeg4_gif",
		InlineQueryResultMpeg4Gif: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultPhoto) isInlineQueryResult() {}
func (m *InlineQueryResultPhoto) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultPhoto
	}{
		Type:                   "photo",
		InlineQueryResultPhoto: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultVenue) isInlineQueryResult() {}
func (m *InlineQueryResultVenue) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultVenue
	}{
		Type:                   "venue",
		InlineQueryResultVenue: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultVideo) isInlineQueryResult() {}
func (m *InlineQueryResultVideo) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultVideo
	}{
		Type:                   "video",
		InlineQueryResultVideo: m,
	}
	return json.Marshal(&data)
}

func (*InlineQueryResultVoice) isInlineQueryResult() {}
func (m *InlineQueryResultVoice) MarshalCustom() ([]byte, error) {
	data := struct {
		Type string `json:"type"`
		*InlineQueryResultVoice
	}{
		Type:                   "voice",
		InlineQueryResultVoice: m,
	}
	return json.Marshal(&data)
}
