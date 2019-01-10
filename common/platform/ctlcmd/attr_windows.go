// +build windows
// Note: this file will be built under windows
package ctlcmd

import "syscall"

func getSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		HideWindow: true,
	}
}
