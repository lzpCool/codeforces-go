直接比较字符串即可。

```py [sol1-Python3]
class Solution:
    def haveConflict(self, event1: List[str], event2: List[str]) -> bool:
        return event1[0] <= event2[1] and event1[1] >= event2[0]
```

```go [sol1-Go]
func haveConflict(event1, event2 []string) bool {
	return event1[0] <= event2[1] && event1[1] >= event2[0]
}
```

#### 复杂度分析

- 时间复杂度：$O(1)$。
- 空间复杂度：$O(1)$。