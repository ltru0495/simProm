package main

import (
	"math/rand"
	"time"
)

type Mod struct {
	Id     string
	Params []string
}

// Simulador de lectura de datos
// se usa el canal de datos c para enviar datos aleatorios
func (m Mod) run(c chan []Data) {
	for {
		var simData []Data
		// se generan datos aleatorios para cada parametro en un modulo
		for _, p := range m.Params {
			data := Data{m.Id, p, rand.Float64()*50 + 50}
			simData = append(simData, data)

		}
		c <- simData
		// Delay entre envio de datos
		time.Sleep(DELAY * time.Second)
	}
}
