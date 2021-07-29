package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/wonpanu/my-golang/pkg/entity"
	"github.com/wonpanu/my-golang/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IVaccine interface {
	GetAll() ([]entity.Vaccine, error)
	GetByID(ID string) (entity.Vaccine, error)
	Create(vc entity.Vaccine) (entity.Vaccine, error)
	Update(ID string, vc entity.Vaccine) (entity.Vaccine, error)
	Delete(ID string) error
}

const (
	database   = "mymongo"
	collection = "vaccines"
)

type VaccineMongo struct {
	DB     *mongo.Database
	Client *mongo.Client
}

func (r VaccineMongo) Create(vc entity.Vaccine) (entity.Vaccine, error) {
	seed := time.Now().Unix()
	vc.ID = util.Hash(fmt.Sprint(seed))
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	_, err := r.DB.Collection(collection).InsertOne(ctx, vc)
	if err != nil {
		return entity.Vaccine{}, err
	}
	return vc, nil
}

func (r VaccineMongo) GetAll() ([]entity.Vaccine, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	cur, err := r.DB.Collection(collection).Find(ctx, bson.M{})
	var vaccine []entity.Vaccine
	if err != nil {
		return vaccine, err
	}
	cur.All(ctx, &vaccine)
	return vaccine, nil
}

func (r VaccineMongo) GetByID(ID string) (entity.Vaccine, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	filter := bson.M{"_id": ID}
	result := r.DB.Collection(collection).FindOne(ctx, filter)
	var vaccine entity.Vaccine
	result.Decode(&vaccine)
	return vaccine, nil
}

func (r VaccineMongo) Update(ID string, vc entity.Vaccine) (entity.Vaccine, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	filter := bson.M{"_id": ID}
	vc.ID = ID
	update := bson.M{
		"$set": vc,
	}
	_, err := r.DB.Collection(collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return entity.Vaccine{}, err
	}
	return vc, nil
}

func (r VaccineMongo) Delete(ID string) error {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	filter := bson.M{"_id": ID}
	_, err := r.DB.Collection(collection).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func NewVaccineRepo(mongo *mongo.Client) VaccineMongo {
	return VaccineMongo{
		DB: mongo.Database(database),
	}
}
