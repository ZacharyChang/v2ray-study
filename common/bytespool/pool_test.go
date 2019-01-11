package bytespool_test

import (
	"math"
	"testing"
	. "v2ray.com/ext/assert"
	. "v2ray.com/v2ray-study/common/bytespool"
)

// Note: test to get pool
func TestGetPool(t *testing.T) {
	assert := With(t)

	pool := GetPool(2049)
	t.Logf("get pool: %v", pool)
	assert(pool, NotEquals, nil)
	pool = GetPool(math.MaxInt32)
	assert(pool, IsNil)
	t.Logf("can not get pool for size: %d", math.MaxInt32)
}
