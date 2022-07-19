package tobucket

import (
	"fmt"
	"time"
)

func storerecording() {
	now := time.Now()
	fmt.Println("Hello world @ &v", now)
}
