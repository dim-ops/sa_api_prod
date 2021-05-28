package database

import (
	"context"
	"encoding/json"
	"go-sa/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataInterface interface {
	Insert(models.Data) (models.Data, error)
	Update(string, interface{}) (models.DataUpdate, error)
	Delete(string) (models.DataDelete, error)
	Get(string) (models.Data, error)
	Search(interface{}) ([]models.Data, error)
}

type DataEntreprise struct {
	Ctx context.Context
	Col *mongo.Collection
}

func (c *DataEntreprise) Insert(docs models.Data) (models.Data, error) {
	data := models.Data{}

	res, err := c.Col.InsertOne(c.Ctx, docs)
	if err != nil {
		return data, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return c.Get(id)
}
func (c *DataEntreprise) Update(id string, update interface{}) (models.DataUpdate, error) {
	result := models.DataUpdate{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	data, err := c.Get(id)
	if err != nil {
		return result, err
	}
	var exist map[string]interface{}
	b, err := json.Marshal(data)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &exist)

	change := update.(map[string]interface{})
	for k := range change {
		if change[k] == exist[k] {
			delete(change, k)
		}
	}

	if len(change) == 0 {
		return result, nil
	}

	res, err := c.Col.UpdateOne(c.Ctx, bson.M{"_id": _id}, bson.M{"$set": change})
	if err != nil {
		return result, err
	}

	newData, err := c.Get(id)
	if err != nil {
		return result, err
	}

	result.ModifiedCount = res.ModifiedCount
	result.Result = newData
	return result, nil
}
func (c *DataEntreprise) Delete(id string) (models.DataDelete, error) {
	result := models.DataDelete{
		DeletedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	res, err := c.Col.DeleteOne(c.Ctx, bson.M{"_id": _id})
	if err != nil {
		return result, err
	}
	result.DeletedCount = res.DeletedCount
	return result, nil
}
func (c *DataEntreprise) Get(id string) (models.Data, error) {
	data := models.Data{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return data, err
	}

	err = c.Col.FindOne(c.Ctx, bson.M{"_id": _id}).Decode((&data))
	if err != nil {
		return data, err
	}

	return data, nil
}
func (c *DataEntreprise) Search(filter interface{}) ([]models.Data, error) {
	datas := []models.Data{}
	if filter == nil {
		filter = bson.M{}
	}

	cursor, err := c.Col.Find(c.Ctx, filter)
	if err != nil {
		return datas, err
	}

	for cursor.Next(c.Ctx) {
		row := models.Data{}
		cursor.Decode(&row)
		datas = append(datas, row)
	}

	return datas, nil
}
