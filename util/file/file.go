package file

import (
	"fmt"
	"os"
)

func IsExist(path string) bool {
	info, err := os.Stat(path)
	fmt.Println(info)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}
	return false
}
