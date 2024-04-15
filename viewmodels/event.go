package viewmodels

type GetEventViewModel struct {
	Name           string
	CurrentUser    string
	Creator        string
	Recipient      string
	InvitationLink string
	Participants   []string
	Presents       PresentListViewModel
}
