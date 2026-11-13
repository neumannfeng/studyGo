//go:build windows

package main

import (
	"syscall"
	"unsafe"
)

const enableVirtualTerminalProcessing = 0x0004

// enableANSI 让传统 Windows 控制台也能解释 ANSI 颜色。
// 返回 false 表示当前控制台不支持，调用方应退回无颜色输出。
func enableANSI() bool {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	getStdHandle := kernel32.NewProc("GetStdHandle")
	getConsoleMode := kernel32.NewProc("GetConsoleMode")
	setConsoleMode := kernel32.NewProc("SetConsoleMode")

	handle, _, _ := getStdHandle.Call(uintptr(0xFFFFFFF5)) // STD_OUTPUT_HANDLE
	var mode uint32
	ok, _, _ := getConsoleMode.Call(handle, uintptr(unsafe.Pointer(&mode)))
	if ok == 0 {
		return false
	}
	changed, _, _ := setConsoleMode.Call(handle, uintptr(mode|enableVirtualTerminalProcessing))
	return changed != 0
}
