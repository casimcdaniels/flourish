package flourish

import "github.com/jmoiron/sqlx"

type MysqlStrainRepository struct {
	DB *sqlx.DB
}

func (r MysqlStrainRepository) Create(s *Strain) (error) {
	tx, err  := r.DB.Begin()

	if err != nil {
		return err
	}

	res, err := tx.Exec(`INSERT INTO strains (name, race) VALUES (?, ?)`, s.Name, s.Race)

	if err != nil {
		tx.Rollback()

		return err
	}

	lastInserted, err := res.LastInsertId()

	if err != nil {
		tx.Rollback()
		return err
	}

	strainId := uint64(lastInserted)

	s.Id = strainId

	for _, flavor := range s.Flavors {
		_, err = tx.Exec(`
			INSERT INTO strain_flavors (strain_id, flavor_id) VALUES (?, (SELECT id FROM flavors WHERE name = ?))
		`, strainId, flavor)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, effect := range s.Effects.Positive {
		_, err = tx.Exec(`INSERT INTO strain_effects (strain_id, effect_id) VALUES (?, (SELECT id FROM effects WHERE name = ? ))`, strainId, effect)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, effect := range s.Effects.Negative {
		_, err = tx.Exec(`INSERT INTO strain_effects (strain_id, effect_id) VALUES (?, (SELECT id FROM effects WHERE name = ? ))`, strainId, effect)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, effect := range s.Effects.Medical {
		_, err = tx.Exec(`INSERT INTO treatments (strain_id, symptom_id) VALUES (?, (SELECT id FROM symptoms WHERE name = ? ))`, strainId, effect)
		if err != nil {
			tx.Rollback()
			return err
		}
	}


	err = tx.Commit()
	return nil
}

func (r MysqlStrainRepository) Save(s *Strain) error {
	tx, err  := r.DB.Begin()

	if err != nil {
		return err
	}

	// Go ahead and update the race/name in the event it was changed
	_, err = tx.Exec(`
		UPDATE strains SET
		race = ?,
		name = ?
		WHERE id = ?
	`, s.Race, s.Name, s.Id)

	if err != nil {
		tx.Rollback()
		return err
	}

	// Since there's not a cache/relationship structure in the application, we will delete all
	// associated records then re-insert
	tx.Exec("DELETE FROM strain_effects WHERE strain_id = ?", s.Id)
	tx.Exec("DELETE FROM strain_flavors WHERE strain_id = ?", s.Id)
	tx.Exec("DELETE FROM treatments WHERE strain_id = ?", s.Id)

	if err !=  nil {
		tx.Rollback()
		return err
	}

	for _, flavor := range s.Flavors {
		_, err = tx.Exec(`
			INSERT INTO strain_flavors (strain_id, flavor_id) VALUES (?, (SELECT id FROM flavors WHERE name = ?))
		`, s.Id, flavor)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, effect := range s.Effects.Positive {
		_, err = tx.Exec(`INSERT INTO strain_effects (strain_id, effect_id) VALUES (?, (SELECT id FROM effects WHERE name = ? ))`, s.Id, effect)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, effect := range s.Effects.Negative {
		_, err = tx.Exec(`INSERT INTO strain_effects (strain_id, effect_id) VALUES (?, (SELECT id FROM effects WHERE name = ? ))`, s.Id, effect)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, effect := range s.Effects.Medical {
		_, err = tx.Exec(`INSERT INTO treatments (strain_id, symptom_id) VALUES (?, (SELECT id FROM symptoms WHERE name = ? ))`, s.Id, effect)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// Removes the strain from the database, and will cascade to any associative tables
func (r MysqlStrainRepository) Delete(id uint64) error {
	_, err := r.DB.Exec("DELETE FROM strains WHERE id = ?", id)
	return err
}

// Fetches records within the strains table based upon the provided optios
func (r MysqlStrainRepository) Search(options StrainSearchOptions) ([]*Strain, error) {
	panic("implement me")
}


// Fetches the strain from the repository and builds the strain object from associations
func (r MysqlStrainRepository) Get(id uint64) (*Strain, error) {

	var strain Strain
	err := r.DB.Get(&strain, "SELECT * FROM strains WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	var flavors []string

	err = r.DB.Select(&flavors, "SELECT name FROM flavors JOIN strain_flavors ON flavors.id = strain_flavors.flavor_id AND strain_flavors.strain_id = ?", id)
	if err != nil {
		return nil, err
	}

	var positiveEffects []string
	err = r.DB.Select(&positiveEffects, "SELECT name FROM effects JOIN strain_effects ON effects.id = strain_effects.effect_id AND strain_effects.strain_id = ? AND effects.rating = 'positive'", id)
	if err != nil {
		return nil, err
	}

	var negativeEffects []string
	err = r.DB.Select(&negativeEffects, "SELECT name FROM effects JOIN strain_effects ON effects.id = strain_effects.effect_id AND strain_effects.strain_id = ? AND effects.rating = 'negative'", id)
	if err != nil {
		return nil, err
	}

	var medicalTreatments []string
	err = r.DB.Select(&medicalTreatments, "SELECT name FROM symptoms JOIN treatments ON symptoms.id = treatments.symptom_id AND treatments.strain_id = ?", id)
	if err != nil {
		return nil, err
	}

	effects := StrainEffects{
		Positive: positiveEffects,
		Negative: negativeEffects,
		Medical: medicalTreatments,
	}

	strain.Flavors = flavors
	strain.Effects = effects
	return &strain, nil


}