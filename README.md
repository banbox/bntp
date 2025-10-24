[中文](README.cn.md)

# BNTP - Go NTP Network Time Synchronization Library

BNTP is a high-performance Go NTP (Network Time Protocol) client library for obtaining and synchronizing network standard time.

## ✨ Features

- **🌍 Multi-Region Support**: Built-in NTP servers for Mainland China, Hong Kong, Taiwan, Japan, South Korea, Singapore, and global regions
- **⚡ High Performance**: Uses atomic operations and caching mechanisms to provide millisecond-level timestamp retrieval
- **💾 Persistence**: Automatically saves time offset to local file for quick recovery after restart
- **🔄 Auto Sync**: Supports periodic automatic time offset refresh (optional)

## 📦 Installation

```bash
go get github.com/banbox/bntp
```

## 🚀 Usage Examples

### Basic Usage

```go
package main

import (
    "log"
    "time"
    "github.com/banbox/bntp"
)

func main() {
    // LangGlobal, LangZhCN, LangZhHK, LangZhTW, LangJaJP, LangKoKr, LangZhSg, LangNone(disabled, default)
    bntp.LangCode = bntp.LangGlobal

    // Get current system time
    sysNow := time.Now()
    log.Printf("Current system time: %s\n", sysNow.Format(time.RFC3339))

    // Get corrected current time
    now := bntp.Now()
    log.Printf("Current standard time: %s\n", now.Format(time.RFC3339))

    // Get corrected timestamp (milliseconds)
    timestamp := bntp.UTCStamp()
    log.Printf("Current timestamp: %d\n", timestamp)

    // Get time offset (milliseconds)
    offset := bntp.GetTimeOffset()
    log.Printf("Time offset: %d ms\n", offset)
}
```

### Custom Configuration

```go
package main

import (
    "log"
    "time"
    "github.com/banbox/bntp"
)

func main() {
    _, err := bntp.SetTimeSync(
        // Set country/region code
        bntp.WithCountryCode(bntp.LangJaJP),
        
        // Set offset file save path (defaults to user cache directory)
        // bntp.WithFilePath("/path/to/ntp_offset.json"),
        
        // Enable loop refresh
        bntp.WithLoopRefresh(true),
        
        // Set sync period (recommended >= 1 hour)
        bntp.WithSyncPeriod(6 * time.Hour),
        
        // Set random fluctuation rate (between 0-1)
        bntp.WithRandomRate(0.15),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Manually refresh time offset
    ts := bntp.GetTimeSync()
    if err := ts.Refresh(); err != nil {
        log.Printf("Refresh failed: %v\n", err)
    }
    
    // Get corrected timestamp (milliseconds)
    timestamp := bntp.UTCStamp()
    log.Printf("Current timestamp: %d\n", timestamp)
}
```

## 📄 License

This project is licensed under the BSD-2 License. See the [LICENSE](LICENSE) file for details.
