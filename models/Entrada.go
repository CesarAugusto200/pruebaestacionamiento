package models

type Entrada struct {
	states      []string
	actualState string
}

func NewEntrada() *Entrada  {
	states := []string{"Entering", "Exiting", "Idle"}
	return &Entrada {
		states:      states,
		actualState: "Idle",
	}
}

func (e *Entrada ) GetState() string {
	return e.actualState
}

func (e *Entrada ) SetState(n int) {
	if n >= 0 && n < len(e.states) {
		e.actualState = e.states[n]
	}
}
