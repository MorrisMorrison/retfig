package viewmodels

type CommentListViewModel struct {
	EventId   string
	PresentId string
	Comments  []*CommentListItemViewModel
}

type CommentListItemViewModel struct {
	PresentId string
	Username  string
	Content   string
	CreatedAt string
}
