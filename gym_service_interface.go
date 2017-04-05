package GoGym

import (
	"reflect"
)

// GymService is a service interface
type GymService interface {
	Prepare(g *Gym)
	InjectServiceContainer(g *Gym)
	GetServiceContainer() *Gym
	CallMethod(method string, param []interface{}) []reflect.Value
}
