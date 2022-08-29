package engine

import (
	"fmt"

	"github.com/ashleycheung/go-game/event"
)

// Implements the base game object struct
type BaseGameObject struct {
	// Game object id
	// It will not have an id
	// until placed in the world
	Id int

	// Event manager
	Event *event.EventManager

	// Reference to the game world.
	// This will be set to nil if
	// the object is not in the world
	World *GameWorld

	// Parent object
	Parent GameObject

	// All children of the object
	Children []GameObject
}

func (g *BaseGameObject) GetEventManager() *event.EventManager {
	return g.Event
}

func (g *BaseGameObject) SetEventManager(m *event.EventManager) {
	g.Event = m
}

// This is called every step by the game
func (g *BaseGameObject) Step(delta float64) {}

// Returns whether the game object has
// been assigned a unique id yet. Essentially
// returns g.Id != 0
func (g *BaseGameObject) HasId() bool {
	return g.Id != 0
}

func (g *BaseGameObject) GetId() int {
	return g.Id
}

func (g *BaseGameObject) SetId(id int) {
	g.Id = id
}

// Adds a child to the tree
// and give it an id if it doesn't exist
func (g *BaseGameObject) AddChild(o GameObject) error {
	// If this game object is already
	// in the world add the given game object
	// and all its children
	if g.World != nil {
		objIter := NewBFSIterator(o)
		for objIter.HasNext() {
			// Check object isnt already in the world
			nextObj := objIter.Next()
			if nextObj.GetWorld() != nil {
				return fmt.Errorf("obj is already in the world %v", nextObj)
			}
			nextObj.SetWorld(g.World)

			g.World.idIncrement++
			// Set the id and world
			nextObj.SetId(g.World.idIncrement)

			// Call enter event
			nextObj.GetEventManager().EmitEvent(event.Event{Name: string(OnSceneEnter)})
		}
	}
	g.Children = append(g.Children, o)
	return nil
}

func (g *BaseGameObject) RemoveChild(o GameObject) {
	// If this game object is in the world
	// remove the world pointer of all children
	// of the object removed
	if g.World != nil {
		objIter := NewBFSIterator(o)
		for objIter.HasNext() {
			// Remove world reference
			nextObj := objIter.Next()
			nextObj.SetWorld(nil)

			// Call exit event
			nextObj.GetEventManager().EmitEvent(event.Event{Name: string(OnSceneExit)})
		}
	}
	// Remove the child
	newChildren := []GameObject{}
	for _, c := range g.Children {
		if c != o {
			newChildren = append(newChildren, c)
		}
	}
	g.SetChildren(newChildren)
}

func (g *BaseGameObject) GetWorld() *GameWorld {
	return g.World
}

func (g *BaseGameObject) SetWorld(w *GameWorld) {
	g.World = w
}

func (g *BaseGameObject) GetChildren() []GameObject {
	return g.Children
}

func (g *BaseGameObject) SetChildren(children []GameObject) {
	g.Children = children
}
