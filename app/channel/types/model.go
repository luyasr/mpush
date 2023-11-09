package types

import (
	"encoding/json"
	"github.com/luyasr/mpush/app/common"
)

type Type struct {
	*common.Meta
	Name string `json:"name" gorm:"not null;uniqueIndex"`
}

func (t *Type) TableName() string {
	return "channel_types"
}

func (t *Type) String() string {
	bytes, _ := json.Marshal(t)
	return string(bytes)
}
