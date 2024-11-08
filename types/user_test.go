package types

import "testing"

func TestUserNameValidation(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
	}{
		{"User", true},
		{"u", false},
		{"thisisaverylongusername", false},
		{"user_name", true},
		{"user name", false},
		{"username_", false},
		{"_username", false},
		{"user$name", false},
		{"user@name", false},
	}

	for _, test := range tests {
		result := userNameValidate(test.name)

		if result != test.expected {
			t.Errorf("usernameValidate(%q) = %v; want %v", test.name, result, test.expected)
		}
	}
}

func TestUserEmailValidation(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"user.name+tag+sorting@example.com", true},
		{"test.email@example.co.uk", true},
		{"invalid-email@", false},
		{"@missingusername.com", false},
		{"missingatsign.com", false},
		{"missingdomain@.com", false},
		{"test@.com", false},
		{"test@com.", false},
		{"test@com", false},
		{"test@-example.com", false},
		{"test@example-.com", false},
	}

	for _, test := range tests {
		result := userEmailValidate(test.email)

		if result != test.expected {
			t.Errorf("emailValidate(%q) = %v; want %v", test.email, result, test.expected)
		}
	}
}
