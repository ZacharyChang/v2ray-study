package errors_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"

	. "v2ray.com/ext/assert"
	. "v2ray.com/v2ray-study/common/errors"
	"v2ray.com/v2ray-study/common/log"
)

func TestError(t *testing.T) {
	assert := With(t)

	err := New("TestError")
	assert(GetSeverity(err), Equals, log.Severity_Info)

	err = New("TestError2").Base(io.EOF)
	assert(GetSeverity(err), Equals, log.Severity_Info)

	err = New("TestError3").Base(io.EOF).AtWarning()
	assert(GetSeverity(err), Equals, log.Severity_Warning)

	err = New("TestError4").Base(io.EOF).AtWarning()
	err = New("TestError5").Base(err)
	assert(GetSeverity(err), Equals, log.Severity_Warning)
	assert(err.Error(), HasSubstring, "EOF")

	t.Log(fmt.Sprintf("TestError5: %s\n", err.Error()))

	err = New("TestError6").Base(New("baseError")).WithPathObj(e{})
	t.Log(fmt.Sprintf("TestError6: %s\n", err.Error()))
}

type e struct{}

func TestErrorMessage(t *testing.T) {
	data := []struct {
		err error
		msg string
	}{
		// Note: compare test
		{
			err: New("a").Base(New("b")).WithPathObj(e{}),
			msg: "v2ray.com/v2ray-study/common/errors_test: a > b",
		},
		{
			err: New("a").Base(New("b").WithPathObj(e{})),
			msg: "a > v2ray.com/v2ray-study/common/errors_test: b",
		},
	}

	for _, d := range data {
		if diff := cmp.Diff(d.msg, d.err.Error()); diff != "" {
			t.Error(diff)
		}
	}
}
