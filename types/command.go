package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommandParamType int

type CommandParam struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Value       string `bson:"value" json:"value"`
}

type CommandTemplate struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID      primitive.ObjectID `bson:"userID" json:"userID"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`

	TemplateParams []CommandParam `bson:"templateParams" json:"templateParams"` // popup menu, editable empty string (probably with some default value)
	ConstantParams []CommandParam `bson:"constantParams" json:"constantParams"`
}

type CreateCommandTemplateParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	TemplateParams []CommandParam `json:"templateParams"`
	ConstantParams []CommandParam `json:"constantParams"`
}

type UpdateCommandTemplateParams struct {
	CreateCommandTemplateParams
}

func NewCommandTemplateFromParams(userID string, params CreateCommandTemplateParams) (*CommandTemplate, error) {
	ouserID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	return &CommandTemplate{
		UserID:         ouserID,
		Name:           params.Name,
		Description:    params.Description,
		TemplateParams: params.TemplateParams,
		ConstantParams: params.ConstantParams,
	}, nil
}
