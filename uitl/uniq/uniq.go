package uniq

import (
	uuid "github.com/satori/go.uuid"
)
var (
	num = make(chan uuid.UUID)
)

func init() {
	go func() {
		for {
			uuid:=uuid.NewV4()
			num <- uuid
		}
	}()
}

func GetUniq() uuid.UUID {
	return <-num
}
