# 最小栈

题目链接: [https://leetcode.cn/problems/min-stack](https://leetcode.cn/problems/min-stack)

## 解题思路：

1. 数据入栈前，对数据进行判断是否小于当前栈中最小的元素，若是，则在入栈后覆盖栈中记录的最小值
2. 数据出栈，减小栈的高度，并判断出栈元素是否为最小元素，若是则重新查找最小元素


```go
type MinStack struct {
    data []int
    min int
    length int
}


func Constructor() MinStack {
    return MinStack{
        data:make([]int,0),
        min:2147483647,
    }
}


func (this *MinStack) Push(val int)  {
	this.data = append(this.data, val)
    if val<this.min{
        this.min=val
    }
}


func (this *MinStack) Pop()  {
    item:=this.data[len(this.data)-1]
    this.data=this.data[:len(this.data)-1]
    if item==this.min{
        // 2^31-1
        min:=2147483647
        for i:=0;i<len(this.data);i++{
            if this.data[i]<min{
                min=this.data[i]
            }
        }
        this.min=min
    }
}


func (this *MinStack) Top() int {
    return this.data[len(this.data)-1]
}


func (this *MinStack) GetMin() int {
    return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍数据，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是栈的高度
- **空间复杂度：** 除了栈本身需要的空间外，只使用了常数个变量，因此空间复杂度为 $$O(1)$$
