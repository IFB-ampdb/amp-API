package app

import (
	"time"
)

type PeptideService interface {
	CreatePeptide(peptide *Peptide) error
	FindPeptideByID(id string) (*Peptide, error)
	FindAllPeptides() ([]*Peptide, error)
}

type peptideService struct {
	repo PeptideRepository
}

func NewPeptideService(repo PeptideRepository) PeptideService {
	return &peptideService{
		repo,
	}
}

func (s *peptideService) CreatePeptide(peptide *Peptide) error {
	peptide.Created = time.Now()
	peptide.Updated = time.Now()
	return s.repo.Create(peptide)
}

func (s *peptideService) FindPeptideByID(id string) (*Peptide, error) {
	return s.repo.FindById(id)
}

func (s *peptideService) FindAllPeptides() ([]*Peptide, error) {
	return s.repo.FindAll()
}
