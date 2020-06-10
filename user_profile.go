package discordgo

type Account struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Verified bool   `json:"verified"`
}

// A UserProfile stores all profile data for an individual Discord user.
type UserProfile struct {
	User              User      `json:"user"`
	ConnectedAccounts []Account `json:"connected_accounts"`
}
