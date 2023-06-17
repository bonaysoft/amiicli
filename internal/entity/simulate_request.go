package entity

type Mode string

const (
	ModeRandom  Mode = "random"
	ModeSpecify Mode = "specify"
)

type SimulateRequest struct {
	Mode  string
	Times int
}
