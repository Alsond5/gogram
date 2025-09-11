package models

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type FieldRule struct {
	FieldName string
	Required  bool
	Forbidden bool
}

type UnionConfig[T any] struct {
	TypeField   *T
	TypeValue   T
	TargetField any
	Rules       []FieldRule
}

func UnmarshalUnion[T any](data []byte, configs []UnionConfig[T]) error {
	var rawMap map[string]any
	if err := json.Unmarshal(data, &rawMap); err != nil {
		return err
	}

	for _, config := range configs {
		if !matchesRules(rawMap, config.Rules) {
			continue
		}

		*config.TypeField = config.TypeValue

		targetValue := reflect.ValueOf(config.TargetField)
		if targetValue.Kind() != reflect.Ptr || targetValue.Elem().Kind() != reflect.Ptr {
			return fmt.Errorf("TargetField must be a pointer to pointer")
		}

		targetType := targetValue.Elem().Type().Elem()
		newInstance := reflect.New(targetType)

		if err := json.Unmarshal(data, newInstance.Interface()); err != nil {
			return err
		}

		targetValue.Elem().Set(newInstance)
		return nil
	}

	return fmt.Errorf("could not determine union type from JSON fields")
}

func MarshalUnion[T comparable](unionType T, cases map[T]any) ([]byte, error) {
	targetStruct, exists := cases[unionType]
	if !exists {
		return nil, fmt.Errorf("unkown union type: %v", unionType)
	}

	if targetStruct == nil {
		return []byte("null"), nil
	}

	return json.Marshal(targetStruct)
}

func RequiredField(fieldName string) FieldRule {
	return FieldRule{FieldName: fieldName, Required: true}
}

func ForbiddenField(fieldName string) FieldRule {
	return FieldRule{FieldName: fieldName, Forbidden: true}
}

func RequiredFields(fieldNames ...string) []FieldRule {
	rules := make([]FieldRule, len(fieldNames))
	for i, name := range fieldNames {
		rules[i] = RequiredField(name)
	}

	return rules
}

func ForbiddenFields(fieldNames ...string) []FieldRule {
	rules := make([]FieldRule, len(fieldNames))
	for i, name := range fieldNames {
		rules[i] = ForbiddenField(name)
	}

	return rules
}

func CombineRules(ruleGroups ...[]FieldRule) []FieldRule {
	var combined []FieldRule
	for _, group := range ruleGroups {
		combined = append(combined, group...)
	}

	return combined
}

func matchesRules(data map[string]any, rules []FieldRule) bool {
	for _, rule := range rules {
		_, fieldExists := data[rule.FieldName]

		if rule.Required && !fieldExists {
			return false
		}

		if rule.Forbidden && fieldExists {
			return false
		}
	}

	return true
}
