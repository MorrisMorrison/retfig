package viewmodels

type PresentListItemViewModel struct {
	Name      string
	Link      string
	Creator   string
	CreatedAt string
}

type PresentListViewModel struct {
	EventId  string
	Presents []PresentListItemViewModel
}
