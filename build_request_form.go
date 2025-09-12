package gogram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"reflect"
	"strings"

	"github.com/Alsond5/gogram/models"
)

type inputMediaOrPaidMedia interface {
	GetData() io.Reader
	GetMedia() string
	MarshalInputMedia() ([]byte, error)
}

func writeParamsToForm[T any](form *multipart.Writer, params *T) error {
	v := reflect.ValueOf(params).Elem()

	for i := range v.NumField() {
		jsonTag := v.Type().Field(i).Tag.Get("json")
		if jsonTag == "-" || jsonTag == "" {
			continue
		}

		parts := strings.Split(jsonTag, ",")

		fieldName := parts[0]
		omitempty := strings.Contains(jsonTag, ",omitempty")

		if omitempty && v.Field(i).IsZero() {
			continue
		}

		var err error

		switch vv := v.Field(i).Interface().(type) {
		case models.InputMedia:
			err = addFormFieldInputMedia(form, fieldName, vv)
		case string:
			err = addFormFieldString(form, fieldName, vv)
		case *models.InputFileUpload:
			err = addFormFieldInputFileUpload(form, fieldName, vv)
		case *models.InputFileString:
			err = addFormFieldString(form, fieldName, vv.Data)
		case []models.InputMedia:
			var ss []inputMediaOrPaidMedia
			for _, m := range vv {
				ss = append(ss, m)
			}

			err = addFormFieldInputMediaSlice(form, fieldName, ss)
		case []models.InputPaidMedia:
			var ss []inputMediaOrPaidMedia
			for _, m := range vv {
				ss = append(ss, m)
			}

			err = addFormFieldInputMediaSlice(form, fieldName, ss)
		default:
			err = addFormFieldDefault(form, fieldName, vv)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func addFormFieldInputMediaSlice(form *multipart.Writer, fieldName string, value []inputMediaOrPaidMedia) error {
	var fields []string
	for _, vv := range value {
		data, err := addFormFieldInputMediaItem(form, vv)
		if err != nil {
			return err
		}

		fields = append(fields, string(data))
	}

	w, err := form.CreateFormField(fieldName)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, strings.NewReader("["+strings.Join(fields, ",")+"]"))
	return err
}

func addFormFieldInputMediaItem(form *multipart.Writer, vv inputMediaOrPaidMedia) ([]byte, error) {
	if !strings.HasPrefix(vv.GetMedia(), "attach://") {
		return vv.MarshalInputMedia()
	}

	filename := strings.TrimPrefix(vv.GetMedia(), "attach://")
	field, err := form.CreateFormFile(filename, filename)

	if err != nil {
		return nil, err
	}

	_, err = io.Copy(field, vv.GetData())
	if err != nil {
		return nil, err
	}

	return vv.MarshalInputMedia()
}

func addFormFieldInputMedia(form *multipart.Writer, fieldName string, vv inputMediaOrPaidMedia) error {
	data, err := addFormFieldInputMediaItem(form, vv)
	if err != nil {
		return err
	}

	w, err := form.CreateFormField(fieldName)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bytes.NewReader(data))
	return err
}

func addFormFieldInputFileUpload(form *multipart.Writer, fieldName string, value *models.InputFileUpload) error {
	if value.Data == nil {
		return fmt.Errorf("nil data for field %s", fieldName)
	}

	w, err := form.CreateFormFile(fieldName, value.Filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, value.Data)
	return err
}

func addFormFieldDefault(form *multipart.Writer, fieldName string, value any) error {
	d, err := json.Marshal(value)
	if err != nil {
		return err
	}

	d = bytes.Trim(d, "\"")
	w, err := form.CreateFormField(fieldName)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bytes.NewReader(d))
	return err
}

func addFormFieldString(form *multipart.Writer, fieldname string, value string) error {
	w, err := form.CreateFormField(fieldname)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, strings.NewReader(value))
	return err
}
