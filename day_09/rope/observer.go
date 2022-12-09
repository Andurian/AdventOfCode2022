package rope

type Observer interface {
	StartMoving(Rope)
	AboutToExecute(Instruction)
	StateChanged(Rope)
	FinishedInstruction()
	FinishedMoving()
}
