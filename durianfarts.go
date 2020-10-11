package durianfarts

import (
	"encoding/json"

	"github.com/iancoleman/strcase"
)

// A Transform is any function that takes a string and returns a
// string.
type Transform func(s string) string

// ToSnake takes a JSON object as a string, and transforms the keys
// into snake case.
func ToSnake(s string) (string, error) {
	return applyTransform(s, strcase.ToSnake)
}

// ToScreamingSnake takes a JSON object as a string, and transforms
// the keys into screaming snake case.
func ToScreamingSnake(s string) (string, error) {
	return applyTransform(s, strcase.ToScreamingSnake)
}

// ToKebab takes a JSON object as a string, and transforms the keys
// into kebab case.
func ToKebab(s string) (string, error) {
	return applyTransform(s, strcase.ToKebab)
}

// ToScreamingKebab takes a JSON object as a string, and transforms
// the keys into screaming kebab case.
func ToScreamingKebab(s string) (string, error) {
	return applyTransform(s, strcase.ToScreamingKebab)
}

// ToCamel takes a JSON object as a string, and transforms the keys
// into camel case.
func ToCamel(s string) (string, error) {
	return applyTransform(s, strcase.ToCamel)
}

// ToLowerCamel takes a JSON object as a string, and transforms the
// keys into (lower) camel case.
func ToLowerCamel(s string) (string, error) {
	return applyTransform(s, strcase.ToLowerCamel)
}

// ToCustomTransform takes a JSON object as a string, and transforms
// the keys using a custom transformation function.
func ToCustomTransform(s string, t Transform) (string, error) {
	return applyTransform(s, t)
}

func applyTransform(s string, t Transform) (string, error) {
	m := map[string]interface{}{}
	err := json.Unmarshal([]byte(s), &m)
	if nil != err {
		return "", err
	}

	outMap := parseMap(m, t)
	out, err := json.Marshal(outMap)
	if nil != err {
		return "", err
	}
	return string(out), nil
}

func parseMap(m map[string]interface{}, transform Transform) map[string]interface{} {
	for key, val := range m {
		newKey := transform(key)
		delete(m, key)
		switch val.(type) {
		case map[string]interface{}:
			m[newKey] = parseMap(val.(map[string]interface{}), transform)
		case []interface{}:
			m[newKey] = parseArray(val.([]interface{}), transform)
		default:
			m[newKey] = val
		}
	}

	return m
}

func parseArray(a []interface{}, transform Transform) []interface{} {
	for i, val := range a {
		switch val.(type) {
		case map[string]interface{}:
			a[i] = parseMap(val.(map[string]interface{}), transform)
		case []interface{}:
			a[i] = parseArray(val.([]interface{}), transform)
		}
	}

	return a
}
