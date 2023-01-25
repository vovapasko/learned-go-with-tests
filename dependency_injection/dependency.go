package dependency_injection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, inputString string) error {
	_, err := fmt.Fprintf(writer, "Hello, %s", inputString)
	if err != nil {
		return err
	}
	return nil
}
