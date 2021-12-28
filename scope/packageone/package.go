package packageone

import (
	"fmt"
)

var privateVar = "i am private"
var PackageVar = "PackageOne var"

func notExported() {

}

func ToString() string {
	return PackageVar
}

func PrintMe(myVar string, blockVar string) {
	fmt.Println("=> ", myVar, blockVar, PackageVar)
}
