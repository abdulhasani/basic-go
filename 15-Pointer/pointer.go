package main

import (
	"fmt"
)

/**function ini berparamter, dimana paramter tidak dapat diubah valuenya
alias tidak refrence */
func zeroval(ival int) {
	ival = 0
}

/**function ini berparamter, dimana paramter bisa diset ulang/dirubah
alias refrence*/
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial1:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}
