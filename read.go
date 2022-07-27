package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"syscall"
	"unsafe"
)

var pid = flag.Uint("p", 0, "PID of progress")
var address = flag.Uint("a", 0, "Address for Beginning")
var size = flag.Uint("s", 0, "Size of memory to read")
var filename = flag.String("o", "dumpData", "name of file to output")
var printType = flag.String("t", "%q", "type to format for showing")
func main() {
	flag.Parse()
	var testString []byte;
	testString=[]byte("If you can see me,you have succeeded")
	if *pid==0 {
       *pid=uint(syscall.Getpid())
       *address=uint(uintptr(unsafe.Pointer(&testString[0])))
       *size= uint(len(testString))
		fmt.Printf("Address of test var is %x,len:%d \n",*address,len(testString))
	}
	var temp []byte = make([]byte, *size)
	var local syscall.Iovec
	    local.Base=&temp[0]
	    local.Len=uint32(*size)
	var remote syscall.Iovec
	remote.Base= (*byte)(unsafe.Pointer(uintptr(*address)))
	remote.Len=uint32(*size)
	fmt.Printf("Remote memory address in iov:%x,len:%d \n",uintptr(unsafe.Pointer(remote.Base)),remote.Len)
	fmt.Printf("Reading memory of progress(pid:%d) from %x to %x \n",*pid,*address,*address+*size)
	fmt.Printf("Parameters of Syscall6 were: %d %d %x %d %x %d %d \n",syscall.SYS_PROCESS_VM_READV,
		uintptr(*pid),
		uintptr(unsafe.Pointer(&local)),
		1,
		uintptr(unsafe.Pointer(&remote)),
		1,
		0)
	_, _, e1 :=syscall.Syscall6(syscall.SYS_PROCESS_VM_READV,
		uintptr(*pid),
		uintptr(unsafe.Pointer(&local)),
		1,
		uintptr(unsafe.Pointer(&remote)),
		1,
		0)

	if e1 != 0 {
		var err error
		err = syscall.Errno(e1)
		fmt.Println("There are something wrong:",err)
	}else{
		if *printType=="%q"{
			fmt.Printf("Things I got :"+*printType+" ...\n",string(temp))
		}else {
			fmt.Printf("Things I got :"+*printType+" ...\n", temp)
		}
		ioutil.WriteFile(*filename, temp, 0644)
	}
}
