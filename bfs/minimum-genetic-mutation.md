# 最小基因变化

题目链接: [https://leetcode.cn/problems/minimum-genetic-mutation](https://leetcode.cn/problems/minimum-genetic-mutation)

## 解题思路：

1. 判断`startGene`与`endGene`是否相等，相等直接返回0
2. 将`bank`转成`map`，判断`endGene`是否在`bank`中，不在直接返回-1
3. 从`startGene`开始处理，将每个字符用`ACGT`中与之不同的字符进行替换，并判断生成的新的基因是否在`bank`中，如果在，则判断是否是`endGene`，若是，则返回`step+1`，否则重新组建`step`存入队列，同时将该基因从`bank`中删除(若后续基因变换出现重复，那变换次数必然比这次多，所以可以直接从`bank`中移除不用判断是否重复)
4. 重复步骤3，直至找到`endGene`或者`queue`清空，若`queue`清空，则返回-1

```go
type step struct {
	value string
	step  int
}

func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	bankMap := make(map[string]bool, len(bank))
	for _, item := range bank {
		bankMap[item] = true
	}
	if !bankMap[endGene] {
		return -1
	}
	queue := []step{{startGene, 0}}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		for i, char := range item.value {
			for _, sig := range "ACGT" {
				if sig != char {
					next := item.value[:i] + string(sig) + item.value[i+1:]
					if bankMap[next] {
						if next == endGene {
							return item.step + 1
						}
						delete(bankMap, next)
						queue = append(queue, step{next, item.step + 1})
					}
				}
			}
		}
	}
	return -1
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(c*n*m)$$,$$n$$为基因序列的长度，$$m$$为`bank`的长度，$$c$$为基因的字符个数`ACGT`
- **空间复杂度：** 空间复杂度为$$O(c*n*m)$$,$$n$$为基因序列的长度，$$m$$为`bank`的长度，$$c$$为基因的字符个数`ACGT`
