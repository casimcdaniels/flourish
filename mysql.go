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

func (r MysqlStrainRepository) Save(*Strain) error {
	panic("implement me")
}

func (r MysqlStrainRepository) Delete(id uint64) error {
	panic("implement me")
}

func (r MysqlStrainRepository) Search(options StrainSearchOptions) ([]*Strain, error) {
	panic("implement me")
}

func (r MysqlStrainRepository) Get(id uint64) (*Strain, error) {
	panic("implement me")
}