package value

// Bool stores v in a new bool value and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int8 stores v in a new int value and returns a pointer to it.
func Int8(v int8) *int8 { return &v }

// Int stores v in a new int value and returns a pointer to it.
func Int(v int) *int { return &v }

// Int32 stores v in a new int32 value and returns a pointer to it.
func Int32(v int32) *int32 { return &v }

// Int64 stores v in a new int64 value and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// Float32 stores v in a new float32 value and returns a pointer to it.
func Float32(v float32) *float32 { return &v }

// Float64 stores v in a new float64 value and returns a pointer to it.
func Float64(v float64) *float64 { return &v }

// Uint32 stores v in a new uint32 value and returns a pointer to it.
func Uint32(v uint32) *uint32 { return &v }

// Uint64 stores v in a new uint64 value and returns a pointer to it.
func Uint64(v uint64) *uint64 { return &v }

// String stores v in a new string value and returns a pointer to it.
func String(v string) *string { return &v }

// Pointer 通用获取指定数的指针
func Pointer[T any](v T) *T {
	return &v
}

// Value 通用获取指定指针的数
func Value[T any](v *T) T {
	var t = new(T)
	if v != nil {
		*t = *v
	}
	return *t
}
