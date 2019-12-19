package entity

// Entity is an interface describing the behaviors of something that can interact with the world state
type Entity interface {
	Update()
}
