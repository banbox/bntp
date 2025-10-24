package bntp

import (
	"log"
	"testing"
	"time"
)

// 使用示例
func TestTimeSync(t *testing.T) {
	// 初始化时间同步器，使用中国区域设置
	SetTimeSync(
		WithCountryCode("zh-CN"),
		WithRandomRate(0.1),
	)
	// 获取校正后的时间戳
	offset := GetTimeOffset()
	trueStr := Now().Format(time.RFC3339)
	curStr := time.Now().Format(time.RFC3339)
	log.Printf("[INFO] ntp sync res, true=%s, cur=%s, offset=%d\n", trueStr, curStr, offset)
}
