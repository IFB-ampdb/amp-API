package psql

import (
	"database/sql"
	"fmt"
	"github.com/ifbampdb/amp-core/app"
	"log"

	_ "github.com/lib/pq"
)

type peptideRepository struct {
	db *sql.DB
}

func PostgresConnection(database string) *sql.DB {
	fmt.Println("Connecting to PostgreSQL DB")
	db, err := sql.Open("postgres", database)
	if err != nil {
		log.Fatalf("%s", err)
		panic(err)
	}
	_, err = db.Query(`CREATE TABLE IF NOT EXISTS peptides (
	  pdb_id VARCHAR(10) NOT NULL,
	  is_amp BOOLEAN,
	  hfobic_area  VARCHAR(96),
	  hfobic_avg FLOAT,
	  hairpin BOOLEAN, 
	  beta_sheet BOOLEAN,
	  alpha_helix BOOLEAN, 
	  alpha_helix_beta_sheet BOOLEAN,
	  alpha_helix_beta_sheet_hairpin BOOLEAN,
	  charge FLOAT,
	  m_dipol FLOAT,
	  charge_amt_atm FLOAT,
	  m_dipol_amt_atm FLOAT,
	  name varchar(256),
	  sequence varchar(256),
	  organism varchar(256),
	  created timestamp NOT NULL DEFAULT current_timestamp,
	  updated timestamp NULL DEFAULT NULL,
	  PRIMARY KEY(pdb_id)
	);`)
	if err != nil {
		log.Fatalf("%s", err)
		panic(err)
	}
	return db
}

func NewPostgresPeptideRepository(db *sql.DB) app.PeptideRepository {
	return &peptideRepository{
		db,
	}
}

func (r *peptideRepository) Create(peptide *app.Peptide) error {
	err := r.db.QueryRow("INSERT INTO peptides(pdb_id, is_amp, hfobic_area, hfobic_avg, hairpin, beta_sheet, alpha_helix, alpha_helix_beta_sheet, alpha_helix_beta_sheet_hairpin, charge, m_dipol, charge_amt_atm, m_dipol_amt_atm, name, sequence, organism, created, updated) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18) RETURNING pdb_id",
		peptide.PdbID, peptide.IsAMP, peptide.HforbicArea, peptide.HfobicAvg, peptide.Hairpin, peptide.BetaSheet, peptide.AlphaHelix, peptide.AlphaHelixBetaSheet, peptide.AlphaHelixBetaSheetHairpin, peptide.Charge, peptide.DipolMomentum, peptide.ChargeAmtAtm, peptide.DipolMomentumAmtAtm, peptide.Name, peptide.Sequence, peptide.Organism, peptide.Created, peptide.Updated).Scan(&peptide.PdbID)
	return err
}

func (r *peptideRepository) FindById(id string) (*app.Peptide, error) {
	peptide := new(app.Peptide)
	err := r.db.QueryRow("SELECT pdb_id, is_amp, hfobic_area, hfobic_avg, hairpin, beta_sheet, alpha_helix, alpha_helix_beta_sheet, alpha_helix_beta_sheet_hairpin, charge, m_dipol, charge_amt_atm, m_dipol_amt_atm, name, sequence, organism, created, updated FROM peptides where pdb_id=$1", id).Scan(&peptide.PdbID, &peptide.IsAMP, &peptide.HforbicArea, &peptide.HfobicAvg, &peptide.Hairpin, &peptide.BetaSheet, &peptide.AlphaHelix, &peptide.AlphaHelixBetaSheet, &peptide.AlphaHelixBetaSheetHairpin, &peptide.Charge, &peptide.DipolMomentum, &peptide.ChargeAmtAtm, &peptide.DipolMomentumAmtAtm, &peptide.Name, &peptide.Sequence, &peptide.Organism, &peptide.Created, &peptide.Updated)
	return peptide, err
}

func (r *peptideRepository) FindAll() (peptides []*app.Peptide, err error) {
	rows, err := r.db.Query("SELECT pdb_id,is_amp, hfobic_area, hfobic_avg, hairpin, beta_sheet, alpha_helix, alpha_helix_beta_sheet, alpha_helix_beta_sheet_hairpin, charge, m_dipol, charge_amt_atm, m_dipol_amt_atm, name, sequence, organism, created, updated FROM peptides")
	defer rows.Close()

	for rows.Next() {
		peptide := new(app.Peptide)
		if err = rows.Scan(&peptide.PdbID, &peptide.IsAMP, &peptide.HforbicArea, &peptide.HfobicAvg, &peptide.Hairpin, &peptide.BetaSheet, &peptide.AlphaHelix, &peptide.AlphaHelixBetaSheet, &peptide.AlphaHelixBetaSheetHairpin, &peptide.Charge, &peptide.DipolMomentum, &peptide.ChargeAmtAtm, &peptide.DipolMomentumAmtAtm, &peptide.Name, &peptide.Sequence, &peptide.Organism, &peptide.Created, &peptide.Updated); err != nil {
			log.Print(err)
			return nil, err
		}

		peptides = append(peptides, peptide)

	}
	return peptides, nil
}
