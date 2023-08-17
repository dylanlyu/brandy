package uuid

import (
	"errors"
	"github.com/google/uuid"
)

func PointerUUID(input uuid.UUID) *uuid.UUID { return &input }

func StringToUUID(input string) uuid.UUID {
	return uuid.MustParse(input)
}

func ValidateUUID(input string) (uuid.UUID, error) {
	id, _ := uuid.Parse(input)
	if id.String() == "00000000-0000-0000-0000-000000000000" {
		return id, errors.New("string is not uuid")
	}

	return id, nil
}

func GenerateUUID() string {
	return uuid.New().String()
}
