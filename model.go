package main

import (
	"database/sql"
)

type doses struct {
	ID              	int     `json:"id"`
	Blockis			int     `json:"Blockis"`
	DrugA           	string  `json:"DrugA"`
	DrugB           	string  `json:"DrugB"`
	DoseA           	float64 `json:"DoseA"`
	DoseB           	float64 `json:"DoseB"`
	Response        	float64 `json:"Response"`
	DSS             	float64 `json:"DSS"`
	Synergy_Bliss     	float64 `json:"Synergy_Bliss"`
	CellLine        	string  `json:"CellLine"`

}

type combination 	struct {
	ID          	int     `json:"id"`

	DrugA       	string  `json:"DrugA"`
	DrugB       	string  `json:"DrugB"`
	CellLine    	string  `json:"CellLine"`
	Source      	string  `json:"Source"`
	DSS         	float64 `json:"DSS"`
	Synergy_Bliss 	float64 `json:"Synergy_Bliss"`
	Blockis		int 	`json:"Blockis"`
}

//add number of cell lines tested
// change add info about matrix comb
type conditions struct {
	ID		int		`json:"id"`
	Source	string	`json:"Source"`
		AssayDetails	struct {
			Readout				string	`json:"Readout"`
			Compound			string	`json:"Compound"`
			ControlNegative		bool	`json:"ControlNegative"`
			ControlPositive		bool	`json:"ControlPositive"`
		}
		NumberOfDrugs	struct {
			Single				int		`json:"Single"`
			DosesSingle			int		`json:"DosesSingle"`
			Combination			int		`json:"Combination"`
			DosesCombination	int		`json:"DosesCombination"`
		}
		Plate			struct {
			Format			int		`json:"Format"`
			CellsPerWell	int		`json:"CellsPerWell"`
			VolumePerWell	int		`json:"VolumePerWell"`
			Unit 			string	`json:"UnitVolume"`
		}

}

type cells struct {
	cell_id          	int     `json:"Id"`
	accession_id		sql.NullString	`json:"AccessionId"`
	cell_name			string	`json:"CellName"`
	tissue_id			int 	`json:"TissueId"`

}

type drugs struct {
	drug_id				int 	`json:"Id"`
	drug_name			string	`json:"DrugName"`
}

func (p *doses) getDose(db *sql.DB) error {
	return db.QueryRow("SELECT ID, DrugA, DrugB, DoseA, DoseB, Response, DSS, Synergy_Bliss, CellLine FROM doses WHERE Blockis=$1",
		p.Blockis).Scan(&p.ID,&p.DrugA, &p.DrugB, &p.DoseA, &p.DoseB, &p.Response, &p.DSS, &p.Synergy_Bliss, &p.CellLine)
}

func (p *doses) updateDose(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE doses SET DrugA=$1, DrugB=$2, DoseA=$3, DoseB=$4, Response=$5, DSS=$6, Synergy_Bliss=$7, CellLine=$8, Blockis=$9 WHERE id=$10",
			p.DrugA, p.DrugB, p.DoseA, p.DoseB, p.Response, p.DSS, p.Synergy_Bliss, p.CellLine, p.Blockis, p.ID)
	return err
}

func (p *doses) deleteDose(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM doses WHERE id=$1", p.ID)
	return err
}

func (p *doses) createDose(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO doses(DrugA, DrugB, DoseA, DoseB, Response, DSS, Synergy_Bliss, CellLine, Blockis) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		p.DrugA, p.DrugB, p.DoseA, p.DoseB, p.Response, p.DSS, p.Synergy_Bliss, p.CellLine, p.Blockis).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func getDoses(db *sql.DB, start, count int) ([]doses, error) {
	rows, err := db.Query(
		"SELECT ID, DrugA, DrugB, DoseA, DoseB, Response, DSS, Synergy_Bliss, CellLine, Blockis FROM doses LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	allDoses := []doses{}

	for rows.Next() {
		var p doses
		if err := rows.Scan(&p.ID, &p.DrugA, &p.DrugB, &p.DoseA, &p.DoseB, &p.Response, &p.DSS, &p.Synergy_Bliss, &p.CellLine, &p.Blockis); err != nil {
			return nil, err
		}
		allDoses = append(allDoses, p)
	}
	return allDoses, nil
}

func getCells(db *sql.DB, start, count int) ([]cells, error) {
	rows, err := db.Query(
		"SELECT cell_id, accession_id, cell_name, tissue_id FROM cells LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	allCells := []cells{}

	for rows.Next() {
		var p cells
		if err := rows.Scan(&p.cell_id, &p.accession_id, &p.cell_name, &p.tissue_id); err != nil {
			return nil, err
		}
		allCells = append(allCells, p)
	}
	return allCells, nil
}

func getDrugs(db *sql.DB, start, count int) ([]drugs, error) {
	rows, err := db.Query(
		"SELECT drug_id, drug_name FROM drugs LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	allDrugs := []drugs{}

	for rows.Next() {
		var p drugs
		if err := rows.Scan(&p.drug_id, &p.drug_name); err != nil {
			return nil, err
		}
		allDrugs = append(allDrugs, p)
	}
	return allDrugs, nil
}

func getDosesByID(db *sql.DB, Blockis int) ([]doses, error) {
	rows, err := db.Query(
		"SELECT ID, DrugA, DrugB, DoseA, DoseB, Response, DSS, Synergy_Bliss, CellLine, Blockis FROM doses WHERE Blockis=$1",
		Blockis)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allDosesByID := []doses{}

	for rows.Next() {
		var p doses
		if err := rows.Scan(&p.ID, &p.DrugA, &p.DrugB, &p.DoseA, &p.DoseB, &p.Response, &p.DSS, &p.Synergy_Bliss, &p.CellLine, &p.Blockis); err != nil {
			return nil, err
		}
		allDosesByID = append(allDosesByID, p)
	}
	return allDosesByID, nil
}

func (p *combination) getCombination(db *sql.DB) error {
	return db.QueryRow("SELECT DrugA, DrugB, CellLine, Source, DSS, Synergy_Bliss, Blockis FROM combination WHERE id=$1", p.ID).Scan(&p.DrugA, &p.DrugB, &p.CellLine, &p.Source, &p.DSS, &p.Synergy_Bliss, &p.Blockis)
}

func (p *combination) updateCombination(db *sql.DB) error {
	_, err := db.Exec("UPDATE combination SET DrugA=$1, DrugB=$2, CellLine=$3, Source=$4, DSS=$5, Synergy_Bliss=$6 WHERE id=$7", p.DrugA, p.DrugB, p.CellLine, p.Source, p.DSS, p.Synergy_Bliss, p.ID)
	return err
}

func (p *combination) deleteCombination(db *sql.DB) error {
	_, err := db.Exec("DELETE from combination WHERE id=$1", p.ID)
	return err
}

func (p *combination) createCombination(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO combination(DrugA, DrugB, CellLine, Source, DSS, Synergy_Bliss, Blockis) VALUES($1, $2, $3, $4, $5, $6) RETURNING id", p.DrugA, p.DrugB, p.CellLine, p.Source, p.DSS, p.Synergy_Bliss, p.Blockis).Scan(&p.ID)

	if err != nil {
		return err
	}
	return nil
}

func getCombinations(db *sql.DB, start, count int) ([]combination, error) {
	rows, err := db.Query(
		"SELECT ID, DrugA, DrugB, CellLine, Source, DSS, Synergy_Bliss, Blockis FROM combination LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	allCombinations := []combination{}

	for rows.Next() {
		var p combination
		if err := rows.Scan(&p.ID, &p.DrugA, &p.DrugB, &p.CellLine, &p.Source, &p.DSS, &p.Synergy_Bliss,&p.Blockis); err != nil {
			return nil, err
		}
		allCombinations = append(allCombinations, p)
	}
	return allCombinations, nil
}


func getConditions(db *sql.DB, start, count int) ([]conditions, error) {
	rows, err := db.Query(
		"SELECT ID, Source, Readout, Compound, ControlNegative, ControlPositive, Single, DosesSingle,Combination, DosesCombination,Format,CellsPerWell,VolumePerWell,Unit FROM conditions LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	allConditions := []conditions{}

	for rows.Next() {
		var p conditions
		if err := rows.Scan(&p.ID, &p.Source, &p.AssayDetails.Readout, &p.AssayDetails.Compound, &p.AssayDetails.ControlNegative, &p.AssayDetails.ControlPositive,
			&p.NumberOfDrugs.Single, &p.NumberOfDrugs.DosesSingle,&p.NumberOfDrugs.Combination,&p.NumberOfDrugs.DosesCombination, &p.Plate.Format, &p.Plate.CellsPerWell, &p.Plate.VolumePerWell, &p.Plate.Unit);
			err != nil {
				return nil, err
			}
		allConditions = append(allConditions, p)
	}
	return allConditions, nil
}

