package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	length1 := m - 1
	length2 := n - 1
	length := m + n - 1

	for length1 >= 0 && length2 >= 0 {
		if nums1[length1] > nums2[length2] {
			nums1[length] = nums1[length1]
			length1--
		} else {
			nums1[length] = nums2[length2]
			length2--
		}

		length--
	}


	for length2 >= 0 {
		nums1[length] = nums2[length2]
		length--
		length2--
	}

}
