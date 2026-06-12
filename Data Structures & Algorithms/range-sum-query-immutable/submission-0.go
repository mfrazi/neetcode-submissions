type NumArray struct {
    prefixSum []int
}

// 1 2 3 4
// 1 3 6 10

func Constructor(nums []int) NumArray {
    prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i:=1; i<len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	return NumArray{
		prefixSum: prefixSum,
	}
}


func (this *NumArray) SumRange(left int, right int) int {
    if left == 0 {
		return this.prefixSum[right]
	}
	return this.prefixSum[right] - this.prefixSum[left-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */