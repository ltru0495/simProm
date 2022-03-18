package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const DELAY = 2

func isIn(s string, arr []string) bool {
	for _, el := range arr {
		if el == s {
			return true
		}
	}
	return false
}

type Data struct {
	ID    string
	param string
	value float64
}

// Array de metricas a medir
// se usa una variable global para evitar duplicados al registrar
var metrics []*Metric

func sendMetric(data []Data) {
	for _, d := range data {

		for _, m := range metrics {
			if m.name == d.param {
				m.p.WithLabelValues(d.ID).Set(d.value)
			}
		}
	}
}

// Recibe los datos simulados mediante el canal c
func receiveData(c chan []Data) {
	// array que almacena los modulos ya registrados
	var mods []string

	// iteracion cada vez que llega un dato al canal c
	for dataArr := range c {
		// en caso no este registrado el modulo, se registra su Id
		// y se verifica si hay una nueva metrica que medir
		if !isIn(dataArr[0].ID, mods) {
			mods = append(mods, dataArr[0].ID)
			NewMetric(dataArr)  // Crear una metrica -> dev.go
			sendMetric(dataArr) // Envia metricas
		} else {
			sendMetric(dataArr) // Envia metricas
		}
	}
}

func main() {
	c := make(chan []Data)
	go receiveData(c)

	//Creacion de 2 modulos separados por 5 segs
	m1 := &Mod{
		Id:     "MOD1",
		Params: []string{"temperature", "light"},
	}
	go m1.run(c)
	time.Sleep(5 * time.Second)

	m2 := &Mod{
		Id:     "MOD2",
		Params: []string{"temperature", "humidity", "co2", "o2"},
	}
	go m2.run(c)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
