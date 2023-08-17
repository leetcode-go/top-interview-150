# 简化路径

题目链接: [https://leetcode.cn/problems/simplify-path](https://leetcode.cn/problems/simplify-path)

## 解题思路：

1. 按照/拆分路径字符串，分割成字符串数组
2. 遍历字符串数组，对于不同元素进行不同的操作
3. 如果字符串是`..`则进行出队操作，如果是`.`或空格则不动，其他元素均进行入队操作
4. 入队出队的顺序按照先进后出的逻辑进行
5. 如果队列长度为0，则不执行出队操作


```go
func simplifyPath(path string) string {
	pathList := strings.Split(path, "/")
	queue := make([]string, 0)
	for _,item:=range pathList{
        // ..表示进入上一层目录，移除掉队尾元素
		if item==".."{
			if len(queue)>0{
				queue=queue[:len(queue)-1]
			}
		}else if item!=""&& item!="."{
            // 既不是空格也不是.就入队
			queue=append(queue,item)
		}
	}
	return fmt.Sprintf("/%s",strings.Join(queue,"/"))
}
```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍字符串，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是字符串 `s` 的长度
- **空间复杂度：** 空间复杂度为 $$O(n)$$，函数使用了一个长度为 $$n$$ 的切片来存储栈。在最坏的情况下，栈的深度可能达到 $$n$$，所以空间复杂度为 $$O(n)$$
