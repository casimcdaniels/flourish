package flourish

type Strain struct {
	Id              uint64
	Race            string
	Flavors         []string
	Effects 		StrainEffects
}

type StrainEffects struct {
	Positive []string
	Negative []string
	Medical []string
}

// StrainFilter
type StrainSearchOptions struct {}

type StrainRepository interface {
	Create(*Strain) (error)
	Update(*Strain) error
	Delete(id uint64) error
	Search(options StrainSearchOptions) ([]*Strain, error)
}