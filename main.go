package main

import (
	"github.com/ashleycheung/go-game/physics"
	"github.com/ashleycheung/go-game/physics/playground"
)

func main() {
	w := physics.NewWorld()
	w.Config.Gravity.Y = 200

	b1 := physics.NewBody(physics.Circle{Radius: 20})
	b1.Position = physics.Vector{X: 140, Y: 500}
	w.AddBody(b1)

	b2 := physics.NewBody(physics.Circle{Radius: 20})
	b2.Position = physics.Vector{X: 160, Y: 500}
	w.AddBody(b2)

	b3 := physics.NewBody(physics.Circle{Radius: 20})
	b3.Position = physics.Vector{X: 200, Y: 150}
	w.AddBody(b3)

	b4 := physics.NewBody(physics.Circle{Radius: 10})
	b4.Position = physics.Vector{X: 110, Y: 400}
	w.AddBody(b4)

	b5 := physics.NewBody(physics.Circle{Radius: 40})
	b5.Position = physics.Vector{X: 110, Y: 100}
	// b5.Velocity = physics.Vector{X: 0, Y: -90}
	// b5.Static = true
	w.AddBody(b5)

	// b6 := physics.NewBody(physics.Rectangle{Size: physics.Vector{X: 40, Y: 100}})
	// b6.Position = physics.Vector{X: 300, Y: 140}
	// w.AddBody(b6)

	b7 := physics.NewBody(physics.Rectangle{Size: physics.Vector{X: 1500, Y: 100}})
	b7.Position = physics.Vector{X: 600, Y: 600}
	b7.Static = true
	w.AddBody(b7)

	// b8 := physics.NewBody(physics.Rectangle{Size: physics.Vector{X: 100, Y: 100}})
	// b8.Position = physics.Vector{X: 700, Y: 450}
	// b8.Velocity = physics.Vector{X: -60, Y: -80}
	// w.AddBody(b8)

	playground := playground.NewPhysicsPlayground(w)
	playground.Run(5070)
}
