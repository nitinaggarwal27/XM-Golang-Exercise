package model

import "testing"

func TestInitAdminAccount(t *testing.T) {
	acc := InitAdminAccount()

	if acc.Email != "admin@xm.com" {
		t.Error("test case fail")
	}
}
