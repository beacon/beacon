package dao

import (
	"math"
	"testing"
	"time"
)

func TestMax(t *testing.T) {
	v := math.MaxUint32
	tm := time.Unix(int64(v), 0)
	t.Log(tm.String())
}
