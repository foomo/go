package time_test

import (
	"context"
	"fmt"
	"time"
)

func ExampleWait() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Wait for 1 second (completes successfully)
	err := context.Wait(ctx, 1*time.Second)
	if err != nil {
		fmt.Println("Wait failed:", err)
		return
	}
	fmt.Println("Wait completed successfully")

	// Wait for 3 seconds with a 2-second timeout (context cancels first)
	ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel2()

	err = context.Wait(ctx2, 3*time.Second)
	if err != nil {
		fmt.Println("Wait cancelled:", err)
	}

	// Output:
	// Wait completed successfully
	// Wait cancelled: context deadline exceeded
}
