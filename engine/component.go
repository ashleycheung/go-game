package engine

// Used to add extra functionality
// to a game object
type Component interface {
	Step(delta float64)
	// Gets the game object
	// this component is added to
	GetGameObject() *GameObject

	// To be overriden by subclass
	// This is called when attached
	// to game object
	OnGameObjectAttach()

	// Called when detached from game object
	OnGameObjectDetach()

	// Sets the game object
	// internally used by game object
	setGameObject(obj *GameObject)
}

// The base component that all
// new components must extend
type BaseComponent struct {
	obj *GameObject
}

func (b *BaseComponent) GetGameObject() *GameObject {
	return b.obj
}

func (b *BaseComponent) OnGameObjectAttach() {}

func (b *BaseComponent) OnGameObjectDetach() {}

func (b *BaseComponent) setGameObject(obj *GameObject) {
	b.obj = obj
}

func (b *BaseComponent) Step(delta float64) {}
