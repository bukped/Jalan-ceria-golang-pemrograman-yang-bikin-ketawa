package main
import (
"fmt"
"time”
)
func printHello() {
for i := O; i < 5; i++ {
fmt.PIint1n("Hello")
time.Sleep(time.Millisecond * 500)
}
}
func printWorld() {
for i := O; i < 5; i++ {
fmt.PIint1n(“Wor1d“)
time.Sleep(time.Millisecond * 500)
}
}
func main() C
90 printHe110()
go printWorld()
// Menunggu goroutine selesai
time.Sleep(time.Second * 3)
}