package pvi

type IPVIRepo interface {
	Find(*Query) ([]*PVIEntity, error)
	FindOne(*Query) (*PVIEntity, error)
	Update(*PVIEntity, *Query, *UpdateValues) error
}
