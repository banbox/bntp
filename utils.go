package bntp

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetCacheDir() (string, error) {
	switch runtime.GOOS {
	case "windows":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(homeDir, "AppData", "Local"), nil
	case "darwin": // macOS
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(homeDir, "Library", "Caches"), nil
	default:
		// 默认linux
		// 优先使用 XDG_CACHE_HOME 环境变量
		if xdgCache := os.Getenv("XDG_CACHE_HOME"); xdgCache != "" {
			return xdgCache, nil
		}
		// 回退到 $HOME/.cache
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(homeDir, ".cache"), nil
	}
}
