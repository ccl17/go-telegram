package telegram

type Chat struct {
	ID                  int      `json:"id"`
	Type                string   `json:"type"`
	Title               string   `json:"title"`
	UserName            string   `json:"username"`
	FirstName           string   `json:"first_name"`
	LastName            string   `json:"last_name"`
	AllMembersAreAdmins bool     `json:"all_members_are_administrators"`
	Description         string   `json:"description,omitempty"`
	InviteLink          string   `json:"invite_link,omitempty"`
	PinnedMessage       *Message `json:"pinned_message"`
	// Photo               *ChatPhoto `json:"photo"`
}
