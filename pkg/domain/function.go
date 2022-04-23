package domain

func GetString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func GetStringSlicePointer(arr []string) *[]string {
	if arr == nil {
		return nil
	}
	return &arr
}
