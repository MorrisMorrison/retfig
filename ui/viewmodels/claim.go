package viewmodels

type ClaimButtonViewModel struct {
	EventId   string
	PresentId string
	IsClaimed bool
	ClaimedBy string
}

func NewClaimButtonViewModel(eventId string, presentId string, isClaimed bool, claimedBy string) *ClaimButtonViewModel {
	return &ClaimButtonViewModel{
		EventId:   eventId,
		PresentId: presentId,
		IsClaimed: isClaimed,
		ClaimedBy: claimedBy,
	}
}
