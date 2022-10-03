package main

import "os"
import "fmt"


type Customer struct {
	id int
	nameAndSurname string
	nid float64
	phoneNumber int
	address string
}

func ej1(path string) (file *os.File, err interface{}) {

	defer func () interface{}{
		err = recover()
		fmt.Println("ejecucion ej1 finalizada")
		return err
	}()

	file, err = os.Open(path)
	if err != nil {
		panic("error: el archivo indicado no fue encontrado o esta dañando")
	}
	return file, err
}

func getLatestCustomer(ids []Customer) *Customer{
	if len(ids) == 0 {
		return nil
	}else {
		return &ids[len(ids)-1]
	}
}

func (c *Customer)validateFields(){

}

func main() {
	path := "./customers.txt"
	_, err := ej1(path)
	fmt.Println(err)
	var customers []Customer
	customer := getLatestCustomer(customers)
	if customer == nil {
		panic("panic: no hay clientes cargados")
	}




}
