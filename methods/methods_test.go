package methods

import (
	"testing"
)

func TestValidateEmail(t *testing.T) {

	//test case 1 -> pass right email format
	email := "testing@test.com"
	err := ValidateEmail(email)
	if err == false {
		t.Error("test case fail", err)
	}

	//test case 2 -> pass wrong email format
	email = "testing-est.com"
	err = ValidateEmail(email)
	if err == true {
		t.Error("test case fail", err)
	}
}
func TestCheckPass(t *testing.T) {

	//test case 1 -> pass blank
	err := CheckPassword("")
	if err == true {
		t.Error("test case fail")
	}

	//test case 2 -> pass only string characters
	err = CheckPassword("testtest")
	if err == true {
		t.Error("test case fail")
	}

	//test case 3 -> pass only special characters
	err = CheckPassword("#$%@*&@^")
	if err == true {
		t.Error("test case fail")
	}

	//test case 4 -> pass only special characters and Upper aphabets
	err = CheckPassword("#$%@*&@^TEST")
	if err == true {
		t.Error("test case fail")
	}

	//test case 5 -> pass only special characters ,lowercase aphabets and Upper aphabets
	err = CheckPassword("#$%@*&@^TESTtest")
	if err == true {
		t.Error("test case fail")
	}

	//test case 6 -> pass right
	err = CheckPassword("Test@123")
	if err == false {
		t.Error("test case fail")
	}

	hashpass := HashForNewPassword("Test@123")

	err = CheckHashForPassword("hashpass", "Test#123")
	if err == true {
		t.Error("test case fail")
	}
	err = CheckHashForPassword(hashpass, "Test@123")
	if err == false {
		t.Error("test case fail")
	}
}

func TestRandomMethod(t *testing.T) {

	str := RandomString(6)
	if len([]byte(str)) != 6 {
		t.Error("test case fail", len([]byte(str)))
	}

	id := randInt()
	if id == 0 {
		t.Error("test case fail")
	}
	str = RandomStringIntegerOnly(6)
	if len([]byte(str)) != 6 {
		t.Error("test case fail", len([]byte(str)))
	}
	str = SlugifyEmail("testing@test.com")
	if str != "testi" {
		t.Error("test case fail")
	}
}
