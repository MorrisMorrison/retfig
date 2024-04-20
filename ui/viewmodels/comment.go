package viewmodels

type CommentListViewModel struct {
	PresentId string
	Comments  []CommentListItemViewModel
}

type CommentListItemViewModel struct {
	Username  string
	Content   string
	CreatedAt string
}
