package badpkg
import "fmt"

func Hello(name string) string {
    message := fmt.Sprintf("Hi, %v. Welcome! the version of badpkg is v1.0.2", name)
    return message
}


