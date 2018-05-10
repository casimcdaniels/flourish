package flourish

type StrainService struct {
	Strains StrainRepository
}

func (s StrainService) Create(race string, flavors []string, effects StrainEffects) (*Strain, error) {}
func (s StrainService) Update(id uint64) error { return nil }
func (s StrainService) Remove(id uint64) error { return nil }
func (s StrainService) Search(options StrainSearchOptions) ([]*Strain, error) { return nil, nil }