


package main


// https://tour.golang.org/

// import "fmt" <- ダメな礼 ↓が良い例
import (
	"fmt"
)

func main() {



	a, b := getHelloMessage("Gemcook")
	fmt.Println(a, b);

	val := url.Values{}
	fmt.Println(a, b);



}



// 名前付きreturn関数
// 関数内で使った変数名を返す
func getHelloMessage(name string) (msg string, msg2 string) {
	msg = "Hello " + name
	msg2 = msg + " thanks"
	// msg, msg2も関数で定義した段階で作られている。
    return
}




