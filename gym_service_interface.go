package GoGym

import (
	"reflect"
)

// GymService is a service interface
type GymService interface {
	Prepare(g *Gym)
	WhoIsYourBoss(g *Gym)
	CallYourBoss() *Gym
	CallMethod(method string, param []interface{}) []reflect.Value
}
