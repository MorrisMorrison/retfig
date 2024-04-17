package viewmodels

type PresentListItemViewModel struct {
	Id        string
	Name      string
	Link      string
	CreatedBy string
	CreatedAt string
}

type PresentListViewModel struct {
	EventId  string
	Presents []PresentListItemViewModel
}
