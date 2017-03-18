package GoGym

type GymService interface {
	Prepare(g *Gym)
	WhoIsYourBoss(g *Gym)
	CallYourBoss() *Gym
}
