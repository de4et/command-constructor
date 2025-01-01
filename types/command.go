package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommandParamType int

const (
	TypeString CommandParamType = iota
	TypeEmpty
	TypeNameless // without name but with value. Omit Name? or leave it empty?
	TypePopupMenu
)

type CommandParam struct {
	Name         string           `bson:"name" json:"name"`
	Description  string           `bson:"description" json:"description"`
	Type         CommandParamType `bson:"type" json:"type"`
	Value        []string         `bson:"value" json:"value"` // popup meny - array of strings, string - empty array, empty - empty array
	DefaultValue string           `bson:"defaultValue" json:"defaultValue"`
	IsConstant   bool             `bson:"isconstant" json:"isconstant"`
}

type CommandTemplate struct {
	ID          string `bson:"_id,omitempty" json:"id,omitempty"`
	UserID      string `bson:"userID,omitempty" json:"userID,omitempty"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`

	CommandName    string         `bson:"commandName" json:"commandName"`
	TemplateParams []CommandParam `bson:"templateParams" json:"templateParams"` // popup menu, editable empty string (probably with some default value)
}

type CreateCommandTemplateParams struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`

	CommandName    string         `json:"commandName" validate:"required"`
	TemplateParams []CommandParam `json:"templateParams" validate:"required"`
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
		UserID:         ouserID.Hex(),
		Name:           params.Name,
		Description:    params.Description,
		CommandName:    params.CommandName,
		TemplateParams: params.TemplateParams,
	}, nil
}
