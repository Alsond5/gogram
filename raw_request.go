package gogram

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type apiResponse[T any] struct {
	OK          bool   `json:"ok"`
	Result      *T     `json:"result,omitempty"`
	Description string `json:"description,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
	Parameters  struct {
		RetryAfter      int `json:"retry_after,omitempty"`
		MigrateToChatID int `json:"migrate_to_chat_id,omitempty"`
	} `json:"parameters,omitzero"`
}

func request[T any, K any](b *Bot, ctx context.Context, method string, params *T, result *K) error {
	if params == nil {
		return simpleRequest(b, ctx, method, result)
	}

	return multipartRequest(b, ctx, method, params, result)
}

func simpleRequest[K any](b *Bot, ctx context.Context, method string, result *K) error {
	url := fmt.Sprintf("%s/bot%s/%s", b.url, b.token, method)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("create request error for %s: %w", method, err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := b.client.Do(req)
	if err != nil {
		return fmt.Errorf("http request error for %s: %w", method, err)
	}
	defer resp.Body.Close()

	return processResponse(resp, method, result)
}

func multipartRequest[T any, K any](b *Bot, ctx context.Context, method string, params *T, result *K) error {
	pr, pw := io.Pipe()
	form := multipart.NewWriter(pw)

	go func() {
		defer func() {
			if err := form.Close(); err != nil {
				pw.CloseWithError(fmt.Errorf("form close error: %w", err))
				return
			}

			pw.Close()
		}()

		if err := writeParamsToForm(form, params); err != nil {
			pw.CloseWithError(fmt.Errorf("write params error: %w", err))
			return
		}
	}()

	url := fmt.Sprintf("%s/bot%s/%s", b.url, b.token, method)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, pr)
	if err != nil {
		pr.CloseWithError(err)
		return fmt.Errorf("create request error for %s: %w", method, err)
	}

	req.Header.Set("Content-Type", form.FormDataContentType())

	resp, err := b.client.Do(req)
	if err != nil {
		return fmt.Errorf("http request error for %s: %w", method, err)
	}
	defer resp.Body.Close()

	return processResponse(resp, method, result)
}

func processResponse[K any](resp *http.Response, method string, result *K) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body for method %s: %w", method, err)
	}

	if len(body) == 0 {
		return nil
	}

	var r apiResponse[K]
	if err := json.Unmarshal(body, &r); err != nil {
		return fmt.Errorf("error decoding response body for method %s: %w", method, err)
	}

	if !r.OK {
		return fmt.Errorf("telegram API error for method %s: %d %s", method, r.ErrorCode, r.Description)
	}

	if result != nil && r.Result != nil {
		*result = *r.Result
	}

	return nil
}
