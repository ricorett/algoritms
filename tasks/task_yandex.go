package task_yandex

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var amount, quality int
	_, err := fmt.Scan(&amount, &quality)
	if err != nil {
		return
	}
	if amount <= 0 {
		fmt.Println(0)
		return
	}

	scanner := bufio.NewReader(os.Stdin)
	arr, _ := scanner.ReadString('\n')
	parts := strings.Fields(arr)
	count := 0
	for i := 0; i < amount && i < len(parts); i++ {
		num, err := strconv.Atoi(parts[i])
		if err != nil {
			continue
		}
		if num > quality {
			count++
		}
	}
	fmt.Println(count)

	return
}

//nums := make([]int, 0, amount)
//for i, v := range parts {
//	if i >= amount {
//		break
//	}
//	num, _ := strconv.Atoi(v)
//	nums = append(nums, num)
//}
//
//if len(nums) == 0 {
//	fmt.Println(0)
//	return
//}
//
//left := 0
//right := len(nums) - 1
//
//if nums[left] <= quality {
//	fmt.Println(0)
//	return
//}
//
//for left <= right {
//	mid := left + (right-left)/2
//	if nums[mid] > quality {
//		left = mid + 1
//	} else {
//		right = mid - 1
//	}
//}
//fmt.Println(left)
