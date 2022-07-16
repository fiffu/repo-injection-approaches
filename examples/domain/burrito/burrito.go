package burrito

type BurritoEntity struct {
	ID    string
	State EntityState
}

func (b *BurritoEntity) EntityID() string {
	return b.ID
}

type EntityState int

const (
	Available = iota
	Consumed
	Deleted
	Unknown
)

type Query struct {
	ID    string
	State EntityState
}

type UpdateValues struct {
	State EntityState
}
