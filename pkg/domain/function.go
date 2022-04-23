package domain

func GetStringSlicePointer(arr []string) *[]string {
	if arr == nil {
		return nil
	}
	return &arr
}
