CREATE TABLE IF NOT EXISTS peptides
(
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
);

-- pdb_id,is_amp, hfobic_area, hfobic_avg, hairpin, beta_sheet, alpha_helix, alpha_helix_beta_sheet, alpha_helix_beta_sheet_hairpin, charge, m_dipol, charge_amt_atm, m_dipol_amt_atm, name, sequence, organism, created, updated