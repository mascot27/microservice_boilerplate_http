package cmd

import "testing"

func TestVerify_credentials_ok(t *testing.T) {
	result := Verify_credentials("admin", "password")
	want := true
	if result != want {
		t.Errorf("creds verification was incorrect, got: %t, want: %t.", result, want)
	}
}

func TestVerify_credentials_bad(t *testing.T) {
	result := Verify_credentials("badUser", "badPassword")
	want := false
	if result != want {
		t.Errorf("creds verification was incorrect, got: %t, want: %t.", result, want)
	}
}

func TestVerify_credentials_badPass(t *testing.T) {
	result := Verify_credentials("admin", "badPassword")
	want := false
	if result != want {
		t.Errorf("creds verification was incorrect, got: %t, want: %t.", result, want)
	}
}

func TestVerify_credentials_badUser(t *testing.T) {
	result := Verify_credentials("badUser", "password")
	want := false
	if result != want {
		t.Errorf("creds verification was incorrect, got: %t, want: %t.", result, want)
	}
}
