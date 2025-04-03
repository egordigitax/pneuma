package jsonschema

import (
	"encoding/json"
	"reflect"
)

type JSONSchemaWrapper struct {
	Format JSONSchemaFormat `json:"format"`
}

type JSONSchemaFormat struct {
	Type   string     `json:"type"`
	Name   string     `json:"name"`
	Schema JSONSchema `json:"schema"`
	Strict bool       `json:"strict"`
}

type JSONSchema struct {
	Type                 string                      `json:"type"`
	Properties           map[string]*JSONSchemaField `json:"properties"`
	Required             []string                    `json:"required"`
	AdditionalProperties bool                        `json:"additionalProperties"`
}

type JSONSchemaField struct {
	Type                 string                      `json:"type,omitempty"`
	Properties           map[string]*JSONSchemaField `json:"properties,omitempty"`
	Required             []string                    `json:"required,omitempty"`
	Items                *JSONSchemaField            `json:"items,omitempty"`
	AdditionalProperties bool                        `json:"additionalProperties"`
	Description          string                      `json:"description,omitempty"`
}

func GenerateJSONSchema(v interface{}) (json.RawMessage, error) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	props, required := buildTypedProperties(t)

	wrapper := JSONSchema{
		Type:                 "object",
		Properties:           props,
		Required:             required,
		AdditionalProperties: false,
	}

	return json.MarshalIndent(wrapper, "", "  ")
}

func GetName(v interface{}) string {
	t := reflect.TypeOf(v)
	if t == nil {
		return "<nil>"
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}

func buildTypedProperties(t reflect.Type) (map[string]*JSONSchemaField, []string) {
	props := make(map[string]*JSONSchemaField)
	required := make([]string, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := field.Name
		required = append(required, name)
		props[name] = buildTypedField(field.Type, field.Tag.Get("pneuma"))
	}

	return props, required
}

func buildTypedField(t reflect.Type, tag string) *JSONSchemaField {
	switch t.Kind() {
	case reflect.Struct:
		props, req := buildTypedProperties(t)
		return &JSONSchemaField{
			Type:                 "object",
			Properties:           props,
			Required:             req,
			Description:          tag,
			AdditionalProperties: false,
		}
	case reflect.Slice:
		return &JSONSchemaField{
			Type:                 "array",
			Items:                buildTypedField(t.Elem(), ""),
			Description:          tag,
			AdditionalProperties: false,
		}
	default:
		return &JSONSchemaField{
			Type:                 goTypeToJSONType(t),
			Description:          tag,
			AdditionalProperties: false,
		}
	}
}

func goTypeToJSONType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.String:
		return "string"
	case reflect.Int, reflect.Int8, reflect.Int64, reflect.Int32,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "integer"
	case reflect.Float32, reflect.Float64:
		return "number"
	case reflect.Bool:
		return "boolean"
	default:
		return "object"
	}
}
