package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Position is used by pop to map your positions database table to your go code.
type Position struct {
	ID            uuid.UUID `json:"id" db:"id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	Asset         string    `json:"asset" db:"asset"`
	AssetType     string    `json:"asset_type" db:"asset_type"`
	Justification string    `json:"justification" db:"justification"`
	Bias          string    `json:"bias" db:"bias"`
	Expiration    time.Time `json:"expiration" db:"expiration"`
	User          User      `belongs_to:"user"`
	UserID        uuid.UUID `json:"user_id" db:"user_id"`
}

// String is not required by pop and may be deleted
func (p Position) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Positions is not required by pop and may be deleted
type Positions []Position

// String is not required by pop and may be deleted
func (p Positions) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Position) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Asset, Name: "Asset"},
		&validators.FuncValidator{
			Field:   p.Bias,
			Name:    "Bias",
			Message: "Bias must be long or short",
			Fn: func() bool {
				return p.Bias == "long" || p.Bias == "short"
			},
		},
		&validators.FuncValidator{
			Field:   p.AssetType,
			Name:    "AssetType",
			Message: "Incorrect Asset Type",
			Fn: func() bool {
				return p.AssetType == "stock" || p.AssetType == "index" || p.AssetType == "crypto" || p.AssetType == "etf"
			},
		},
	), err
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Position) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Position) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
