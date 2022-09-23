package main

import "fmt"
import "test"

func main(){
	ej1()
	ej2()
}

func ej1(){
	myName := "Tomas Cambiasso"
	myAddress := "Calle Falsa 123"

	fmt.Println(myName)
	fmt.Println(myAddress)
	return
}

func ej2(){
	var temp float32 = 10
	var humidity int = 20
	var pressure float32 = 30

	fmt.Printf("La temp es %f, la humedad es %d%% y lo otro es %f\n",temp, humidity, pressure)

	return 
}

func ej3(){
	// var 1nombre string INCORRECTO
	//var nombre string //ARREGLADO
	//var apellido string // CORRECTO
	//var int edad INCORRECTO, mal orden
	// var edad int // ARREGLADO
	// 1apellido := 6 INCORRECTO, no se ni que seria la forma correcta porque ni idea que quizo decir con apellido int
	// var licencia_de_conducir = true // INCORRECTO
	// var licenciaDeConducir = true ARREGLADO
	// var estatura de la persona int INCORRECTO
	//var estaturaDeLaPersona float32 // ARREGLADO
	//cantidadDeHijos := 2 // CORRECTO
}

func ej4(){
	/*
	var apellido string = "Gomez"
	var edad int = "35" / Sacar " "
	boolean := "false"; // Sacar " " y sacar ; y boolean al menos que ese sea el nombre de la variable, en tal caso, cambiar el nombre.
	var sueldo string = 45857.90 // Cambiar var a float
	var nombre string = "Julián"
	*/
  
}