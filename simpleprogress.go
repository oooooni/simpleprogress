package simpleprogress

import(
	"fmt"
	"time"
)


func Simpleprogress(message string) chan<- string {
	ch := make(chan string)
	length := len(message) + 3
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("\rdone")
				return
			default:
				fmt.Print("\r")
				for i:=0; i<length; i++ {
					fmt.Print(" ")	
				}
				fmt.Printf("\r%s", message)
				for i := 0; i < 3; i++ {
					fmt.Print(".")
					time.Sleep(1 * time.Second)
				}
			}	
		}
	}()
	return ch	
}	
