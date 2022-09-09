package main

import (
	"fmt"
	"sync"
)

type Argument int

const (
	Attack Argument = iota
	Defense
)

type Observer interface {
	Handle(q *Query)
}

type Observable interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Fire(q *Query)
}

type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

type Game struct {
	observers sync.Map
}

func (g *Game) Subscribe(o Observer) {
	g.observers.Store(o, struct{}{})
}

func (g *Game) Unsubscribe(o Observer) {
	g.observers.Delete(o)
}

func (g *Game) Fire(q *Query) {
	g.observers.Range(func(key, value any) bool {
		if key == nil {
			return false
		}
		key.(Observer).Handle(q)
		return true
	})
}

type CreatureChain struct {
	game            *Game
	Name            string
	attack, defense int
}

func (c *CreatureChain) Attack() int {
	q := Query{c.Name, Attack, c.attack}
	c.game.Fire(&q)

	return q.Value
}

func (c *CreatureChain) Defense() int {
	q := Query{c.Name, Defense, c.defense}
	c.game.Fire(&q)

	return q.Value
}

func (c *CreatureChain) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.attack, c.defense)
}

func NewCreatureChan(game *Game, name string, attack int, defense int) *CreatureChain {
	return &CreatureChain{
		game,
		name,
		attack,
		defense,
	}
}

type CreatureModifierChain struct {
	game     *Game
	Creature *CreatureChain
}

func (c *CreatureModifierChain) Handle(q *Query) {}

type DoubleAttackModifierChain struct {
	CreatureModifierChain
}

func (c *DoubleAttackModifierChain) Handle(q *Query) {
	if q.CreatureName == c.Creature.Name && q.WhatToQuery == Attack {
		q.Value *= 2
	}
}

func (c *DoubleAttackModifierChain) Close() error {
	c.game.Unsubscribe(c)
	return nil
}

func NewDoubleAttackModifierChain(game *Game, creature *CreatureChain) *DoubleAttackModifierChain {
	d := &DoubleAttackModifierChain{CreatureModifierChain{game, creature}}
	game.Subscribe(d)
	return d
}

func main() {
	game := &Game{sync.Map{}}

	gobling := NewCreatureChan(game, "strong gobling", 2, 2)
	var count = 3
	fmt.Println(gobling.String())
	{
		m := NewDoubleAttackModifierChain(game, gobling)
		fmt.Println(gobling.String())
		_ = m.Close()
		count = 5
	}

	fmt.Println(count)
	fmt.Println(gobling.String())
}
