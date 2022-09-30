package main

import "fmt"
import "os"
import "errors"

type customError struct{

}

func (e *customError) Error() string{
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func ej2(salary int) error{
	if salary < -150 {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return nil
}

func ej3(salary int) error{
	if salary < 150 {
		return fmt.Errorf("error: el salario minimo imponible es de 150k y el salario ingresado es de %d",salary)
	}
	return nil
}

func calculateTotalSalary( hours int, rate float64) (salary float64, err error){
	return
}

func main() {
	var salary int
	salary = 100
	var err customError
	if salary < -5 {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
	
	err2 := ej2(salary)
	if err2 != nil{
		fmt.Println(err2.Error())
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")

	err3 := ej3(salary)

	if err3 != nil{
		err := fmt.Errorf("%w",err3)
		fmt.Println(err)
	}

	//calculateTotalSalary()
}