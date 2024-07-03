package greeting

import (
	"fmt"

	"github.com/google/uuid"
)

func Hello() {
	fmt.Println("Hello")
}

func GenerateUUID() {
	fmt.Println(uuid.NewString())
}
