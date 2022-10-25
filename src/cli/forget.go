package cli

import (
	"D7024E/node/id"
	"D7024E/node/stored"
)

func Forget(input string) string {
	valueID, err := id.String2KademliaID(input)
	if err != nil {
		return err.Error()
	}
	store := stored.GetInstance()
	store.StopRefresh(*valueID)
	return valueID.String()
}
