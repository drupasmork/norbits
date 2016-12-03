package main

import (
"ethos/syscall"
"ethos/ethos"
ethosLog "ethos/log"
"ethos/efmt"	
)

func main () {

path := "/programs/"

status := ethosLog.RedirectToLog("DirInfo")
if status != syscall.StatusOk {
efmt.Fprintf(syscall.Stderr, "Error opening %v: %v\n", path, status)
syscall.Exit(syscall.StatusOk)
}

var result = DirR { count: 0 , size: 0 }
fname := ".."

fd, status := ethos.OpenDirectoryPath(path)
if status != syscall.StatusOk {
efmt.Fprintf(syscall.Stderr, "Error Path Opening %v: %v\n", path, status)
syscall.Exit(syscall.StatusOk)
}

efmt.Printf("\n%-18s %-18s \n\n","File Name","Size(bytes)")

for fname!="null" {

name, stat := ethos.GetNextName(fd,fname)

if stat != syscall.StatusOk {
break
efmt.Fprintf(syscall.Stderr, "Error Getting Name %v: %v\n", name, stat)
syscall.Exit(syscall.StatusOk)
}

fname = string(name)
result.count = result.count + 1
filepath := path + fname

finfo, sta := ethos.GetFileInformationPath(filepath)

efmt.Printf("%-18s %-18v \n", string(name), finfo.Size)

if sta != syscall.StatusOk {
efmt.Fprintf(syscall.Stderr, "Error getting file info for %v: %v", finfo, sta)
syscall.Exit(syscall.StatusOk)
}

result.size = result.size + finfo.Size

}

efmt.Printf("%15s \n","Result")
efmt.Printf("Total files in directory : %v\n", result.count)
efmt.Printf("Size of directory : %v bytes\n", result.size)

efmt.Print("Project Completed !!!\n")

}
