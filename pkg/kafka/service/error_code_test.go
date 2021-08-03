package service

import (
	"fmt"
	"testing"
)

func TestErrorCode(t *testing.T) {
	fmt.Println(NONE)
	if NONE != 0 {
		t.Errorf("error code malformed")
	}
	fmt.Println(OFFSET_OUT_OF_RANGE)
	if OFFSET_OUT_OF_RANGE != 1 {
		t.Errorf("error code malformed")
	}
	fmt.Println(MEMBER_ID_REQUIRED)
	if MEMBER_ID_REQUIRED != 79 {
		t.Errorf("error code malformed")
	}
}
