package badpkg
import "fmt"

func Hello(name string) string {
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}

func Test() string {
    msg := 'hhh i modify v1.0.1'
    return msg
}
