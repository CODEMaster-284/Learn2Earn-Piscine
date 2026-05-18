package quadcmd

import (
	"os"
	"strconv"
)

func ParseSize() (int, int, bool) {
	if len(os.Args) != 3 {
		return 0, 0, false
	}

	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return 0, 0, false
	}

	y, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return 0, 0, false
	}

	return x, y, true
}
