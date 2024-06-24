package internal
import (
	"fmt"
	"os"
	"strings"
)
func Env() (err error) {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		values := strings.Split(pair[1], ";")
		for _, value := range values {
			fmt.Println(pair[0], "=", value)
		}
	}
	return nil
}