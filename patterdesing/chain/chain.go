package main

import (
	"fmt"
)

type Modifier interface {
	Add(m Modifier, ident int)
	Handle()
}

type Creature struct {
	Name            string
	Attack, Defense int
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

func NewCreature(name string, attack, defense int) *Creature {
	return &Creature{name, attack, defense}
}

type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func (c *CreatureModifier) Add(m Modifier, ident int) {
	if c.next != nil {
		c.next.Add(m, ident)
		return
	}

	c.next = m
}

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

func NewCreatureModifier(creature *Creature) Modifier {
	return &CreatureModifier{creature: creature}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func (d *DoubleAttackModifier) Handle() {
	d.creature.Attack *= 2
	fmt.Printf("Doubling %v attack %v \n", d.creature.Name, d.creature.Attack)
	d.CreatureModifier.Handle()
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{creature: c}}
}

type IncreaseDefenseModifier struct {
	CreatureModifier
}

func (d *IncreaseDefenseModifier) Handle() {
	fmt.Printf("Doubling Defense %v attack %v \n", d.creature.Name, d.creature.Defense*2)
	d.CreatureModifier.Handle()
}

func NewIncraseDefeseModifier(c *Creature) *IncreaseDefenseModifier {
	return &IncreaseDefenseModifier{CreatureModifier{creature: c}}
}

type NoBonusesModifier struct {
	CreatureModifier
}

func NewNoBonusesModifier(
	c *Creature) *NoBonusesModifier {
	return &NoBonusesModifier{CreatureModifier{
		creature: c}}
}

func (n *NoBonusesModifier) Handle() {
	// nothing here!
}

func main() {
	gobling := NewCreature("Gobling", 1, 1)
	fmt.Println(gobling.String())

	root := NewCreatureModifier(gobling)
	root.Add(NewIncraseDefeseModifier(gobling), 1)
	root.Add(NewNoBonusesModifier(gobling), 2)
	root.Add(NewDoubleAttackModifier(gobling), 2)
	root.Add(NewDoubleAttackModifier(gobling), 3)
	root.Add(NewDoubleAttackModifier(gobling), 4)
	root.Add(NewDoubleAttackModifier(gobling), 5)
	root.Handle()

	fmt.Println(gobling.String())
}
