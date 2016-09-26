// monoserv compute package
// author: davethecust

package c_src

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lcompute

#include <compute.h>

*/
import "C"

func PrintFromCFiles() {
	C.printPlayer1()
	C.printPlayer2()
}

func main() {
	PrintFromCFiles()
}
