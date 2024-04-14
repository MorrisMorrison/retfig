package viewmodels

type GetEventViewModel struct {
	Name           string
	CreatorEmail   string
	Recipient      string
	InvitationLink string
	Participants   []string
}
