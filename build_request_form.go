package gogram

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"reflect"
	"strings"
)

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
		case string:
			err = addFormFieldString(form, fieldName, vv)
		default:
			err = addFormFieldDefault(form, fieldName, vv)
		}

		if err != nil {
			return err
		}
	}

	return nil
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
