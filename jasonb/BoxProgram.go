package main

import (
"ethos/syscall"
"ethos/ethos"
ethosLog "ethos/log"
"ethos/efmt"	
)

func main () {
me := syscall.GetUser()
path := "/user/" + me + "/myDir/"

status := ethosLog.RedirectToLog("BoxProgram")
if status != syscall.StatusOk {
efmt.Fprintf(syscall.Stderr, "Error opening %v: %v\n", path, status)
syscall.Exit(syscall.StatusOk)
}

var coor1 Box
coor1.x1=-10
coor1.y1=2
coor1.x2=-5
coor1.y2=10

var coor2 Box
coor2.x1 = coor1.y1
coor2.y1 = coor1.x1
coor2.x2 = coor1.y2
coor2.y2 = coor1.x2

var midx float32
var midy float32
midx = float32(coor1.x1 + coor1.x2)
midy = float32(coor1.y1 + coor1.y2)
midx = midx / 2
midy = midy / 2

efmt.Printf("Points of Box \n [x1,y1] are : [ %v,%v ]\n [x2,y2] are : [ %v,%v ]\n", coor1.x1, coor1.y1, coor1.x2, coor1.y2)
efmt.Printf("Middle Coordinate of the Box is: [ %v,%v ]\n",midx,midy)
efmt.Printf("Inverse Box coordinates are -\nLower Left : [%v,%v]\nUpper Right : [%v,%v]\n",coor2.x1,coor2.y1,coor2.x2,coor2.y2)

fd, status := ethos.OpenDirectoryPath(path)
coor1.Write(fd)
coor1.WriteVar(path)

efmt.Print("Project Completed !!!\n")

}
