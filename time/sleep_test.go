package time_test

import (
	"context"
	"fmt"
	"time"

	gotime "github.com/foomo/go/time"
)

func ExampleSleep() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Wait for 1 second (completes successfully)
	err := gotime.Sleep(ctx, 1*time.Second)
	if err != nil {
		fmt.Println("Sleep failed:", err)
		return
	}

	fmt.Println("Sleep completed successfully")

	// Sleep for 3 seconds with a 2-second timeout (context cancels first)
	ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel2()

	err = gotime.Sleep(ctx2, 3*time.Second)
	if err != nil {
		fmt.Println("Sleep cancelled:", err)
	}

	// Output:
	// Sleep completed successfully
	// Sleep cancelled: context deadline exceeded
}
