package main

/*
	The pipeline pattern is used to process data through a series of stages,
	where each stage runs in its own goroutine and passes results to
	the next stage via channels.
*/

/*
	in our example, we will take a array of nums, put it to a channel, and sqrt it
*/

func sliceToChannel(arr []int) <-chan int {

	out := make(chan int)
	go func() {
		for _, i := range arr {
			out <- i
		}

		close(out)
	}()

	return out
}

func sqrt(nums <-chan int) <-chan int {

	out := make(chan int)
	go func() {
		for n := range nums {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func main() {

	// input
	var arr = []int{1, 2, 3, 4}
	//stage 1
	nums := sliceToChannel(arr)
	// stage 2
	finalChannel := sqrt(nums)
	// stage 3
	for res := range finalChannel {
		println("sqrt of one of elem is ", res)
	}
}
