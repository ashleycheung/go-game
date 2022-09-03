package engine

import (
	"fmt"

	"github.com/ashleycheung/go-game/event"
)

// Implements the base game object struct
type baseGameObject struct {
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

	// Local storage of all the groups
	// this object is a part of
	groupsSet map[string]bool
}

func (g *baseGameObject) GetEventManager() *event.EventManager {
	return g.Event
}

func (g *baseGameObject) SetEventManager(m *event.EventManager) {
	g.Event = m
}

// This is called every step by the game
func (g *baseGameObject) Step(delta float64) {}

// Returns whether the game object has
// been assigned a unique id yet. Essentially
// returns g.Id != 0
func (g *baseGameObject) HasId() bool {
	return g.Id != 0
}

func (g *baseGameObject) GetId() int {
	return g.Id
}

func (g *baseGameObject) SetId(id int) {
	g.Id = id
}

// Adds a child to the tree
// and give it an id if it doesn't exist
func (g *baseGameObject) AddChild(o GameObject) error {
	// If this game object is already
	// in the world add the given game object
	// and all its children
	if g.World != nil {
		objIter := newBFSIterator(o)
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

			// Add groups to world cache
			for _, groupName := range nextObj.GetGroups() {
				g.World.addObjectToGroup(nextObj, groupName)
			}

			// Call enter event
			nextObj.GetEventManager().EmitEvent(event.Event{Name: string(OnSceneEnter)})
		}
	}
	g.Children = append(g.Children, o)
	return nil
}

func (g *baseGameObject) RemoveChild(o GameObject) {
	// If this game object is in the world
	// remove the world pointer of all children
	// of the object removed
	if g.World != nil {
		objIter := newBFSIterator(o)
		for objIter.HasNext() {
			nextObj := objIter.Next()

			// Call exit event
			nextObj.GetEventManager().EmitEvent(event.Event{Name: string(OnSceneExit)})

			// Remove groups from world cache
			for _, groupName := range nextObj.GetGroups() {
				g.World.removeObjectFromGroup(nextObj, groupName)
			}

			// Remove world reference
			nextObj.SetWorld(nil)
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

func (g *baseGameObject) AddToGroup(groupName string) {
	g.groupsSet[groupName] = true
	if g.World != nil {
		g.World.addObjectToGroup(g, groupName)
	}
}

func (g *baseGameObject) RemoveFromGroup(groupName string) {
	delete(g.groupsSet, groupName)
	if g.World != nil {
		g.World.removeObjectFromGroup(g, groupName)
	}
}

func (g *baseGameObject) GetGroups() []string {
	outGroups := []string{}
	for group := range g.groupsSet {
		outGroups = append(outGroups, group)
	}
	return outGroups
}

func (g *baseGameObject) GetWorld() *GameWorld {
	return g.World
}

func (g *baseGameObject) SetWorld(w *GameWorld) {
	g.World = w
}

func (g *baseGameObject) GetChildren() []GameObject {
	return g.Children
}

func (g *baseGameObject) SetChildren(children []GameObject) {
	g.Children = children
}
