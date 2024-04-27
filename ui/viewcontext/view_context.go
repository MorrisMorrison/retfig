package viewcontext

type ViewContext struct {
	CurrentUser  string
	IsAuthorized bool
}

func NewViewContext(username string, isAuthorized bool) *ViewContext {
	return &ViewContext{
		CurrentUser:  username,
		IsAuthorized: isAuthorized,
	}
}
