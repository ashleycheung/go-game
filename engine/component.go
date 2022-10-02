package engine

import "github.com/ashleycheung/go-game/event"

// Used to add extra functionality
// to a game object
type Component interface {
	Step(delta float64)
	// Gets the game object
	// this component is added to
	GetGameObject() *GameObject

	// Used internally, and called
	// when game object attached initially.
	baseOnGameObjectAttach()
	baseOnGameObjectDetach()

	// Called when scene enter
	OnSceneEnter()

	// Called when scene exits
	OnSceneExit()

	// To be overriden by subclass
	// This is called when attached
	// to game object
	OnGameObjectAttach()

	// Called when detached from game object
	OnGameObjectDetach()

	// Returns whether the component is in the world yet.
	// A component is in the world if is attached to a game object
	// and the object is in the world
	IsInScene() bool

	// Calls this function when the component
	// is inside the game scene.
	// If already in the scene the function
	// is called instantly
	CallWhenInScene(fn func())

	// Sets the game object
	// internally used by game object
	setGameObject(obj *GameObject)
}

// The base component that all
// new components must extend
type BaseComponent struct {
	obj *GameObject
	// Stores all the functions needed
	// to be called once in the world
	inWorldFuncsQueue []func()

	// Removes the on scene enter listener
	onSceneEnterRemoveListener func()

	onSceneExitRemoveListener func()
}

func (b *BaseComponent) GetGameObject() *GameObject {
	return b.obj
}

func (b *BaseComponent) baseOnGameObjectAttach() {
	// Check if in the world now
	if b.IsInScene() {
		b.callInWorldFuncsQueue()
	} else {
		// Otherwise wait for gameobject to join scene
		b.onSceneEnterRemoveListener = b.obj.Event.AddOneTimeListener(
			OnSceneEnterEvent,
			func(e event.Event[GameObjectEvent]) error {
				b.callInWorldFuncsQueue()
				b.onSceneEnterRemoveListener = nil
				// Add one time listener for scene exit event
				b.onSceneExitRemoveListener = b.obj.Event.AddOneTimeListener(
					OnSceneExitEvent,
					func(e event.Event[GameObjectEvent]) error {
						b.OnSceneExit()
						return nil
					},
				)
				return nil
			},
		)
	}
}

func (b *BaseComponent) baseOnGameObjectDetach() {
	// If detached then remove the listener
	// These need to be called because detaching a component
	// should remove them from the scene
	if b.onSceneEnterRemoveListener != nil {
		b.onSceneEnterRemoveListener()
	}
	if b.onSceneExitRemoveListener != nil {
		b.onSceneExitRemoveListener()
	}
	// If still in scene and about
	// to be removed then call on scene exit
	if b.IsInScene() {
		b.OnSceneExit()
	}
}

// Called all the functions in the queue
func (b *BaseComponent) callInWorldFuncsQueue() {
	// Uninitialised so return
	if b.inWorldFuncsQueue == nil {
		return
	}
	// Call all funcs in queue and clear
	for _, f := range b.inWorldFuncsQueue {
		f()
	}
	b.inWorldFuncsQueue = []func(){}
}

func (b *BaseComponent) IsInScene() bool {
	return b.obj != nil && b.obj.World != nil
}

func (b *BaseComponent) CallWhenInScene(fn func()) {
	// If already in world call the function
	if b.IsInScene() {
		fn()
		return
	}
	// If slice is empty initialise
	if b.inWorldFuncsQueue == nil {
		b.inWorldFuncsQueue = []func(){}
	}
	// Not in function yet so place in buffer
	b.inWorldFuncsQueue = append(b.inWorldFuncsQueue, fn)
}

func (b *BaseComponent) setGameObject(obj *GameObject) {
	b.obj = obj
}

func (b *BaseComponent) Step(delta float64)  {}
func (b *BaseComponent) OnSceneEnter()       {}
func (b *BaseComponent) OnSceneExit()        {}
func (b *BaseComponent) OnGameObjectAttach() {}
func (b *BaseComponent) OnGameObjectDetach() {}
