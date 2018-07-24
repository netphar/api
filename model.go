package main

import (
	"database/sql"
)

type doses struct {
	ID          int     `json:"id"`
	DrugA       string  `json:"DrugA"`
	DrugB       string  `json:"DrugB"`
	DoseA       float64 `json:"DoseA"`
	DoseB       float64 `json:"DoseB"`
	Response    float64 `json:"Response"`
	DSS         float64 `json:"DSS"`
	Synergy_HSA float64 `json:"Synergy_HSA"`
	CellLine    string  `json:"CellLine"`
}

type combination struct {
	ID          int     `json:"id"`
	DrugA       string  `json:"DrugA"`
	DrugB       string  `json:"DrugB"`
	CellLine    string  `json:"CellLine"`
	Source      string  `json:"Source"`
	DSS         float64 `json:"DSS"`
	Synergy_HSA float64 `json:"Synergy_HSA"`
}

func (p *doses) getDose(db *sql.DB) error {
	return db.QueryRow("SELECT DrugA, DrugB, DoseA, DoseB, Response, DSS, Synergy_HSA, CellLine FROM doses WHERE id=$1",
		p.ID).Scan(&p.DrugA, &p.DrugB, &p.DoseA, &p.DoseB, &p.Response, &p.DSS, &p.Synergy_HSA, &p.CellLine)
}

func (p *doses) updateDose(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE doses SET DrugA=$1, DrugB=$2, DoseA=$3, DoseB=$4, Response=$5, DSS=$6, Synergy_HSA=$7, CellLine=$8 WHERE id=$9",
			p.DrugA, p.DrugB, p.DoseA, p.DoseB, p.Response, p.DSS, p.Synergy_HSA, p.CellLine, p.ID)
	return err
}

func (p *doses) deleteDose(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM doses WHERE id=$1", p.ID)
	return err
}

func (p *doses) createDose(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO doses(DrugA, DrugB, DoseA, DoseB, Response, DSS, Synergy_HSA, CellLine) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		p.DrugA, p.DrugB, p.DoseA, p.DoseB, p.Response, p.DSS, p.Synergy_HSA, p.CellLine).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func getDoses(db *sql.DB, start, count int) ([]doses, error) {
	rows, err := db.Query(
		"SELECT ID, DrugA, DrugB, DoseA, DoseB, Response, DSS, Synergy_HSA, CellLine FROM doses LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	allDoses := []doses{}

	for rows.Next() {
		var p doses
		if err := rows.Scan(&p.ID, &p.DrugA, &p.DrugB, &p.DoseA, &p.DoseB, &p.Response, &p.DSS, &p.Synergy_HSA, &p.CellLine); err != nil {
			return nil, err
		}
		allDoses = append(allDoses, p)
	}
	return allDoses, nil
}

func (p *combination) getCombination(db *sql.DB) error {
	return db.QueryRow("SELECT DrugA, DrugB, CellLine, Source, DSS, Synergy_HSA FROM combination WHERE id=$1", p.ID).Scan(&p.DrugA, &p.DrugB, &p.CellLine, &p.Source, &p.DSS, &p.Synergy_HSA)
}

func (p *combination) updateCombination(db *sql.DB) error {
	_, err := db.Exec("UPDATE combination SET DrugA=$1, DrugB=$2, CellLine=$3, Source=$4, DSS=$5, Synergy_HSA=$6 WHERE id=$7", p.DrugA, p.DrugB, p.CellLine, p.Source, p.DSS, p.Synergy_HSA, p.ID)
	return err
}

func (p *combination) deleteCombination(db *sql.DB) error {
	_, err := db.Exec("DELETE from combination WHERE id=$1", p.ID)
	return err
}

func (p *combination) createCombination(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO combination(DrugA, DrugB, CellLine, Source, DSS, Synergy_HSA) VALUES($1, $2, $3, $4, $5, $6) RETURNING id", p.DrugA, p.DrugB, p.CellLine, p.Source, p.DSS, p.Synergy_HSA).Scan(&p.ID)
	if err != nil {
		return nil
	}
	return nil
}

func getCombinations(db *sql.DB, start, count int) ([]combination, error) {
	rows, err := db.Query(
		"SELECT ID, DrugA, DrugB, CellLine, Source, DSS, Synergy_HSA FROM combination LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	allCombinations := []combination{}

	for rows.Next() {
		var p combination
		if err := rows.Scan(&p.ID, &p.DrugA, &p.DrugB, &p.CellLine, &p.Source, &p.DSS, &p.Synergy_HSA); err != nil {
			return nil, err
		}
		allCombinations = append(allCombinations, p)
	}
	return allCombinations, nil
}
