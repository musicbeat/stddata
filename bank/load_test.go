package bank
// Keep reading: http://golang.org/doc/code.html#Testing
import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	fmt.Println("Test: bank.Load")
	bank := new(Bank)
	n, err := bank.Load()
	// check count or something.
	fmt.Printf("%d\n", n)
	fmt.Printf("%s\n", err)
}