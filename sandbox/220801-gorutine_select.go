package main
import "fmt"

func evenSum(from, to int, ch chan int) {
    result := 0
    for i:=from; i<=to; i++ {
        if i%2 == 0 {
            result += i
        }    
    }
    ch <- result
}
func squareSum(from, to int, ch chan int) {
    result := 0
    for i:=from; i<=to; i++ {
        if i%2 == 0 {
            result += i*i
        }    
    }
    ch <- result
}

func main() {
    evenCh := make(chan int)
    sqCh := make(chan int)

    go evenSum(0, 100, evenCh)
    go squareSum(0, 100, sqCh)

/* Waits for whichever channel receives a value first using select:
 * Because both channels are unbuffered, each goroutine will block at
 * ch <- result
 * until main’s select chooses that case.
 * The select will pick the first case that becomes ready.
 */
    select {
        case x := <- evenCh:
            fmt.Println(x,"evenSum wins")
        case y := <- sqCh:
            fmt.Println(y,"squareSum wins")
    }
}