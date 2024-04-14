package viewmodels

type PresentListItemViewModel struct {
	Name string
	Link string
}

type PresentListViewModel struct {
	Presents []PresentListItemViewModel
}
