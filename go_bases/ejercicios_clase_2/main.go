package main

import "fmt"

type Client struct {
	age       int
	employed  bool
	seniority int
	salary    float32
}

func main() {
	var myOtherArray = [...]int{1, -2, 3, 4, 34, 123}
	fmt.Println("myOtherArray", myOtherArray)
    fmt.Println("//------- Slice ----") //Tiene un array como base. El tamanio puede cambiar dinamicamente
    slice := myOtherArray[3:3]
    fmt.Println("slice", slice)
    fmt.Println("slice len:", len(slice))
    fmt.Println("cap len:", cap(slice))
	word := "word"
	ej1(word)
	client := Client{
		10, true, 2, 100000,
	}
	ej2(client)
	ej3(12)
	ej4("Benjamin")

}

func ej1(word string) {
	fmt.Println("La cantidad de letras es", (len(word)))
	for _, char := range word {
		fmt.Println(string(char))
	}
	return
}

func ej2(client Client) {
	if client.age < 22 || !client.employed || client.seniority < 1 {
		fmt.Println("No se puede dar prestamo")
	} else if client.salary > 100000 {
		fmt.Println("Prestamo sin interes")
	} else {
		fmt.Println("Prestamo con interes")
	}
	return
}

func ej3(month int) {
	months := [12]string{"Enero", "Feb", "Mar", "Abr", "May", "Jun", "jul", "Agosto", "Sept", "Oct", "Nov", "Dic"}

	fmt.Println(months[month-1])
	return
}

func ej4(name string){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println(name, "tiene" ,employees[name])
	overAge := 0
	for _, age := range employees{
		if age > 21{
			overAge++
		}
	}
	fmt.Println("Los empleados mayor a 21 son",overAge)
	employees["Federico"] = 25
	delete(employees,"Pedro")
	fmt.Println(employees)
	return
}