package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var SubArrayMaxSumCmd = &cobra.Command{
	Use:   "subArrayMaxSum",
	Short: "Calculate maximum sub array sum",
	Run:   SubArrayMaxSum,
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func SubArrayMaxSum(cmd *cobra.Command, args []string) {
	fmt.Println("Enter numbers for array separated by newline.")
	fmt.Println("Enter empty line to calculate")

	scanner := bufio.NewScanner(os.Stdin)
	var array []int
	for scanner.Scan() {
		num := scanner.Text()
		if len(num) == 0 {
			break
		}

		arrayNum, err := strconv.Atoi(num)
		fatalOnError(err)

		array = append(array, arrayNum)
	}

	if len(array) == 0 {
		log.Fatalln("Please put valid array")
	}

	var tSum, maxSum int
	for _, val := range array {
		tSum = max(val, tSum+val)
		maxSum = max(maxSum, tSum)
	}

	log.Printf("Max subarray sum: %v", maxSum)
}
