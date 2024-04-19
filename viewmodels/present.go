package viewmodels

type PresentListItemViewModel struct {
	Id            string
	Name          string
	Link          string
	UpvoteCount   int32
	DownvoteCount int32
	ShowComments  bool
	CreatedBy     string
	CreatedAt     string
	Comments      CommentListViewModel
}

type PresentListViewModel struct {
	EventId  string
	Presents []PresentListItemViewModel
}
