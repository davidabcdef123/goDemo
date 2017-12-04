package main

func main() {

}

func getAverage(arr []int, size int) int {
	var i,sum int
	var avg int

	for i = 0; i < size; i++ {
		sum += arr[i];
	}
	avg = sum / size
	return avg
}


