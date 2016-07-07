//Notes on Code Here People!!!

//Great examples here
func gen(nums ...int) chan int {
	fmt.Printf("TYPE OF NUMS %T\n", nums) // just FYI

	out := make(chan int)
	go func() {
		for _, n := range nums {      //this ranges and gets the value, not the index from a slice of int (or map)
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {           //this ranges over a channel of ints and gets the values
			out <- n * n
		}
		close(out)
	}()
	return out
}

//////////

func merge(cs ...chan int) chan int {
	fmt.Printf("TYPE OF CS: %T\n", cs) // just FYI

	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)                     //why do we have this c here? has to do with closure + scope of channels
	}
