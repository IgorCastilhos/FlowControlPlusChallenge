package main

import "FlowControl/consoleApp"

func main() {
	//	Three part loop
	//for i := 10; i >= 0; i-- {
	//	fmt.Println("i is", i)
	//}f

	// While loop
	//rand.Seed(time.Now().UnixNano())
	//
	//j := 1000
	//
	//for j > 100 {
	//	j = rand.Intn(1000) + 1
	//	fmt.Println("j is ", j)
	//	if j > 100 {
	//		fmt.Println("j is", j, "so loop just keeps going")
	//	}
	//}
	//fmt.Println("Got", j, "and broke out of loop")

	//	Infinite loop
	//reader := bufio.NewReader(os.Stdin)
	//ch := make(chan string)
	//
	//go mylogger.ListenForLog(ch)
	//
	//fmt.Println("Enter something")
	//
	//for i := 0; i < 5; i++ {
	//	fmt.Print("-> ")
	//	input, _ := reader.ReadString('\n')
	//	ch <- input
	//	time.Sleep(time.Second)
	//}

	//for i := 1; i <= 10; i++ {
	//	fmt.Print("i is ", i)
	//	for j := 1; j <= 3; j++ {
	//		fmt.Print("   j: ", j)
	//	}
	//	fmt.Println()
	//}

	consoleApp.ConsoleApplication()
}
