package api

import (
	"fmt"
	"testing"
)

func Test_ApiCode(t *testing.T) {
	fmt.Println(JoinGroup)
	if JoinGroup != 11 {
		t.Errorf("api code malformed")
	}
}
