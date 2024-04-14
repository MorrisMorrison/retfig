package services

import "github.com/MorrisMorrison/retfig/repositories"

type PresentService struct {
	presentRepository repositories.PresentRepository
}

func NewPresentService(presentRepository *repositories.PresentRepository) *PresentService {
	return &PresentService{presentRepository: *presentRepository}
}
