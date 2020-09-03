package utils

import "strconv"

func HexToInt(hex string) (int, error) {
	c, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return 0, err
	}

	return int(c), nil
}
