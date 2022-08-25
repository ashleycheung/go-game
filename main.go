package main

import (
	"github.com/ashleycheung/go-game/physics"
	"github.com/ashleycheung/go-game/physics/playground"
)

func main() {
	w := physics.NewWorld()
	// b1 := w.NewBody(physics.Circle{Radius: 20})
	// b1.Position = physics.Vector{X: 400, Y: 400}
	// w.AddBody(b1)

	// b2 := w.NewBody(physics.Circle{Radius: 20})
	// b2.Position = physics.Vector{X: 140, Y: 400}
	// b2.Velocity = physics.Vector{X: 30, Y: 0}
	// w.AddBody(b2)

	b1 := physics.NewBody(physics.Circle{Radius: 20})
	b1.Position = physics.Vector{X: 140, Y: 140}
	b1.Velocity = physics.Vector{X: -50, Y: 50}
	w.AddBody(b1)

	b2 := physics.NewBody(physics.Circle{Radius: 20})
	b2.Position = physics.Vector{X: 10, Y: 110}
	b2.Velocity = physics.Vector{X: 30, Y: 30}
	w.AddBody(b2)

	b3 := physics.NewBody(physics.Circle{Radius: 20})
	b3.Position = physics.Vector{X: 200, Y: 120}
	w.AddBody(b3)

	b4 := physics.NewBody(physics.Circle{Radius: 20})
	b4.Position = physics.Vector{X: 110, Y: 140}
	w.AddBody(b4)

	b5 := physics.NewBody(physics.Circle{Radius: 20})
	b5.Position = physics.Vector{X: 200, Y: 80}
	b5.Velocity = physics.Vector{X: -20, Y: 10}
	w.AddBody(b5)

	// physics.NewPhysicsServer(w, 8070)
	playground := playground.NewPhysicsPlayground(w)
	playground.Run(5050)
}
