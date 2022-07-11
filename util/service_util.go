package util

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateId(key string) string {

	id := fmt.Sprintf("%v%v", key, uuid.New().String())

	return id
}
