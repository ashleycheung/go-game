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

	// b1 := physics.NewBody(physics.Circle{Radius: 10})
	// b1.Position = physics.Vector{X: 100, Y: 100}
	// b1.Velocity = physics.Vector{X: 100}
	// w.AddBody(b1)

	// b2 := physics.NewBody(physics.Circle{Radius: 10})
	// b2.Position = physics.Vector{X: 200, Y: 100}
	// b2.Sensor = true
	// b2.GetEvent().AddListener(string(physics.BodyCollideEvent), func(e event.Event) error {
	// 	w.RemoveBody(b2)
	// 	return nil
	// })
	// w.AddBody(b2)

	playground := playground.NewPhysicsPlayground(w)
	playground.Run(5080)
}
