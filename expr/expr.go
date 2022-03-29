package expr

// 如为真返回参数一，否则返回参数二
func If[T any](e bool, v1, v2 T) T {
	if e {
		return v1
	}
	return v2
}
