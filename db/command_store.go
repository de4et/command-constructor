package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/de4et/command-constructor/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const commandColl = "commands"

type CommandStore interface {
	Dropper

	GetCommands(context.Context) ([]*types.CommandTemplate, error)
	GetCommandByID(ctx context.Context, id string) (*types.CommandTemplate, error)
	InsertCommand(ctx context.Context, command *types.CommandTemplate) (*types.CommandTemplate, error)
	SearchCommands(ctx context.Context, userID string, name string) ([]*types.CommandTemplate, error)
	UpdateCommand(ctx context.Context, id string, command *types.CommandTemplate) error
	IsExists(ctx context.Context, userID, id string) (bool, error)
	GetCommandsOfUser(ctx context.Context, id string) ([]*types.CommandTemplate, error)
	DeleteCommandByID(ctx context.Context, id string) error
}

type MongoCommandStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoCommandStore(client *mongo.Client, dbname string) *MongoCommandStore {
	return &MongoCommandStore{
		client: client,
		coll:   client.Database(dbname).Collection(commandColl),
	}
}

func (s *MongoCommandStore) GetCommands(ctx context.Context) ([]*types.CommandTemplate, error) {
	return s.getCommandsByFilter(ctx, bson.M{})
}

func (s *MongoCommandStore) GetCommandByID(ctx context.Context, id string) (*types.CommandTemplate, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.getCommandByFilter(ctx, bson.M{"_id": oid})
}

func (s *MongoCommandStore) GetCommandsOfUser(ctx context.Context, userID string) ([]*types.CommandTemplate, error) {
	return s.getCommandsByFilter(ctx, bson.M{
		"userID": userID,
	})
}

func (s *MongoCommandStore) InsertCommand(ctx context.Context, command *types.CommandTemplate) (*types.CommandTemplate, error) {
	res, err := s.coll.InsertOne(ctx, command)
	if err != nil {
		return nil, err
	}

	command.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return command, nil
}

func (s *MongoCommandStore) SearchCommands(ctx context.Context, userID string, name string) ([]*types.CommandTemplate, error) {
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	fmt.Println(oid)
	cur, err := s.coll.Find(ctx, bson.M{
		// "$text": bson.M{
		// 	"$search": name,
		// },
		"name": bson.M{
			"$regex": primitive.Regex{
				Pattern: name,
				Options: "i",
			},
		},
		"userID": userID,
	})
	if err != nil {
		return nil, err
	}
	var commands = []*types.CommandTemplate{}
	if err := cur.All(ctx, &commands); err != nil {
		return nil, err
	}
	return commands, nil
}

func (s *MongoCommandStore) getCommandByFilter(ctx context.Context, filter bson.M) (*types.CommandTemplate, error) {
	var command types.CommandTemplate
	if err := s.coll.FindOne(ctx, filter).Decode(&command); err != nil {
		return nil, err
	}
	return &command, nil
}

func (s *MongoCommandStore) getCommandsByFilter(ctx context.Context, filter bson.M) ([]*types.CommandTemplate, error) {
	cur, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	commands := []*types.CommandTemplate{}
	if err := cur.All(ctx, &commands); err != nil {
		return nil, err
	}

	return commands, nil
}

func (s *MongoCommandStore) UpdateCommand(ctx context.Context, id string, command *types.CommandTemplate) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	tobson, err := ConvertStructToBson(*command)
	if err != nil {
		return nil
	}
	// convert whole document instead of certain fields
	// update := bson.M{"$set": tobson}
	res, err := s.coll.ReplaceOne(ctx, bson.M{"_id": oid}, tobson)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (s *MongoCommandStore) Drop(ctx context.Context) error {
	return s.coll.Drop(ctx)
}

func (s *MongoCommandStore) IsExists(ctx context.Context, userID, id string) (bool, error) {
	u, err := s.GetCommandByID(ctx, id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}

	if u.UserID == userID {
		return true, nil
	}
	if u == nil {
		return false, nil
	}

	return true, nil

}

func (s *MongoCommandStore) DeleteCommandByID(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if _, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid}); err != nil {
		return err
	}
	return nil
}
