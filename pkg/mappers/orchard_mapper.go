package mappers

import (
	"mango_crm/internal/orchard/domain/entity"
	"mango_crm/internal/orchard/infraestructure/model"
	"time"
)

func ToEntity(m model.Orchard) *entity.Orchard {
	return &entity.Orchard{
		ID:       m.ObjectID,
		Name:     m.Name,
		Location: m.Location,
		Owner:    m.Owner,
		Active:   m.Active,
		CreateAt: m.CreatedAt,
	}
}

func ToModel(e *entity.Orchard) model.Orchard {
	return model.Orchard{
		ObjectID:  e.ID,
		Name:      e.Name,
		Location:  e.Location,
		Owner:     e.Owner,
		Active:    e.Active,
		CreatedAt: time.Now(),
	}
}

func ToModelWithoutID(e *entity.Orchard) map[string]interface{} {
	return map[string]interface{}{
		"name":       e.Name,
		"location":   e.Location,
		"owner":      e.Owner,
		"active":     e.Active,
		"created_at": e.CreateAt,
	}
}
