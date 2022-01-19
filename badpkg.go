package badpkg
import "fmt"
import "rsc.io/quote"

func Hello(name string) string {
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}

func Test() string {
    msg := fmt.Sprintf("Hi")
    return quote.Hello()
}
