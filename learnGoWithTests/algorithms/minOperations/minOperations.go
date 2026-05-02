package minoperations

func minOperations(nums []int, k int) int {
    var somatorio int

    for _ , num := range nums{
        somatorio += num
    }
    return somatorio % k
}
