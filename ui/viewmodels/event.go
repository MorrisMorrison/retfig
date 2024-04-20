package viewmodels

type GetEventViewModel struct {
	Name           string
	CurrentUser    string
	CreatedBy      string
	Recipient      string
	InvitationLink string
	Participants   []string
	Presents       *PresentListViewModel
}
