package viewmodels

type PresentListItemViewModel struct {
	PresentId         string
	EventId           string
	Name              string
	Link              string
	UpvoteCount       int32
	DownvoteCount     int32
	CommentCount      int32
	CreatedBy         string
	CreatedAt         string
	Comments          *CommentListViewModel
	IsClaimed         bool
	ClaimedBy         string
	IsUpvotedByUser   bool
	IsDownvotedByUser bool
}

func NewPresentListItemViewModel(
	presentId string,
	eventId string,
	name string,
	link string,
	upvoteCount int32,
	downvoteCount int32,
	commentCount int32,
	createdBy string,
	createdAt string,
	comments *CommentListViewModel,
	isClaimed bool,
	claimedBy string,
	isUpvotedByUser bool,
	isDownvotedByUser bool,
) *PresentListItemViewModel {
	return &PresentListItemViewModel{
		PresentId:         presentId,
		EventId:           eventId,
		Name:              name,
		Link:              link,
		UpvoteCount:       upvoteCount,
		DownvoteCount:     downvoteCount,
		CommentCount:      commentCount,
		CreatedBy:         createdBy,
		CreatedAt:         createdAt,
		Comments:          comments,
		IsClaimed:         isClaimed,
		ClaimedBy:         claimedBy,
		IsUpvotedByUser:   isUpvotedByUser,
		IsDownvotedByUser: isDownvotedByUser,
	}
}

type PresentListViewModel struct {
	EventId  string
	Presents []*PresentListItemViewModel
}

func NewPresentListViewModel(eventId string, presents []*PresentListItemViewModel) *PresentListViewModel {
	return &PresentListViewModel{
		EventId:  eventId,
		Presents: presents,
	}
}
