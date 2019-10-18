package app

import "time"

// Peptide struct declares the information in a peptide
type Peptide struct {
	PdbID                      string    `json:"PdbID" db:"pdb_id"`
	Name                       string    `json:"name" db:"name"`
	Sequence                   string    `json:"sequence" db:"sequence"`
	Organism                   string    `json:"organism" db:"organism"`
	IsAMP                      bool      `json:"is_amp" db:"is_amp"`
	HforbicArea                string    `json:"hfobic_area" db:"hfobic_area"`
	HfobicAvg                  float32   `json:"hfobic_avg" db:"hfobic_avg"`
	Hairpin                    bool      `json:"hairpin" db:"hairpin"`
	BetaSheet                  bool      `json:"beta_sheet" db:"beta_sheet"`
	AlphaHelix                 bool      `json:"alpha_helix" db:"alpha_helix"`
	AlphaHelixBetaSheet        bool      `json:"alpha_helix_beta_sheet" db:"alpha_helix_beta_sheet"`
	AlphaHelixBetaSheetHairpin bool      `json:"alpha_helix_beta_sheet_hairpin" db:"alpha_helix_beta_sheet_hairpin"`
	Charge                     float32   `json:"charge" db:"charge"`
	DipolMomentum              float32   `json:"DipolMomentum" db:"DipolMomentum"`
	ChargeAmtAtm               float32   `json:"charge_amt_atm" db:"charge_amt_atm"`
	DipolMomentumAmtAtm        float32   `json:"m_dipol_amt_atm" db:"m_dipol_amt_atm"`
	Created                    time.Time `json:"created" db:"created"`
	Updated                    time.Time `json:"updated" db:"updated"`
}
