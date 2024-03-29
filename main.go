package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/experimental/devices/mcp9808"
	"periph.io/x/periph/host"
)

var dev bool
var port int

func init() {
	flag.BoolVar(&dev, "dev", false, "development mode")
	flag.IntVar(&port, "port", 80, "port to serve on")
	flag.Parse()
}

func readTemp() int {
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Open default I²C bus.
	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatalf("failed to open I²C: %v", err)
	}
	defer bus.Close()

	// Create a new temperature sensor.
	sensor, err := mcp9808.New(bus, &mcp9808.DefaultOpts)
	if err != nil {
		log.Fatalln(err)
	}

	// Read values from sensor.
	measurement, err := sensor.SenseTemp()

	if err != nil {
		log.Fatalln(err)
	}

	return int(measurement.Fahrenheit())
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))

	temp := 65
	if !dev {
		temp = readTemp()
	}

	tmpl.Execute(w, temp)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)

	addr := fmt.Sprintf(":%v", port)
	log.Printf("Listening on http://localhost%v ...\n", addr)

	http.ListenAndServe(addr, nil)
}
