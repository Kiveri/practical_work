package main

import (
	"fmt"
	"math"
)

/*

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    merged := make([]int, 0, len(nums1)+len(nums2))
    i, j := 0, 0

    for i < len(nums1) && j < len(nums2) {
        if nums1[i] < nums2[j] {
            merged = append(merged, nums1[i])
            i++
        } else {
            merged = append(merged, nums2[j])
            j++
        }
    }

    for i < len(nums1) {
        merged = append(merged, nums1[i])
        i++
    }
    for j < len(nums2) {
        merged = append(merged, nums2[j])
        j++
    }

    mid := len(merged) / 2
    if len(merged)%2 == 0 {
        return (float64(merged[mid-1]) + float64(merged[mid])) / 2.0
    } else {
        return float64(merged[mid])
    }
}

*/

// findMedianSortedArrays находит медиану двух отсортированных массивов
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Убедимся, что nums1 — это меньший массив
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m
	medianPos := (m + n + 1) / 2

	for left <= right {
		i := (left + right) / 2 // Индекс для первого массива
		j := medianPos - i      // Индекс для второго массива

		if i < m && nums1[i] < nums2[j-1] {
			// i слишком мал, нужно увеличить
			left = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			// i слишком велик, нужно уменьшить
			right = i - 1
		} else {
			// Нашли правильное разделение
			var maxLeft float64
			if i == 0 {
				maxLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxLeft = float64(nums1[i-1])
			} else {
				maxLeft = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}

			if (m+n)%2 == 1 {
				return maxLeft
			}

			var minRight float64
			if i == m {
				minRight = float64(nums2[j])
			} else if j == n {
				minRight = float64(nums1[i])
			} else {
				minRight = math.Min(float64(nums1[i]), float64(nums2[j]))
			}

			return (maxLeft + minRight) / 2.0
		}
	}

	return 0
}

func main() {
	// Пример использования
	nums1 := []int{1, 3}
	nums2 := []int{2, 5}
	median := findMedianSortedArrays(nums1, nums2)
	fmt.Println("Median:", median) // Ожидаемый результат: 2.0
}
