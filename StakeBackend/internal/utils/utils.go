package utils


func IsValidAddress(address string) bool {
	return len(address) == 42 && address[:2] == "0x"
}