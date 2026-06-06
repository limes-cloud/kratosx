package config

import (
	"fmt"
	"os"
	"strings"
)

func resolvePlaceholders(input map[string]any) error {
	mapper := placeholderMapper(input)
	resolveMap(input, mapper)
	return nil
}

func resolveMap(input map[string]any, mapper func(string) string) {
	for key, val := range input {
		input[key] = resolveValue(val, mapper)
	}
}

func resolveValue(val any, mapper func(string) string) any {
	switch vt := val.(type) {
	case string:
		return expandPlaceholders(vt, mapper)
	case map[string]any:
		resolveMap(vt, mapper)
		return vt
	case []any:
		for i, item := range vt {
			vt[i] = resolveValue(item, mapper)
		}
		return vt
	default:
		return val
	}
}

func expandPlaceholders(input string, mapper func(string) string) string {
	return os.Expand(input, mapper)
}

func placeholderMapper(input map[string]any) func(string) string {
	return func(name string) string {
		args := strings.SplitN(strings.TrimSpace(name), ":", 2)
		key := strings.TrimSpace(args[0])
		if key == "" {
			return ""
		}

		if val, ok := readPlaceholderValue(input, key); ok {
			return stringifyPlaceholderValue(val)
		}
		if val, ok := os.LookupEnv(key); ok {
			return val
		}
		if len(args) > 1 {
			return args[1]
		}
		return ""
	}
}

func readPlaceholderValue(input map[string]any, path string) (any, bool) {
	keys := strings.Split(path, ".")
	var val any = input
	for _, key := range keys {
		next, ok := val.(map[string]any)
		if !ok {
			return nil, false
		}
		val, ok = next[key]
		if !ok {
			return nil, false
		}
	}
	return val, true
}

func stringifyPlaceholderValue(val any) string {
	switch vt := val.(type) {
	case nil:
		return ""
	case string:
		return vt
	case []byte:
		return string(vt)
	case fmt.Stringer:
		return vt.String()
	case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return fmt.Sprint(vt)
	default:
		return ""
	}
}
