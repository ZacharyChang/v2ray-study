package common_test

import (
	"errors"
	"testing"

	. "v2ray.com/ext/assert"
	. "v2ray.com/v2ray-study/common"
)

func TestMust(t *testing.T) {
	assert := With(t)

	f := func() error {
		return errors.New("test error")
	}

	assert(func() { Must(f()) }, Panics)
}

func TestMust2(t *testing.T) {
	assert := With(t)

	f := func() (interface{}, error) {
		return nil, errors.New("test error")
	}

	assert(func() { Must2(f()) }, Panics)
}
