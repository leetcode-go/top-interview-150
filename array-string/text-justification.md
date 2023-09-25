# 文本左右对齐

题目链接: [https://leetcode.cn/problems/text-justification](https://leetcode.cn/problems/text-justification)

# 解题思路

1. 字符串分行有以下四种情况
    - 空格可以均匀分布，则每两个字符串之间用相同长度的空格字符串拼接
    - 空格不可以均匀分布，则左侧的字符串之间使用的空格字符串长度比右侧长
    - 一行只有一个字符串的情况下，右侧空位用空格填满
    - 最后一行剩余的空位，用空格填满
2. 遍历字符串数组，判断长度是否构成一行，若构成，则开始按情况进行处理
3. 对于第一种情况与第二种情况，可以一起处理，先按照平均分布计算出单词之间空格个数，再计算出剩余空格数量，平均分配到左侧每个空格字符串末尾
    

```go
func fillSpace(n int)string{
    return strings.Repeat(" ", n)
}

func fullJustify(words []string, maxWidth int) []string {
    right,n :=0,len(words)
    res:=make([]string,0)
    for{
        // 每一行的开始都是上一行最后一个字符串的下一个字符串
        left:=right
        subLength:=0
        for right<n && subLength+len(words[right])+right-left<=maxWidth{
            subLength+=len(words[right])
            right++
        }
        // 最后一行，左对齐
        if right==n{
            s:=strings.Join(words[left:]," ")
            res=append(res,s+fillSpace(maxWidth-len(s)))
            return res
        }

        // 当前行包含的字符串个数
        wordsCount:=right-left
        // 当前行应包含的空格数
        spacesCount:=maxWidth-subLength
        if wordsCount==1{
            res=append(res,words[left]+fillSpace(spacesCount))
            continue
        }
        // 计算当前行两个字符串之间平均应有多少空格
        avgSpaces:=spacesCount/(wordsCount-1)
        // 将多余的空格平均到左侧部分字符串之间
        extraSpaces:=spacesCount%(wordsCount-1)
        // 左侧部分字符串之间多一个空格，可以达成条件空格不均匀分配时左侧空格多于右侧
        s1:=strings.Join(words[left:left+extraSpaces+1],fillSpace(avgSpaces+1))
        s2:=strings.Join(words[left+extraSpaces+1:right],fillSpace(avgSpaces))
        res=append(res,s1+fillSpace(avgSpaces)+s2)
    }
}
```

## 复杂度分析

- **时间复杂度:** $$O(n)$$
- **空间复杂度:** $$O(n)$$
