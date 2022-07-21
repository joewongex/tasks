package g

import (
	"fmt"
	"runtime"
	"testing"
)

func TestNotify(t *testing.T) {
	// err := os.Chdir("./root/code/go/tasks")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// wd, err := os.Getwd()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Println(wd)
	// Init()

	_, file, _, _ := runtime.Caller(1)
	fmt.Println("xxxx" + file)
}
