# O(1)时间插入、删除和获取随机元素


题目链接: [https://leetcode.cn/problems/insert-delete-getrandom-o1](https://leetcode.cn/problems/insert-delete-getrandom-o1)

## 解题思路：

1. 插入、删除、读取操作都需要在O(1)的时间内实现，所以直接用链表不合适，直接用数组也不合适
2. 可以采用切片+哈希表的方式进行实现
3. 通过切片能实现O(1)时间的插入以及随机读取，再通过哈希表记录每个元素的位置，在删除的时候直接选择元素覆盖缩减切片即可


```go
type RandomizedSet struct {
    data []int
    idxMap map[int]int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        data:[]int{},
        idxMap: map[int]int{},
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _,ok:=this.idxMap[val];ok{
        return false
    }
    // 记录元素插入的位置
    this.idxMap[val]=len(this.data)
    this.data = append(this.data,val)
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    id,ok:=this.idxMap[val]
    if !ok{
        return false
    }
    last:=len(this.data)-1
    // 覆盖元素
    this.data[id]=this.data[last]
    // 修正最后一个元素的位置记录
    this.idxMap[this.data[id]] = id
    // 缩减切片
    this.data=this.data[:last]
    // 删除哈希表中的位置记录
    delete(this.idxMap,val)
    return true
}


func (this *RandomizedSet) GetRandom() int {
    return this.data[rand.Intn(len(this.data))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
```

## 复杂度分析一

- **时间复杂度：** 因此时间复杂度为 $$O(1)$$
- **空间复杂度：** 空间复杂度为 $$O(n)$$，其中 $$n$$ 是数组 $$输入数据的长度$$ 的长度