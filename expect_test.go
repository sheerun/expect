package expect_test

import (
	"github.com/sheerun/expect"
	"testing"
)

func TestFail(t *testing.T) {
	expect.Fail(t, "foobar")
}
