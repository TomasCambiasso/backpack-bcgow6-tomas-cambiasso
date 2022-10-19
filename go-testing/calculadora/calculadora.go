package calculadora

import "fmt"

func Restar(num1, num2 int) int {
	return num1 - num2
}

func QuickSort(list []int, low int, high int) {
	if low < high {
		pi := partition(list, low, high)
		QuickSort(list, low, pi-1)
		QuickSort(list, pi+1, high)
	}
}

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}

func partition(list []int, low int, high int) int {
	pivot := list[high]
	i := (low - 1)

	for j := low; j < high; j++ {
		if list[j] <= pivot {
			i++
			swap(&list[i], &list[j])
		}
	}
	swap(&list[i+1], &list[high])

	return i + 1
}

func Divide(num int, dem int) (int, error) {
	if dem == 0 {
		return 0, fmt.Errorf("El denominador no puede ser 0")
	}
	return num / dem, nil
}
