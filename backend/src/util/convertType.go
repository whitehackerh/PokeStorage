package util

func IntToPointer(value int) *int {
	return &value
}

func PointerToInt(value *int) int {
	if value == nil {
		return 0
	}
	return *value
}

func StringToPointer(value string) *string {
	return &value
}

func PointerToString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
