package util

import (
	"fmt"
	"time"
)

type ServiceUtil struct{}

func (s *ServiceUtil) GenerateId(key string) string {

	now := time.Now().Format("20060102150405")
	id := fmt.Sprintf("%v%v", key, now)

	return id
}
