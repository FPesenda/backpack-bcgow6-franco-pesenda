/*
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).
*/

package main

type product struct {
	Name     string
	Price    float64
	Quantity int
}

type services struct {
	Name          string
	Price         float64
	MinutesWorked int
}

type maintenance struct {
	Name  string
	Price float64
}

func PriceOfAllProducts(productsList []product) (price float64) {
	price = 0.0
	for _, v := range productsList {
		price += v.Price * float64(v.Quantity)
	}
	return
}

func PriceOfAllServices(services []services) (price float64) {
	price = 0.0
	for _, v := range services {
		if v.MinutesWorked < 30 {
			price += v.Price * 30
		} else {
			price += v.Price * float64(v.MinutesWorked)
		}
	}
	return
}

func PriceOfAllMaintenance(maintenances []maintenance) (price float64) {
	price = 0.0
	for _, v := range maintenances {
		price += v.Price
	}
	return
}

func main() {

}
