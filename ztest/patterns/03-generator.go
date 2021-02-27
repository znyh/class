package patterns

func generator(count int) chan int{
	ch := make(chan int)

	go func(ch chan int) {
		for i := 0; i < count; i++ {
			ch <- i
		}
		close(ch)
	}(ch)
	return ch
}
