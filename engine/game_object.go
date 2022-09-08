package engine

import (
	"fmt"

	"github.com/ashleycheung/go-game/event"
)

type GameEvent string

const (
	// Called when a body first enters the scene
	OnSceneEnterEvent GameEvent = "onSceneEnter"

	// Called when the body leaves the scene
	OnSceneExitEvent GameEvent = "onSceneExit"

	// Called when a step begins for the given game object
	OnGameObjectStepEvent GameEvent = "onGameObjectStepEvent"
)

// Creates a new game object
func NewGameObject() *GameObject {
	obj := &GameObject{}
	obj.groupsSet = map[string]bool{}
	obj.Children = []*GameObject{}
	obj.Event = event.NewEventManager()
	obj.components = map[string]Component{}
	return obj
}

// Implements the base game object struct
type GameObject struct {
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
	Parent *GameObject

	// All children of the object
	Children []*GameObject

	// Local storage of all the groups
	// this object is a part of
	groupsSet map[string]bool

	// Maps the name of the component
	// to the component itself.
	// Components should be added by reference
	components map[string]Component
}

// This is called every step by the game
func (g *GameObject) Step(delta float64) {
	g.Event.EmitEvent(event.Event{
		Name: string(OnGameObjectStepEvent),
	})
	// Call components
	for _, c := range g.components {
		c.Step(delta)
	}
}

// Returns whether the game object has
// been assigned a unique id yet. Essentially
// returns g.Id != 0
func (g *GameObject) HasId() bool {
	return g.Id != 0
}

// Adds a child to the tree
// and give it an id if it doesn't exist
func (g *GameObject) AddChild(o *GameObject) error {
	// If this game object is already
	// in the world add the given game object
	// and all its children
	if g.World != nil {
		objIter := newBFSIterator(o)
		for objIter.HasNext() {
			// Check object isnt already in the world
			nextObj := objIter.Next()
			if nextObj.World != nil {
				return fmt.Errorf("obj is already in the world %v", nextObj)
			}
			nextObj.World = g.World

			g.World.idIncrement++
			// Set the id and world
			nextObj.Id = g.World.idIncrement

			// Add groups to world cache
			for _, groupName := range nextObj.GetGroups() {
				g.World.addObjectToGroup(nextObj, groupName)
			}

			// Call enter event
			nextObj.Event.EmitEvent(event.Event{Name: string(OnSceneEnterEvent)})
		}
	}
	g.Children = append(g.Children, o)
	return nil
}

// Removes a child of this object
func (g *GameObject) RemoveChild(o *GameObject) {
	// If this game object is in the world
	// remove the world pointer of all children
	// of the object removed
	if g.World != nil {
		objIter := newBFSIterator(o)
		for objIter.HasNext() {
			nextObj := objIter.Next()

			// Call exit event
			nextObj.Event.EmitEvent(event.Event{Name: string(OnSceneExitEvent)})

			// Remove groups from world cache
			for _, groupName := range nextObj.GetGroups() {
				g.World.removeObjectFromGroup(nextObj, groupName)
			}

			// Remove world reference
			nextObj.World = nil
		}
	}
	// Remove the child
	newChildren := []*GameObject{}
	for _, c := range g.Children {
		if c != o {
			newChildren = append(newChildren, c)
		}
	}
	g.Children = newChildren
}

// Gets all the children of the given object
func (g *GameObject) GetChildren() []*GameObject {
	return g.Children
}

// Adds a component with the given name.
// If the component with name already exists,
// it is overriden
func (g *GameObject) AddComponent(name string, component Component) {
	g.components[name] = component
	component.setGameObject(g)
	component.OnGameObjectAttach()
}

// Removes the component of the given name
func (g *GameObject) RemoveComponent(name string) {
	delete(g.components, name)
}

// Gets a component of the given name, Returns nil if it doesnt exist
func (g *GameObject) GetComponent(name string) (c Component) {
	c, exists := g.components[name]
	if !exists {
		return nil
	}
	return c
}

// Adds object to a given group
func (g *GameObject) AddToGroup(groupName string) {
	g.groupsSet[groupName] = true
	if g.World != nil {
		g.World.addObjectToGroup(g, groupName)
	}
}

// Removes object from group
func (g *GameObject) RemoveFromGroup(groupName string) {
	delete(g.groupsSet, groupName)
	if g.World != nil {
		g.World.removeObjectFromGroup(g, groupName)
	}
}

// Gets all the groups the object is a part of
func (g *GameObject) GetGroups() []string {
	outGroups := []string{}
	for group := range g.groupsSet {
		outGroups = append(outGroups, group)
	}
	return outGroups
}

// Returns whether the game object is part of the given group
func (g *GameObject) InGroup(groupName string) bool {
	return g.groupsSet[groupName]
}
