# 删除有序数组中的重复项

题目链接: [https://leetcode.cn/problems/remove-duplicates-from-sorted-array/](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/)

解题思路一(用时间换空间):

1. 数组是有序的，那么重复出现的元素都有一个特性，除第一次出现外，后续多次出现的元素都是与前一个元素相同的，可用此方式判定数据重复
2. 仅需要将每个第一次出现的元素按先后顺序从数据头部重新覆盖写入

```go
func removeDuplicates(nums []int) int {
	n := 0
	for _, num := range nums {
		if n == 0 {
			n++
			continue
		} else {
			if nums[n-1] == num {
				continue
			} else {
				nums[n] = num
				n++
			}
		}
	}
	return n
}
```

解题思路二(用空间换时间):

1. 通过map存储出现过的元素，若元素未出现过，则记录并按先后顺序从数组头部覆盖写入
2. 仅需要将每个第一次出现的元素按先后顺序从数据头部重新覆盖写入

```go
func removeDuplicates(nums []int) int {
    numsMap:=map[int]int{}
    n:=0
    for _,num:=range nums{
        if _,has:=numsMap[num];!has{
            nums[n]=num
            n++
            numsMap[num]=1
        }   
    }
    return n
}
```
