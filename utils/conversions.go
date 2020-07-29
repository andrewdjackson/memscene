package utils

// ConvertBooltoInt converts boolean to integer
func ConvertBooltoInt(b bool) uint8 {
	if b {
		return 1
	}

	return 0
}
