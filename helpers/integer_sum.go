/*
Write a program that uses goroutines to calculate the sum of an array of integers.
The program should divide the array into four equal parts and
launch a goroutine to calculate the sum of each part.
The program should wait for all the goroutines to complete and then calculate the final sum.
*/

package helpers

import (
	"fmt"
	"sync"
)

func ArraySum(wg *sync.WaitGroup, mu *sync.Mutex) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	chunckSize := len(arr) / 4

	fmt.Println("chuck size of array", chunckSize)

	sum := 0

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()

			end := start + chunckSize
			fmt.Println("end :", end)

			if end == 3 {
				end = len(arr)
			}

			partSum := 0

			for j := start; j < end; j++ {
				partSum += arr[j]
			}

			mu.Lock()

			sum += partSum

			mu.Unlock()

		}(i * chunckSize)
		fmt.Printf("Sum of array %d\n", sum)
		wg.Wait()

	}

}

func ArraySumFunc1(wg *sync.WaitGroup, mu *sync.Mutex) {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chunckArray := DivideArray(arr, 4)
	sum := 0

	for i := 0; i < len(chunckArray); i++ {
		wg.Add(1)
		go func(subArray []int) {
			defer wg.Done()
			for _, value := range subArray {
				mu.Lock()
				sum += value
				mu.Unlock()
			}
		}(chunckArray[i])
		wg.Wait()

	}
	fmt.Println("Sum of all sub array is : ", sum)

	fmt.Println("Array length: ", len(chunckArray))

}

func DivideArray(nums []int, parts int) [][]int {
	// Calculate the size of each part
	size := (len(nums) + parts - 1) / parts

	fmt.Println("size: ", size)

	// Create a slice of slices to hold the subarrays
	subarrays := make([][]int, parts)

	// Loop through the array and extract the subarrays
	for i := 0; i < len(nums); i += size {
		end := i + size
		if end > len(nums) {
			end = len(nums)
		}
		subarrays[i/size] = nums[i:end]
	}

	return subarrays
}
