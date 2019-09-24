package register

import (
    "testing"
)

func TestNewConsul (t *testing.T) {
	r := NewConsul()
    if r != nil {
        t.Errorf("", r)
    }
}