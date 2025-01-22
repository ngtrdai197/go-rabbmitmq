package helper

import "github.com/google/uuid"

func ValidateUUID(v string) bool {
	uuid, err := uuid.Parse(v)
	if err != nil {
		return false
	}
	return uuid.String() == v
}
