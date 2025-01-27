// lpp
package lpp

import (

	//	"fmt"

	"math"

	"sort"
)

func COORD_SORT(array [][2]int, status int) [][2]int {
	cache_hash := make(map[int][][2]int)
	result := [][2]int{}
	cache_data := []int{}
	for _, data := range array {
		num := data[0]
		_, ok := cache_hash[num]
		if !ok {
			cache_hash[num] = [][2]int{}

		}
		cache_hash[num] = append(cache_hash[num], data)

	}
	for key, _ := range cache_hash {
		cache_data = append(cache_data, key)
	}
	if status == 0 {

		sort.Ints(cache_data)
	} else {

		sort.Sort(sort.Reverse(sort.IntSlice(cache_data)))
	}
	for _, element := range cache_data {
		for _, value := range cache_hash[element] {
			result = append(result, value)
		}
	}

	return result
}
func COORD_MERGE(array [][2]int, status int) int {
	array = COORD_SORT(array, status)

	length := 0
	result_list := [][2]int{}
	for i, data := range array {
		if i == 0 {
			result_list = append(result_list, data)
		} else {
			length := len(result_list) - 1
			if data[0] > data[1] {
				if data[0] >= result_list[length][1] {
					if data[1] < result_list[length][1] {
						result_list[length][1] = data[1]
					}
				} else {
					result_list = append(result_list, data)
				}
			} else {
				if data[0] <= result_list[length][1] {
					if data[1] >= result_list[length][1] {
						result_list[length][1] = data[1]
					}
				} else {
					result_list = append(result_list, data)
				}
			}
		}
	}

	for _, data := range result_list {
		length += int(math.Abs(float64(data[0]-data[1])) + 1)
	}
	return length
}
func COORD_CHAIN(array [][2]int, raw_array [][2]int, status int) ([][2]int, [][2]int, int) {
	//status 0是递增，status 1是递减

	i, j, k, max := 0, 0, 0, 0
	var result, result2 [][2]int
	length := len(array)

	//变长数组参数，C99新特性，用于记录当前各元素作为最大元素的最长递增序列长度
	liss := make([]int, length)

	//前驱元素数组，记录当前以该元素作为最大元素的递增序列中该元素的前驱节点，用于打印序列用
	pre := make([]int, length)

	for i = 0; i < length; i++ {

		liss[i] = 0
		pre[i] = i
	}

	max = 0
	k = 0
	for i = 1; i < length; i++ {
		//找到以array[i]为最末元素的最长递增子序列
		for j = 0; j < i; j++ {
			if status == 0 {
				if array[j][1] <= array[i][1] && array[i][0] < array[i][1] {
					cache := liss[j] + array[i][1] - array[i][0]

					if cache > liss[i] {
						liss[i] = cache
						if cache > max {
							max = cache
							k = i
						}
						if array[j][0] < array[j][1] {

							pre[i] = j
						}
					}
				}
			} else {
				if array[j][0] > array[i][0] && array[j][1] > array[i][1] && array[i][0] > array[i][1] {
					cache := liss[j] + array[i][0] - array[i][1]

					if cache > liss[i] {
						liss[i] = cache
						if cache > max {
							max = cache
							k = i
						}
						if array[j][0] > array[j][1] {
							pre[i] = j
						}
					}
				}

			}
			//如果要求非递减子序列只需将array[j] < array[i]改成<=，
			//如果要求递减子序列只需改为>

		}
	}
	//	fmt.Println(pre, k)
	//输出序列
	//	= max - 1

	result = [][2]int{array[k]}
	result2 = [][2]int{raw_array[k]}
	for {
		if pre[k] == k {
			break
		}

		k = pre[k]
		result = append([][2]int{array[k]}, result...)
		result2 = append([][2]int{raw_array[k]}, result2...)
	}
	//	fmt.Println(result)
	total_length := COORD_MERGE(result, status)
	return result, result2, total_length
}

func LCS(array []int, raw_array []int, status int) ([]int, []int) {
	//status 0是递增，status 1是递减

	i, j, k, max := 0, 0, 0, 0
	var result, result2 []int
	length := len(array)

	//变长数组参数，C99新特性，用于记录当前各元素作为最大元素的最长递增序列长度
	liss := make([]int, length)

	//前驱元素数组，记录当前以该元素作为最大元素的递增序列中该元素的前驱节点，用于打印序列用
	pre := make([]int, length)

	for i = 0; i < length; i++ {

		liss[i] = 1
		pre[i] = i
	}

	max = 1
	k = 0
	for i = 1; i < length; i++ {
		//找到以array[i]为最末元素的最长递增子序列
		for j = 0; j < i; j++ {
			//如果要求非递减子序列只需将array[j] < array[i]改成<=，
			//如果要求递减子序列只需改为>
			if status == 0 {

				if array[j] < array[i] && liss[j]+1 > liss[i] {

					liss[i] = liss[j] + 1
					pre[i] = j

					//得到当前最长递增子序列的长度，以及该子序列的最末元素的位置
					if max < liss[i] {

						max = liss[i]
						k = i
					}
				}
			} else {

				if array[j] > array[i] && liss[j]+1 > liss[i] {

					liss[i] = liss[j] + 1
					pre[i] = j

					//得到当前最长递增子序列的长度，以及该子序列的最末元素的位置
					if max < liss[i] {

						max = liss[i]
						k = i
					}

				}
			}
		}
	}

	//输出序列
	i = max - 1

	result = []int{array[k]}
	result2 = []int{raw_array[k]}
	for {
		if pre[k] == k {
			break
		}

		k = pre[k]
		result = append([]int{array[k]}, result...)
		result2 = append([]int{raw_array[k]}, result2...)
	}

	return result, result2
}

// min of two integers
func min(a int, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return
}

// max of two integers
func maxI(a int, b int) (res int) {
	if a < b {
		res = b
	} else {
		res = a
	}

	return
}

// max of two float64s
func max(a float64, b float64) (res float64) {
	if a < b {
		res = b
	} else {
		res = a
	}

	return
}

const GAP_COST = float64(0.5)

func getCost(r1 []rune, r1Index int, r2 []rune, r2Index int) float64 {
	if r1[r1Index] == r2[r2Index] {
		return 1.0
	} else {
		return -2.0
	}
}

// SmithWaterman computes the Smith-Waterman local sequence alignment for the
// two input strings. This was originally designed to find similar regions in
// strings representing DNA or protein sequences.
func SmithWaterman(s1 string, s2 string) float64 {
	var cost float64

	// index by code point, not byte
	r1 := []rune(s1)
	r2 := []rune(s2)

	r1Len := len(r1)
	r2Len := len(r2)

	if r1Len == 0 {
		return float64(r2Len)
	}

	if r2Len == 0 {
		return float64(r1Len)
	}

	d := make([][]float64, r1Len)
	for i := range d {
		d[i] = make([]float64, r2Len)
	}

	var maxSoFar float64
	for i := 0; i < r1Len; i++ {
		// substitution cost
		cost = getCost(r1, i, r2, 0)
		if i == 0 {
			d[0][0] = max(0.0, max(-GAP_COST, cost))
		} else {
			d[i][0] = max(0.0, max(d[i-1][0]-GAP_COST, cost))
		}

		// save if it is the biggest thus far
		if d[i][0] > maxSoFar {
			maxSoFar = d[i][0]
		}
	}

	for j := 0; j < r2Len; j++ {
		// substitution cost
		cost = getCost(r1, 0, r2, j)
		if j == 0 {
			d[0][0] = max(0, max(-GAP_COST, cost))
		} else {
			d[0][j] = max(0, max(d[0][j-1]-GAP_COST, cost))
		}

		// save if it is the biggest thus far
		if d[0][j] > maxSoFar {
			maxSoFar = d[0][j]
		}
	}

	for i := 1; i < r1Len; i++ {
		for j := 1; j < r2Len; j++ {
			cost = getCost(r1, i, r2, j)

			// find the lowest cost
			d[i][j] = max(
				max(0, d[i-1][j]-GAP_COST),
				max(d[i][j-1]-GAP_COST, d[i-1][j-1]+cost))

			// save if it is the biggest thus far
			if d[i][j] > maxSoFar {
				maxSoFar = d[i][j]
			}
		}
	}

	return maxSoFar
}
