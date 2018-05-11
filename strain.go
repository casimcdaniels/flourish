package flourish

type Strain struct {
	Id      uint64 `json:"id"`
	Name    string `json:"name"`
	Race    string `json:"race"`
	Flavors []string
	Effects StrainEffects
}

type StrainEffects struct {
	Positive []string
	Negative []string
	Medical  []string
}

// StrainFilter
type StrainSearchOptions struct{}

type StrainRepository interface {
	Create(*Strain) (error)
	Save(*Strain) error
	Delete(id uint64) error
	Search(options StrainSearchOptions) ([]*Strain, error)
	Get(id uint64) (*Strain, error)
}
