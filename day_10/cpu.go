package main

type CallbackId struct {
	start     int
	increment int
}

type Callback func(int, int)

type CPU struct {
	x     int
	cycle int

	callbacks map[CallbackId]Callback
}

func (c *CPU) tick() {
	c.cycle += 1
	for id := range c.callbacks {
		c.tryTriggerCallback(id)
	}
}

func (c *CPU) tryTriggerCallback(id CallbackId) {
	d := c.cycle - id.start
	if d >= 0 && d%id.increment == 0 {
		c.callbacks[id](c.cycle, c.x)
	}
}

func (c *CPU) Execute(instruction Instruction) {
	for t := 0; t < instruction.CycleCount()-1; t += 1 {
		c.tick()
	}
	instruction.Execute(c)
	c.tick()
}

func (c *CPU) ExecuteMultiple(instructions []Instruction) {
	for _, i := range instructions {
		c.Execute(i)
	}
}

func (c *CPU) AddCallback(id CallbackId, f Callback) {
	c.callbacks[id] = f
	c.tryTriggerCallback(id)
}

func NewCPU() *CPU {
	return &CPU{1, 1, make(map[CallbackId]Callback)}
}
