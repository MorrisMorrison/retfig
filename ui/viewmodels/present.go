package viewmodels

type PresentListItemViewModel struct {
	PresentId     string
	EventId       string
	Name          string
	Link          string
	UpvoteCount   int32
	DownvoteCount int32
	CommentCount  int32
	CreatedBy     string
	CreatedAt     string
	Comments      *CommentListViewModel
}

type PresentListViewModel struct {
	EventId  string
	Presents []*PresentListItemViewModel
}
