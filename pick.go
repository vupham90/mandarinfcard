package mandarinfcard

import (
	"time"
)

func Pick(now time.Time, i, size int) int {
	hsh := dateHash(now)
	baseid := (hsh + i) % size

	return slope(baseid, size)
}

func dateHash(now time.Time) int {
	return int(now.Unix()) / 86400
}

func slope(x, size int) int {
	return x % size
}
