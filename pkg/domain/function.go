package domain

import "strings"

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

func JoinBody(arr ...string) string {
	list := []string{}
	for _, s := range arr {
		if s == "" {
			continue
		}
		list = append(list, s)
	}
	return strings.Join(list, "\n")
}
