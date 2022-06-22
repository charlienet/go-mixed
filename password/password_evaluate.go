package password

import (
	"fmt"
	"regexp"
)

const (
	LevelD = iota
	LevelC
	LevelB
	LevelA
	LevelS
)

var patterns = []*regexp.Regexp{
	regexp.MustCompile(`[0-9]+`),          // 数字
	regexp.MustCompile(`[a-z]+`),          // 小字字母
	regexp.MustCompile(`[A-Z]+`),          // 大写字母
	regexp.MustCompile(`[~!@#$%^&*?_-]+`), // 特殊符号
}

func PasswordEvaluate(minLength, maxLength int, minLevel int, password string) error {
	// 首先校验密码长度是否在范围内
	if len(password) < minLength {
		return fmt.Errorf("无效密码: 密码长度小于%d个字符", minLength)
	}
	if len(password) > maxLength {
		return fmt.Errorf("无效密码: 密码的长度大于%d个字符", maxLength)
	}

	var level int = LevelD
	for _, pattern := range patterns {
		if pattern.MatchString(password) {
			level++
		}
	}

	// 如果最终密码强度低于要求的最低强度，返回并报错
	if level < minLevel {
		return fmt.Errorf("密码强度不满足当前策略要求。")
	}
	return nil
}
