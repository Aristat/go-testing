package file_examples

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

var (
	fileScanner = &cobra.Command{
		Use:           "scanner",
		Short:         "Scanner example",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run: func(_ *cobra.Command, _ []string) {
			file, err := os.Open("resources/scanner.txt")

			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			go func() {
				ch := make(chan os.Signal, 1)
				signal.Notify(ch, os.Interrupt)
				<-ch
				fmt.Println("os.Signal Interrupt\n")
				os.Exit(1)
			}()

			start := int64(0)
			_, err = file.Seek(start, 0)

			scanner := bufio.NewScanner(file)
			pos := start
			scanLines := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
				advance, token, err = bufio.ScanLines(data, atEOF)
				pos += int64(advance)

				fmt.Println("current pos", pos)

				return
			}
			scanner.Split(scanLines)

			for scanner.Scan() {
				//time.Sleep(5 * time.Second)
				fmt.Println("line", scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		},
	}
	Cmd = &cobra.Command{
		Use:           "file",
		Short:         "File examples",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
)

func init() {
	Cmd.AddCommand(fileScanner)
}
