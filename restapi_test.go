package discordgo

import (
	"testing"
)

//////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////// START OF TESTS

// TestLogout tests the Logout() function. This should not return an error.
func TestLogout(t *testing.T) {

	if dg == nil || dg.Token == "" {
		t.Skip("Cannot test logout, dg.Token not set.")
	}

	err := dg.Logout()
	if err != nil {
		t.Errorf("Logout() returned error: %+v", err)
	}
}

func TestUserAvatar(t *testing.T) {
	if !isConnected() {
		t.Skip("Skipped, Not connected to Discord.")
	}

	a, err := dg.UserAvatar("@me")
	if err != nil {
		if err.Error() == `HTTP 404 NOT FOUND, {"message": ""}` {
			t.Skip("Skipped, @me doesn't have an Avatar")
		}
		t.Errorf(err.Error())
	}

	if a == nil {
		t.Errorf("a == nil, should be image.Image")
	}
}

func TestUserUpdate(t *testing.T) {

	if envEmail == "" || envPassword == "" {
		t.Skip("Skipping, DG_USERNAME or DG_PASSWORD not set")
		return
	}

	if !isConnected() {
		t.Skip("Skipped, Not connected to Discord.")
	}

	u, err := dg.User("@me")
	if err != nil {
		t.Errorf(err.Error())
	}

	s, err := dg.UserUpdate(envEmail, envPassword, "testname", u.Avatar, "")
	if err != nil {
		t.Error(err.Error())
	}
	if s.Username != "testname" {
		t.Error("Username != testname")
	}
	s, err = dg.UserUpdate(envEmail, envPassword, u.Username, u.Avatar, "")
	if err != nil {
		t.Error(err.Error())
	}
	if s.Username != u.Username {
		t.Error("Username != " + u.Username)
	}
}

//func (s *Session) UserChannelCreate(recipientID string) (st *Channel, err error) {

func TestUserChannelCreate(t *testing.T) {

	if !isConnected() {
		t.Skip("Skipped, Not connected to Discord.")
	}

	if envAdmin == "" {
		t.Skip("Skipped, DG_ADMIN not set.")
	}

	_, err := dg.UserChannelCreate(envAdmin)
	if err != nil {
		t.Errorf(err.Error())
	}

	// TODO make sure the channel was added
}

func TestUserChannels(t *testing.T) {

	if !isConnected() {
		t.Skip("Skipped, Not connected to Discord.")
	}

	_, err := dg.UserChannels()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestUserGuilds(t *testing.T) {

	if !isConnected() {
		t.Skip("Skipped, Not connected to Discord.")
	}

	_, err := dg.UserGuilds()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestUserSettings(t *testing.T) {

	if !isConnected() {
		t.Skip("Skipped, Not connected to Discord.")
	}

	_, err := dg.UserSettings()
	if err != nil {
		t.Errorf(err.Error())
	}
}