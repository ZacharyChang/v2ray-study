package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func getCurrentPkg() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Base(path), nil
}

func main() {

	filename := "errors.generated.go"

	pkg, err := getCurrentPkg()
	if err != nil {
		log.Fatal("Failed to get current package: ", err.Error())
		return
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to generate errors.generated.go: %v", err)
		return
	}

	fmt.Fprintln(file, "package", pkg)
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "import \"v2ray.com/core/common/errors\"")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "type errPathObjHolder struct {}")
	fmt.Fprintln(file, "func newError(values ...interface{}) *errors.Error { return errors.New(values...).WithPathObj(errPathObjHolder{}) }")
	file.Close()

	res, err := exec.Command("go", "fmt", filename).Output()
	fmt.Println("result:", string(res))
}
