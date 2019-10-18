package app

type PeptideRepository interface {
	Create(peptide *Peptide) error
	FindById(id string) (*Peptide, error)
	FindAll() ([]*Peptide, error)
}
