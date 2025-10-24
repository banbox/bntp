[English](./README.md)

# BNTP - Go NTP 网络时间同步库

BNTP 是一个高性能的 Go 语言 NTP（网络时间协议）客户端库，用于获取和同步网络标准时间。

## ✨ 特性

- **🌍 多地区支持**：内置中国大陆、香港、台湾、日本、韩国、新加坡及全球 NTP 服务器
- **⚡ 高性能**：使用原子操作和缓存机制，提供毫秒级时间戳获取
- **💾 持久化**：自动保存时间偏移到本地文件，重启后快速恢复
- **🔄 自动同步**：支持定期自动刷新时间偏移（可选）

## 📦 安装

```bash
go get github.com/banbox/bntp
```

## 🚀 使用示例

### 基础使用

```go
package main

import (
    "log"
    "time"
    "github.com/banbox/bntp"
)

func main() {
    // LangGlobal, LangZhCN, LangZhHK, LangZhTW, LangJaJP, LangKoKr, LangZhSg, LangNone(不启用，默认)
    bntp.LangCode = bntp.LangGlobal

    // 获取系统当前时间
    sysNow := time.Now()
    log.Printf("当前系统时间: %s\n", sysNow.Format(time.RFC3339))

    // 获取校正后的当前时间
    now := bntp.Now()
    log.Printf("当前标准时间: %s\n", now.Format(time.RFC3339))

    // 获取校正后的时间戳（毫秒）
    timestamp := bntp.UTCStamp()
    log.Printf("当前时间戳: %d\n", timestamp)

    // 获取时间偏移量（毫秒）
    offset := bntp.GetTimeOffset()
    log.Printf("时间偏移: %d ms\n", offset)
}
```

### 自定义配置

```go
package main

import (
    "log"
    "time"
    "github.com/banbox/bntp"
)

func main() {
    _, err := bntp.SetTimeSync(
        // 设置国家/地区代码
        bntp.WithCountryCode(bntp.LangJaJP),
        
        // 设置偏移文件保存路径（默认使用用户缓存目录）
        // bntp.WithFilePath("/path/to/ntp_offset.json"),
        
        // 启用循环刷新
        bntp.WithLoopRefresh(true),
        
        // 设置同步周期（建议 >= 1 小时）
        bntp.WithSyncPeriod(6 * time.Hour),
        
        // 设置随机波动率（0-1 之间）
        bntp.WithRandomRate(0.15),
    )
    if err != nil {
        log.Fatal(err)
    }

    // 手动刷新时间偏移
    ts := bntp.GetTimeSync()
    if err := ts.Refresh(); err != nil {
        log.Printf("刷新失败: %v\n", err)
    }
    
    // 获取校正后的时间戳（毫秒）
    timestamp := bntp.UTCStamp()
    log.Printf("当前时间戳: %d\n", timestamp)
}
```

## 📄 许可证

本项目采用 BSD-2 许可证。详见 [LICENSE](LICENSE) 文件。
