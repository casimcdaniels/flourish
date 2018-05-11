package flourish

type StrainService struct {
	Strains StrainRepository
}

func (s StrainService) Get(id uint64) (*Strain, error) {
	return s.Strains.Get(id)
}

func (s StrainService) Create(name string, race string, flavors []string, effects StrainEffects) (*Strain, error) {
	strain := &Strain{
		Name:    name,
		Race:    race,
		Flavors: flavors,
		Effects: effects,
	}

	err := s.Strains.Create(strain)
	if err != nil {
		return nil, err
	}

	return strain, nil
}

func (s StrainService) Update(id uint64, name *string, race *string, flavors *[]string, effects *StrainEffects) error {
	strain, err := s.Get(id)
	if err != nil {
		return err
	}

	if name != nil {
		strain.Name = *name
	}

	if race != nil {
		strain.Race = *race
	}

	if flavors != nil {
		strain.Flavors = *flavors
	}

	if effects != nil {
		strain.Effects = *effects
	}

	err = s.Strains.Save(strain)

	if err != nil {
		return err
	}

	return nil
}

func (s StrainService) Remove(id uint64) error {
	return s.Strains.Delete(id)
}

func (s StrainService) Search(options StrainSearchOptions) ([]*Strain, error) {
	return s.Strains.Search(options)
}
