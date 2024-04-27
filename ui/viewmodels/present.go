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

type PresentListViewModel struct {
	EventId  string
	Presents []*PresentListItemViewModel
}
