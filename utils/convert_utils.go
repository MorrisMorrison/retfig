package utils

import (
	"math/rand"
	"time"
)

func ToDuration(value int) time.Duration {
	return time.Duration(rand.Int31n(int32(value)))
}
