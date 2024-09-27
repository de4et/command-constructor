package db

import (
	"context"

	"github.com/de4et/command-constructor/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const commandColl = "commands"

type CommandStore interface {
	Dropper

	GetCommands(context.Context) ([]*types.CommandTemplate, error)
	InsertCommand(ctx context.Context, command *types.CommandTemplate) (*types.CommandTemplate, error)
	SearchCommand(ctx context.Context, userID string, name string) ([]*types.CommandTemplate, error)
	UpdateCommand(ctx context.Context, id string, command *types.CommandTemplate) error
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
	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var commands []*types.CommandTemplate
	if err := cur.All(ctx, &commands); err != nil {
		return nil, err
	}

	return commands, nil
}

func (s *MongoCommandStore) InsertCommand(ctx context.Context, command *types.CommandTemplate) (*types.CommandTemplate, error) {
	res, err := s.coll.InsertOne(ctx, command)
	if err != nil {
		return nil, err
	}

	command.ID = res.InsertedID.(primitive.ObjectID)
	return command, nil
}

func (s *MongoCommandStore) SearchCommand(ctx context.Context, userID string, name string) ([]*types.CommandTemplate, error) {
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	cur, err := s.coll.Find(context.Background(), bson.M{
		"name": bson.M{
			"$regex": primitive.Regex{
				Pattern: name,
				Options: "i",
			},
		},
		"userID": oid,
	})
	if err != nil {
		return nil, err
	}
	var commands = []*types.CommandTemplate{}
	if err := cur.All(context.Background(), &commands); err != nil {
		return nil, err
	}
	return commands, nil
}

func (s *MongoCommandStore) GetCommandByFilter(ctx context.Context, filter bson.M) (*types.CommandTemplate, error) {
	var command types.CommandTemplate
	if err := s.coll.FindOne(ctx, filter).Decode(&command); err != nil {
		return nil, err
	}
	return &command, nil
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
