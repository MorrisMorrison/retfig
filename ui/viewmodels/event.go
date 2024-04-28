package viewmodels

type GetEventViewModel struct {
	Name           string
	CreatedBy      string
	Recipient      string
	InvitationLink string
	Participants   []string
	Presents       *PresentListViewModel
}

func NewGetEventViewModel(name string, createdBy string, recipient string, invitationLink string, participants []string, presents *PresentListViewModel) *GetEventViewModel {
	return &GetEventViewModel{
		Name:           name,
		CreatedBy:      createdBy,
		Recipient:      recipient,
		InvitationLink: invitationLink,
		Participants:   participants,
		Presents:       presents,
	}
}
