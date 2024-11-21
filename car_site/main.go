package main

import (
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"strings"
)

var cars = map[string]string{
	"id1": "Renault Logan",
	"id2": "Renault Duster",
	"id3": "BMW X6",
	"id4": "BMW M5",
	"id5": "VW Passat",
	"id6": "VW Jetta",
	"id7": "Audi A4",
	"id8": "Audi Q7",
}

// carsListFunc — вспомогательная функция для вывода всех машин.
func carsListFunc() []string {
	var list []string
	for _, c := range cars {
		list = append(list, c)
	}
	return list
}

// carFunc — вспомогательная функция для вывода определённой машины.
func carFunc(id string) string {
	if c, ok := cars[id]; ok {
		return c
	}
	return "unknown identifier " + id
}

func carsHandle(rw http.ResponseWriter, r *http.Request) {
	carsList := carsListFunc()
	io.WriteString(rw, strings.Join(carsList, ", "))
}

func carHandle(rw http.ResponseWriter, r *http.Request) {
	carID := chi.URLParam(r, "id")
	if carID == "" {
		http.Error(rw, "carID param is missed", http.StatusBadRequest)
		return
	}
	rw.Write([]byte(carFunc(carID)))
}

func brandHandle(rw http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")
	rw.Write([]byte(strings.Join(getCarsByBrand(brand), ", ")))

}

func modelHandle(rw http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")
	model := chi.URLParam(r, "model")
	rw.Write([]byte(strings.Join(getCarByBrandAndModel(brand, model), ", ")))
}

func getCarsByBrand(brand string) []string {
	var list []string
	for _, c := range cars {
		if strings.Contains(c, brand) {
			list = append(list, c)
		}
	}
	return list
}

func getCarByBrandAndModel(brand string, model string) []string {
	var list []string
	for _, c := range cars {
		if strings.Contains(c, brand) {
			if strings.Contains(c, model) {
				list = append(list, c)
			}
		}
	}
	return list
}

func main() {

	r := chi.NewRouter()

	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("chi"))
	})

	r.Get("/cars", carsHandle)

	r.Route("/cars", func(r chi.Router) {
		r.Get("/", carsHandle) // GET /cars
		// Route можно вкладывать один в другой
		r.Route("/{brand}", func(r chi.Router) {
			r.Get("/", brandHandle)        // GET /cars/renault
			r.Get("/{model}", modelHandle) // GET /cars/renault/duster
		})
	})
	//r.Get("/car/{id}", carHandle)

	log.Fatal(http.ListenAndServe(":8080", r))

	//fmt.Printf("%+v", getCarsByBrand("Renault"))
	//fmt.Printf("%+v", getCarByBrandAndModel("Renault", "Logan"))

}
