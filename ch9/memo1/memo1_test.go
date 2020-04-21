package memo1

import (
	"testing"

	"gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func TestMemo1(t *testing.T) {
	m := New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	memotest.Concurrent(t, m)
}
