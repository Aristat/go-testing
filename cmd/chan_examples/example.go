package chan_examples

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var (
	chanBase = &cobra.Command{
		Use:           "chan_base",
		Short:         "Default chan example",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run: func(_ *cobra.Command, _ []string) {
			ch := make(chan int)
			done := make(chan struct{}) // create DONE channel

			var wg sync.WaitGroup

			numbers := []int{1, 2, 3}

			wg.Add(len(numbers))

			for _, n := range numbers {
				go func(n int) {
					defer wg.Done()
					ch <- n
				}(n)
			}

			go func() {
				defer close(done) // close done channel to tell that all jobs is done
				for c := range ch {
					fmt.Printf("routine start %v\n", c)
					time.Sleep(1 * time.Second) // for better understanding
					fmt.Printf("routine done %v\n", c)
				}
			}()

			wg.Wait()
			close(ch) // after all routines push payload then we close channel to release RANGE
			<-done    // wait when all jobs is done

			fmt.Printf("Done!\n")
		},
	}
	Cmd = &cobra.Command{
		Use:           "chan",
		Short:         "Chan examples",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
)

func init() {
	Cmd.AddCommand(chanBase)
}
