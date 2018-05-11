package flourish

type Strain struct {
	Id      uint64 `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Race    string `db:"race" json:"race"`
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
