package main

import (
	"math/rand"

	"github.com/ashleycheung/go-game/physics"
	"github.com/ashleycheung/go-game/physics/playground"
)

func main() {
	w := physics.NewWorld()
	w.Config.AirResistance = 0

	// Add random amount of objects
	for i := 0; i < 2000; i++ {
		b := physics.NewBody(physics.Circle{Radius: 5})
		b.Position = physics.Vector{X: rand.Float64() * 1000, Y: rand.Float64() * 1000}
		b.Velocity = physics.Vector{X: rand.Float64()*200 - 100, Y: rand.Float64()*200 - 100}
		w.AddBody(b)
	}

	playground := playground.NewPhysicsPlayground(w)
	playground.Run(5070)
}
