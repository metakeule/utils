package utf8

import (
	"testing"
)

func TestRPad(t *testing.T) {
	corpus := map[string]string{
		"a":     "a   ",
		"abcd":  "abcd",
		"äbcde": "äbcde",
		"äbcd":  "äbcd",
		"äbc":   "äbc ",
	}

	for k, v := range corpus {
		r := RPad(k, 4)
		if r != v {
			t.Errorf("wrong RPad 4 for %#v, expected: %#v, got: %#v", k, v, r)
		}
	}
}

func TestRPadWith(t *testing.T) {
	corpus := map[string]string{
		"a":     "a---",
		"abcd":  "abcd",
		"äbcde": "äbcde",
		"äbcd":  "äbcd",
		"äbc":   "äbc-",
	}

	for k, v := range corpus {
		r := RPadWith(k, 4, "-")
		if r != v {
			t.Errorf("wrong RPadWith 4, '-' for %#v, expected: %#v, got: %#v", k, v, r)
		}
	}
}

func TestLPad(t *testing.T) {
	corpus := map[string]string{
		"a":     "   a",
		"abcd":  "abcd",
		"äbcde": "äbcde",
		"äbcd":  "äbcd",
		"äbc":   " äbc",
	}

	for k, v := range corpus {
		r := LPad(k, 4)
		if r != v {
			t.Errorf("wrong LPad 4 for %#v, expected: %#v, got: %#v", k, v, r)
		}
	}
}

func TestLPadWith(t *testing.T) {
	corpus := map[string]string{
		"a":     "---a",
		"abcd":  "abcd",
		"äbcde": "äbcde",
		"äbcd":  "äbcd",
		"äbc":   "-äbc",
	}

	for k, v := range corpus {
		r := LPadWith(k, 4, "-")
		if r != v {
			t.Errorf("wrong LPadWith 4, '-' for %#v, expected: %#v, got: %#v", k, v, r)
		}
	}
}
