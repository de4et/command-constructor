package db

import (
	"context"
	"errors"

	"github.com/de4et/command-constructor/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

type UserStore interface {
	Dropper

	GetUserByName(ctx context.Context, name string) (*types.User, error)
	GetUserByID(ctx context.Context, id string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(ctx context.Context, user *types.User) (*types.User, error)
	DeleteUserByID(ctx context.Context, id string) error
	IsExist(ctx context.Context, name string) (bool, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, dbname string) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(dbname).Collection(userColl),
	}
}

func (s *MongoUserStore) Search(part string) error {
	cur, err := s.coll.Find(context.Background(), bson.M{
		"name": bson.M{
			"$regex": primitive.Regex{
				Pattern: part,
				Options: "i",
			},
		},
	})
	if err != nil {
		return err
	}
	var users []*types.User
	if err := cur.All(context.Background(), &users); err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) DeleteUserByID(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if _, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid}); err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.GetUserByFilter(ctx, bson.M{"_id": oid})
}

func (s *MongoUserStore) GetUserByName(ctx context.Context, name string) (*types.User, error) {
	return s.GetUserByFilter(ctx, bson.M{"name": name})
}

func (s *MongoUserStore) GetUserByFilter(ctx context.Context, filter bson.M) (*types.User, error) {
	var user types.User
	if err := s.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	users := []*types.User{}
	if err := cur.All(ctx, &users); err != nil {
		return []*types.User{}, err
	}

	return users, nil
}

func (s *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return user, nil
}

func (s *MongoUserStore) IsExist(ctx context.Context, name string) (bool, error) {
	u, err := s.GetUserByName(ctx, name)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	if u == nil {
		return false, nil
	}

	return true, nil

}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	return s.coll.Drop(ctx)
}
