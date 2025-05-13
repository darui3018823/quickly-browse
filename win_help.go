//go:build windows
// +build windows

// 2025 darui3018823 All rights reserved.
package main

import (
	"syscall"
	"unsafe"
)

var (
	user32          = syscall.NewLazyDLL("user32.dll")
	procMessageBoxW = user32.NewProc("MessageBoxW")
)

func showWindowsHelpDialog() {
	title := "Quickly-Browse Help"
	content := `Usage:
  q-brow [options] "search terms"

Options:
  -g        Google
  -y        YouTube
  -t        Twitter
  -d        DuckDuckGo
  --help    Show this message
`

	procMessageBoxW.Call(
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(content))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		0x00000040, // MB_ICONINFORMATION
	)
}
