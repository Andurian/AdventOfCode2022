package observer

import "andurian/adventofcode/2022/day_09/rope"

type EmptyObserver struct{}

func (o *EmptyObserver) StartMoving(rope.Rope)           {}
func (o *EmptyObserver) AboutToExecute(rope.Instruction) {}
func (o *EmptyObserver) StateChanged(rope.Rope)          {}
func (o *EmptyObserver) FinishedInstruction()            {}
func (o *EmptyObserver) FinishedMoving()                 {}
