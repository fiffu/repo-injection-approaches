package repository

import "github.com/fiffu/repo-injection-approaches/examples/domain/pvi"

type EntityDBModel struct {
	ID    []byte
	State string
}

func Serialize(e pvi.PVIEntity) *EntityDBModel {
	return &EntityDBModel{
		ID:    []byte(e.ID),
		State: EncodeEntityState(e.State),
	}
}

func Deserialize(model *EntityDBModel) *pvi.PVIEntity {
	return &pvi.PVIEntity{
		ID:    string(model.ID),
		State: DecodeEntityState(model.State),
	}
}

func EncodeEntityState(s pvi.EntityState) string {
	switch s {
	case pvi.Available:
		return "available"
	case pvi.Consumed:
		return "consumed"
	case pvi.Deleted:
		return "deleted"
	}
	return ""
}

func DecodeEntityState(str string) pvi.EntityState {
	switch str {
	case "available":
		return pvi.Available
	case "consumed":
		return pvi.Consumed
	case "deleted":
		return pvi.Deleted
	}
	return pvi.Unknown
}
