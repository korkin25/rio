package mappers

import (
	"github.com/rancher/norman/pkg/types"
	"github.com/rancher/norman/pkg/types/mapper"
)

func NewObject(kind, apiVersion string) types.Mapper {
	return types.Mappers{
		mapper.SetValue{Field: "kind", InternalValue: kind},
		mapper.SetValue{Field: "apiVersion", InternalValue: apiVersion},
		mapper.Drop{
			Field: "kind",
		},
		mapper.Drop{
			Field: "apiVersion",
		},
		mapper.Move{
			From: "metadata/labels",
			To:   "labels",
		},
		mapper.Move{
			From: "metadata/annotations",
			To:   "annotations",
		},
		mapper.Drop{
			Field: "metadata",
		},
		mapper.Drop{
			Optional: true,
			Field:    "status",
		},
		&mapper.Embed{
			Field:    "spec",
			Optional: true,
		},
		LabelCleaner{},
	}
}
