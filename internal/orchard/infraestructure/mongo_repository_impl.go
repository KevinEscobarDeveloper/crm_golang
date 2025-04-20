package infraestructure

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"mango_crm/internal/orchard/domain/entity"
	"mango_crm/internal/orchard/domain/repository"
	"mango_crm/internal/orchard/infraestructure/model"
	"mango_crm/pkg/logger"
	"mango_crm/pkg/mappers"
)

type MongoRepositoryImpl struct {
	collection *mongo.Collection
	logger     *logger.Logger
}

func NewMongoRepositoryImpl(col *mongo.Collection, logger *logger.Logger) repository.Repository {
	return &MongoRepositoryImpl{collection: col, logger: logger}
}
func (m *MongoRepositoryImpl) Save(ctx context.Context, o *entity.Orchard) error {
	_, err := m.collection.InsertOne(ctx, mappers.ToModel(o))
	if err != nil {
		m.logger.Error("Error saving orchard: " + err.Error())
		return err
	}
	return nil
}

func (m *MongoRepositoryImpl) FindById(ctx context.Context, id string) (*entity.Orchard, error) {
	filter := bson.M{"_id": id}

	var orchard model.Orchard
	err := m.collection.FindOne(ctx, filter).Decode(&orchard)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			m.logger.Warn("Orchard not found with ID: " + id)
			return nil, nil
		}
		m.logger.Error("Error finding orchard by ID: " + err.Error())
		return nil, err
	}

	return mappers.ToEntity(orchard), nil
}
func (m *MongoRepositoryImpl) FindAll(ctx context.Context) ([]*entity.Orchard, error) {
	var orchardModels []model.Orchard
	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		m.logger.Error("Error finding all orchards: " + err.Error())
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			m.logger.Error("Error closing cursor: " + err.Error())
		}
	}(cursor, ctx)

	if err = cursor.All(ctx, &orchardModels); err != nil {
		m.logger.Error("Error decoding orchards: " + err.Error())
		return nil, err
	}

	var orchards []*entity.Orchard
	for _, m := range orchardModels {
		orchards = append(orchards, mappers.ToEntity(m))
	}

	return orchards, nil
}

func (m *MongoRepositoryImpl) Update(ctx context.Context, o *entity.Orchard) error {
	_, err := m.collection.UpdateOne(ctx, bson.M{"_id": o.ID}, bson.M{"$set": mappers.ToModelWithoutID(o)})
	if err != nil {
		m.logger.Error("Error updating orchard: " + err.Error())
		return err
	}
	return nil
}

func (m *MongoRepositoryImpl) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	_, err := m.collection.DeleteOne(ctx, filter)
	if err != nil {
		m.logger.Error("Error deleting orchard: " + err.Error())
		return err
	}
	return nil
}
