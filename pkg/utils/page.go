package utils

// 计算总页数
func TotalPage(size, total int) int {
	if size == 0 {
		return 0
	}
	t := total / size
	if total%size > 0 {
		t += 1
	}
	return t
}
