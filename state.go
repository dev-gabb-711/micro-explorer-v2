package main

// GlobalState tracks data that survives scene transitions
type GlobalState struct {
	Health int // Player health
	Energy int // Player energy used for attacks/stuns
	Level  int // Current mission (1-10)
}

func NewGlobalState() *GlobalState {
	return &GlobalState{
		Health: 100,
		Energy: 250, // Default starting energy
		Level:  1,
	}
}
