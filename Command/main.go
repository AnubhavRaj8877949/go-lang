package main

import "fmt"

type ICommand interface {
	execute()
	undo()
}

// ----------------------Light------------------------------
type Light struct{}

func (l *Light) On() {
	fmt.Println("Light Turn on")
}
func (l *Light) Off() {
	fmt.Println("Light Turn off")
}

type LightCommand struct {
	Light *Light
}

func (l *LightCommand) execute() {
	l.Light.On()
}

func (l *LightCommand) undo() {
	l.Light.Off()
}

// --------------------Fan-----------------------------------

type Fan struct{}

func (f *Fan) On() {
	fmt.Println("Fan Turn on")
}

func (f *Fan) Off() {
	fmt.Println("Fan Turn off")
}

type FanCommand struct {
	Fan *Fan
}

func (f *FanCommand) execute() {
	f.Fan.On()
}
func (f *FanCommand) undo() {
	f.Fan.Off()
}

// --------------------Ac----------------------------------
type Ac struct{}

func (a *Ac) On() {
	fmt.Println("Turn on AC")
}
func (a *Ac) Off() {
	fmt.Println("Turn off AC")
}

type ACCommand struct {
	Ac *Ac
}

func (a *ACCommand) execute() {
	a.Ac.On()
}
func (a *ACCommand) undo() {
	a.Ac.Off()
}

// ---------Remote------------------------------------
type Remote struct {
	command []ICommand
	pressed []bool
}

func NewRemote(slots int) *Remote {
	return &Remote{
		command: make([]ICommand, slots),
		pressed: make([]bool, slots),
	}
}
func (r *Remote) SetCommand(cmd ICommand, idx int) {
	if idx >= 0 && idx < len(r.command) {
		r.command[idx] = cmd
		r.pressed[idx] = false
	}

}
func (r *Remote) PressButton(idx int) {
	if idx < 0 || idx >= len(r.command) || r.command[idx] == nil {
		return
	}

	if r.pressed[idx] == true {
		r.command[idx].undo()
		r.pressed[idx] = false

	} else {
		r.command[idx].execute()
		r.pressed[idx] = true

	}

}

func main() {
	lightcommand := &LightCommand{
		Light: &Light{},
	}

	fancommand := &FanCommand{
		Fan: &Fan{},
	}

	accommand := &ACCommand{
		Ac: &Ac{},
	}

	remote := NewRemote(3)

	remote.SetCommand(lightcommand, 0)
	remote.SetCommand(fancommand, 1)
	remote.SetCommand(accommand, 2)

	remote.PressButton(0)
	remote.PressButton(1)
	remote.PressButton(2)
	// ----------------

	remote.PressButton(0)
	remote.PressButton(1)
	remote.PressButton(2)

}

//  real life use
// -> TextEditor (undo feature)
// ->PhotoShop  (undo feature)
// keyboard shortcut
// Encapsulate  a request  as an object  thereby letting  you parameterige  client  with different requesr
