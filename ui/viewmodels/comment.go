package viewmodels

type CommentListViewModel struct {
	EventId   string
	PresentId string
	Comments  []*CommentListItemViewModel
}

func NewCommentListViewModel(eventId string, presentId string, comments []*CommentListItemViewModel) *CommentListViewModel {
	return &CommentListViewModel{
		EventId:   eventId,
		PresentId: presentId,
		Comments:  comments,
	}
}

type CommentListItemViewModel struct {
	PresentId string
	Username  string
	Content   string
	CreatedAt string
}

func NewCommentListItemViewModel(presentId string, username string, content string, createdAt string) *CommentListItemViewModel {
	return &CommentListItemViewModel{
		PresentId: presentId,
		Username:  username,
		Content:   content,
		CreatedAt: createdAt,
	}
}
