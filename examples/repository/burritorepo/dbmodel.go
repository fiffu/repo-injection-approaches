package burritorepo

import "github.com/fiffu/repo-injection-approaches/examples/domain/burrito"

type BurritoDBModel struct {
	ID    []byte
	State string
}

func Serialize(e burrito.BurritoEntity) *BurritoDBModel {
	return &BurritoDBModel{
		ID:    []byte(e.ID),
		State: EncodeEntityState(e.State),
	}
}

func Deserialize(model *BurritoDBModel) *burrito.BurritoEntity {
	return &burrito.BurritoEntity{
		ID:    string(model.ID),
		State: DecodeEntityState(model.State),
	}
}

func EncodeEntityState(s burrito.EntityState) string {
	switch s {
	case burrito.Available:
		return "available"
	case burrito.Consumed:
		return "consumed"
	case burrito.Deleted:
		return "deleted"
	}
	return ""
}

func DecodeEntityState(str string) burrito.EntityState {
	switch str {
	case "available":
		return burrito.Available
	case "consumed":
		return burrito.Consumed
	case "deleted":
		return burrito.Deleted
	}
	return burrito.Unknown
}
