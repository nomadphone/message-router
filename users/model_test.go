package users

import "testing"

func TestCreateUser(t *testing.T) {
	user := User{}
	user.TwillioPhone = "1234567890"
	user.TelegramUsername = "test"
	user.TelegramChatID = 1234567890
	user.Name = "test"
	if user.TwillioPhone != "1234567890" {
		t.Errorf("TwillioPhone not set correctly")
	}
	if user.TelegramUsername != "test" {
		t.Errorf("TelegramUsername not set correctly")
	}
	if user.TelegramChatID != 1234567890 {
		t.Errorf("TelegramChatID not set correctly")
	}
	if user.Name != "test" {
		t.Errorf("Name not set correctly")
	}
}
