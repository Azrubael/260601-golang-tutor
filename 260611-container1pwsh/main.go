package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"unsafe"
)

var (
    kernel32                     = syscall.NewLazyDLL("kernel32.dll")
    procCreateJobObject          = kernel32.NewProc("CreateJobObjectW")
    procAssignProcessToJobObject = kernel32.NewProc("AssignProcessToJobObject")
    procSetInformationJobObject  = kernel32.NewProc("SetInformationJobObject")
    procOpenProcess              = kernel32.NewProc("OpenProcess")
)

const PROCESS_ALL_ACCESS = 0x1F0FFF

type JOBOBJECT_BASIC_LIMIT_INFORMATION struct {
    PerProcessUserTimeLimit int64
    PerJobUserTimeLimit     int64
    LimitFlags              uint32
    MinimumWorkingSetSize   uintptr
    MaximumWorkingSetSize   uintptr
    ActiveProcessLimit      uint32
    Affinity                uintptr
    PriorityClass           uint32
    SchedulingClass         uint32
}

type IO_COUNTERS struct {
    ReadOperationCount  uint64
    WriteOperationCount uint64
    OtherOperationCount uint64
    ReadTransferCount   uint64
    WriteTransferCount  uint64
    OtherTransferCount  uint64
}

type JOBOBJECT_EXTENDED_LIMIT_INFORMATION struct {
    BasicLimitInformation JOBOBJECT_BASIC_LIMIT_INFORMATION
    IoInfo                IO_COUNTERS
    ProcessMemoryLimit    uintptr
    JobMemoryLimit        uintptr
    PeakProcessMemoryUsed uintptr
    PeakJobMemoryUsed     uintptr
}

type JOBOBJECT_CPU_RATE_CONTROL_INFORMATION struct {
    ControlFlags uint32
    CpuRate      uint32
}

func main() {
    cmd := exec.Command("powershell.exe", "-NoLogo", "-NoProfile", "-Command", /*"Get-Date",*/ "Write-Output 'Привіт з контейнера!'")
    cmd.Env = append(os.Environ(), "MY_CONTAINER_ENV=TestEnv")
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

		// Перенаправлення виводу
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

    if err := cmd.Start(); err != nil {
        fmt.Println("Помилка запуску:", err)
        return
    }

    hJob, _, err := procCreateJobObject.Call(0, 0)
    if hJob == 0 {
        fmt.Println("Не вдалося створити Job Object:", err)
        return
    }

    hProcess, _, err := procOpenProcess.Call(
        uintptr(PROCESS_ALL_ACCESS),
        uintptr(0),
        uintptr(cmd.Process.Pid),
    )
    if hProcess == 0 {
        fmt.Println("Не вдалося отримати HANDLE процесу:", err)
        return
    }

    r1, _, err := procAssignProcessToJobObject.Call(hJob, hProcess)
    if r1 == 0 {
        fmt.Println("Не вдалося прив’язати процес:", err)
        return
    }

    // RAM ліміт 100 MB
    memInfo := JOBOBJECT_EXTENDED_LIMIT_INFORMATION{}
    memInfo.ProcessMemoryLimit = 100 * 1024 * 1024
    memInfo.JobMemoryLimit = 100 * 1024 * 1024
    memInfo.BasicLimitInformation.LimitFlags = 0x200 | 0x100 // JOB_OBJECT_LIMIT_PROCESS_MEMORY + JOB_OBJECT_LIMIT_JOB_MEMORY

    r2, _, err := procSetInformationJobObject.Call(
        hJob,
        uintptr(9), // JobObjectExtendedLimitInformation
        uintptr(unsafe.Pointer(&memInfo)),
        unsafe.Sizeof(memInfo),
    )
    if r2 == 0 {
        fmt.Println("Помилка встановлення RAM ліміту:", err)
    }

    // CPU ліміт 30%
    cpuInfo := JOBOBJECT_CPU_RATE_CONTROL_INFORMATION{
        ControlFlags: 0x1 | 0x4, // ENABLE + HARD_CAP
        CpuRate:      3000,      // 30% (30 * 100)
    }

    r3, _, err := procSetInformationJobObject.Call(
        hJob,
        uintptr(15), // JobObjectCpuRateControlInformation
        uintptr(unsafe.Pointer(&cpuInfo)),
        unsafe.Sizeof(cpuInfo),
    )
    if r3 == 0 {
        fmt.Println("Помилка встановлення CPU ліміту:", err)
    }

    cmd.Wait()

		// Повідомлення після завершення
    fmt.Println("✅ Контейнер завершив роботу успішно і зупинений.")
}
