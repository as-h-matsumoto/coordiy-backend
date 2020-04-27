


package main


// https://tour.golang.org/

// import "fmt" <- ダメな礼 ↓が良い例
import (
	"fmt"
	"os"
	"time"
	"math"
	"math/rand"
	"math/cmplx"
	"io"
	"strings"
	"sync"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	// 型　variable
	var intv int
	var float64v float64
	var stringv string
	var boolv bool
	// rand.Intn <- 最初の文字は大文字にする（Goのルールとして、大文字はエクスポートされる。パッケージをインポートするときはエクスポートされた名前のみ参照できる。）
	fmt.Printf("\nintv: %d, float64v: %f, stringv: %s, boolv: %t, randInt: %d, Sqrt: %g\n\n", intv, float64v, stringv, boolv, rand.Intn(10), math.Sqrt(7))

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var int2 int = 2
	//var int3 int = 3
	var int4 int = 4
	var int5 int = 5
	var tea float64 = math.Sqrt( float64(20) )
	var teaa float64 = math.Sqrt( float64(int4*int5) )
	var teaaa float64 = math.Sqrt( float64( (int2*int2) * int5 ) )
	fmt.Println(tea, teaa, teaaa)

	// 宣言代入パターン
	var msg1 string
	msg1 = "hello world msg1"
	var msg2 string = "Hello World msg2"
	var msg3 = "Hello Hello msg3"
	msg4 := "Super Hello msg4"

	// 複数の宣言代入パターン
	var c3, d4 int
	c3, d4 = 10, 15
	c3, d4 = 9, 14 //上書き
	var (
	  e int
	  f string
	)
	e = 20
	f = "hoge"
	var (
	  c2 = 20
	  d2 = "hoge"
	)
	fmt.Print("\n\n")


	fmt.Print("\n print \n")
	fmt.Print(msg1)
	// 接頭辞 F : Fprint(), Fprintf(), Fprintln() : 書き込み先を明示的に指定
	fmt.Fprint(os.Stdout, msg2)
	// 接頭辞 S : Sprint(), Sprintf(), Sprintln() : 変数に代入する際
	msgs := fmt.Sprintln(msg3)
	// 接尾辞 f : Printf(), Fprintf(), Sprintf() : フォーマットを指定
	fmt.Printf("%s\n", msg4)
	// 接尾辞 ln : Println(), Fprintln(), Sprintln() : 改行が追加
	fmt.Println(msgs, c3, d4, e, f, c2, d2)
	fmt.Print("\n\n")


	// 関数
	fmt.Print(" 関数 ")
	sayHi()
    sayName("gcfuji")
	// 関数もデータ型なので、変数に代入可能。その際は関数名はいらない
	funcf := func(a int, b int) (int, int) { return b, a }
	fmt.Println(getMessage("gcfuji"), getHelloMessage("Gemcook"))
	fmt.Println(swap(5, 2))
	fmt.Println(funcf(234, 2342))
	fmt.Print("\n\n")


	// 配列
	fmt.Print("\n 配列 \n")
	var arra [5]int // a[0] - a[4]
	arra[2] = 3
    arra[4] = 10
	arrb := [3]int{1, 3, 6}
	arrc := [...]int{2, 4, 7, 90, 5, 19, 12, 421}
	fmt.Println(arra, arrb, arrc)
	fmt.Print("\n\n")


	// スライス
	fmt.Print("\n スライス 配列の参照 \n")
	// スライス　：　スライスは配列への参照なので、値を変更すると元の配列も値が変更される
	// a[:]		全要素
	// a[2:]	2番目から最後まで
	// a[2:5]	2番目から(5-1)番目まで
	// a[:5]	初めから(5-1)番目まで
    barrc := arrc[2:4] // [7, 90]
    carrc := arrc[2:]  // [7 90 5 19 12 421]
    darrc := arrc[:4]  // [2 4 7 90]
    earrc := arrc[:]   // [2 4 7 90 5 19 12 421]
	fmt.Println(barrc, carrc, darrc, earrc)
	// スライスの長さを取得
	fmt.Println(len(barrc))
	fmt.Println(len(earrc))
    // スライスの最大容量を取得
	fmt.Println(cap(barrc))
	fmt.Println(cap(earrc))
	// スライス変更
	barrc[0] = 1092
	barrc[1] = 29487
	fmt.Println(arrc, barrc)
	barrcChange1(barrc)
	fmt.Println(arrc, barrc)
	// makeでスライス作成
	s1 := make([]int, 3) // [0 0 0]
	s2 := []int{1, 3, 5} // 値もいれる。配列の宣言と似ている
	s3 := append(s2, 8, 2, 10)
	s3copy := make([]int, len(s3))
	ansS3copy := copy(s3copy, s3)
	fmt.Println(s1, s2, s3, ansS3copy, s3copy)
	fmt.Print("\n\n")


	// マップ
	fmt.Print("\n マップ \n")
	// map : マップ ：　連想配列（key=>value的な）
	m1 := map[string]int{"fujimoto": 100, "arita": 200}
	m1v, m1ok := m1["fujimoto"]
	fmt.Println(m1, m1["fujimoto"], m1v, m1ok)
	delete(m1, "fujimoto")
	fmt.Println(m1)
	fmt.Print("\n\n")


	// ポインタ
	fmt.Print("\n ポインタ \n")
	var pa *int // int型ポインタ
    pa = &e     // ポインタにアドレスをいれる
	fmt.Printf("int address: %p, int value: %d\n",pa, *pa)
	var pmap *map[string]int // map型ポインタ
    pmap = &m1
	fmt.Printf("map address: %p, map value: %v\n",pmap, *pmap)
	var pfunc *func(a int, b int) (int, int) // func型ポインタ
    pfunc = &funcf
	fmt.Printf("func address: %p, func value: %v\n",pfunc, *pfunc)
	fmt.Print("\n\n")

	
	// if
	// score := 52; はなくてもいい。
	// したのパターンはscoreはブロック内でしか利用できない。
	fmt.Print("\n if \n")
    if score := 52; score > 80 {
        fmt.Println("Great!")
    } else if score > 60 {
        fmt.Println("Good!")
    } else {
        fmt.Println("soso")
	}
	fmt.Print("\n\n")
	

	// switch
	fmt.Print("\n switch \n")
    switch m1["arita"] {
    case 100:
        fmt.Println("fujimoto")
    case 200:
        fmt.Println("arita")
    default:
        fmt.Println("wrong")
    }
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Today is Saturday.")
	case time.Sunday:
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Today is a weekday.")
	}
	fmt.Print("\n\n")


	// for (whileはない)
	fmt.Print("\n for \n")
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		} else if i == 8 {
			break
		}
		fmt.Printf("i: %d, ",i)
	}
	fmt.Print("\n")
	i := 10
	for i < 20 {
		fmt.Printf("i: %d, ",i)
		i++
	}
	fmt.Print("\n\n")
	/* これはダメ
	for i := 20; i < 30 {
		fmt.Printf("i: %d, ",i)
		n++
	}
	*/


	// range （配列の数だけ回す）
	// スライス
	fmt.Print("\n range （配列の数だけ回す） \n")
	srange := []int{2, 3, 8}
	for i, v := range srange {
		fmt.Printf("i: %v, v: %v, ", i, v)
	}
	fmt.Print("\n")
	// ブランク修飾子
	// _にしておくと何が入ってきてもそれを廃棄してくれるという仕様
	for _, v := range srange {
		fmt.Printf("v: %v, ", v)
	}
	fmt.Print("\n")
	// マップ
	mrange := map[string]int{"fujimoto": 200, "arita": 300}
	for k, v := range mrange {
		fmt.Printf("k: %v, v: %v, ", k, v)
	}
	fmt.Print("\n\n")
	

	//構造体
	fmt.Print("\n 構造体 \n")
	// newでuser構造体分の領域を確保して、初期化して、そのポインタを返す
	u1 := new(user)
	// ポインタが返ってきているので下記の書き方でもOK
	// (*u).name = "fujimoto"
	u1.name = "fujimoto"
	u1.score = 200
	fmt.Println(u1, *u1, (*u1).name, (*u1).score)
	// 初期化する時に直接値を与える事も可能
	// こちらの場合はポインタではない構造体のデータが入ってくる。
	u2 := user{name: "arita", score: 300}
	fmt.Println(u2)
	stte := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(stte)
	mVertex111 = make(map[string]Vertex)
	mVertex111["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(mVertex111["Bell Labs"])
	fmt.Print("\n\n")


	// メソッド
	// 構造体と関数をレシーバーで紐付けてメソッドとする
	fmt.Print("\n メソッド \n")
	uMeso := user{name: "fujimoto", score: 500}
    uMeso.hit()
	uMeso.show()
	// newポインタのパターンでもレシーバーを経由すると結果は同じになる
	uMesoTest := new(user)
	uMesoTest.name = "tasoya"
	uMesoTest.score = 20971
	uMesoTest.hit()
	uMesoTest.show()
	fmt.Print("\n\n")


	// インタフェース
	// メソッドシグネチャの集合として利用される。
	fmt.Print("\n メソッド \n")
	greeters := []greeter{japanese{}, american{}}
    for _, greeter := range greeters {
		greeter.greet()
		show(greeter)
	}
	fmt.Print("\n\n")


	// goroutine
	// goを付けて並行処理にする
	fmt.Print("\n goroutine \n")
    //go task1()
    //go task2()
    // goroutineが終わる前に main 関数自体が終わる為、待ち時間をつける。
	//time.Sleep(time.Second * 3)
	fmt.Print("\n\n")
	
	
	// channel
	// goroutine では return ができないので chan をつけてわたす
	// channel はスライスなどと同じ参照型のデータでmake で作成
	fmt.Print("\n channel goにリターンさせる方法 同期重要 \n")
    //result := make(chan string)
    // goを付けて並行処理にする
    //go taskChan1(result)
    //go task2()
    // resultの中に何も入ってなければ、入ってくるまで待つ仕様になっている
    //fmt.Println(<-result)
    // goroutineが終わる前に main 関数自体が終わる為、待ち時間をつける。
	//time.Sleep(time.Second * 3)
	fmt.Print("\n\n")


	// defer (遅延処理)
	// defer 定義ではない物が最初に終わり、defer定義はしたから順に処理される
	fmt.Print("\n defer \n")
	//defer fmt.Println("defer。最後に表示されてしまいます。")
	//defer fmt.Println("サポート")
	//defer fmt.Println("world")
	fmt.Println("hello")
	//defer fmt.Println("系統")
	//defer fmt.Println("Ieko3#")
	//for i := 0; i < 10; i++ {
	//	defer fmt.Println(i)
	//}
	//defer fmt.Println("\n\nstart defer。最後に表示されてしまいます。")
	fmt.Print("\n\n")

	
	// クロージャー
	fmt.Print("\n クロージャー \n")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	fmt.Print("\n クロージャー2 \n")
	pos2, neg2 := adder2(), adder2()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		fmt.Println(
			pos2(i),
			neg2(-2*i),
		)
	}
	fmt.Print("\n\n")


	// フェボナッチ数列計算をmapで実施した例
	memo := make(map[int]int)
    for i := 0; i < 40; i++ {
		Memorize(i, memo)
		fmt.Printf("i: %v, memo[i]: %v\n", i,memo[i])
	}
	fmt.Print("\n\n")
	

	// タイプアサーション
	fmt.Print("\n タイプアサーション インタフェース \n")
	doAI(21)
	doAI("hello")
	doAI(true)
	fmt.Print("\n\n")
   

	// strings インタフェース
	// importしているパッケージはインタフェースとのこと？
	fmt.Print("\n インタフェース \n")
	rrr := strings.NewReader("Hello, Reader!")
	fmt.Println(rrr);

	bbb := make([]byte, 10)
	fmt.Println(bbb[2:6]);
	bbb[2] = 255
	fmt.Println(bbb);
	
	for {
		n, err := rrr.Read(bbb)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, bbb)
		fmt.Printf("bbb[:n] = %q\n", bbb[:n])
		if err == io.EOF {
			break
		}
	}
	fmt.Print("\n\n")


	// チャネル
	fmt.Print("\n チャネル \n")
	c := make(chan int, 10)
	go chFibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Print("\n\n")
	//上記の様に１０個の場合はrangeが有効だが、処理の中で終了を判定する場合は、以下の形でclose channelにシグナルを送る方法がある。
	cchan := make(chan int)
	cquit := make(chan int)
	fmt.Println(cchan, cquit);
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("<-c: %v\n", <-cchan)
		}
		cquit <- 0
	}()
	chFibonacci2(cchan, cquit)
	fmt.Print("\n\n")




	// lock unlock 安全に 更新する。
	fmt.Print("\n lock \n")
	cSafeCounter := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go cSafeCounter.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println("lockValue: %d", cSafeCounter.Value("somekey"))
	fmt.Print("\n\n")


	Crawl("https://golang.org/", 4, fetcher)

	

	/*
	fmt.Print("\n tick Boom \n")
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Print("\n end tick boom \n")
	*/

}




//　クローラー
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}
// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}
// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult
type fakeResult struct {
	body string
	urls []string
}
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}
// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}




// lock unlock 構造体の編集を安全に行う。
// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}
// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}
// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}



//
func chFibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func chFibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
		    fmt.Printf("before x: %d, y: %d\n", x, y)
			x, y = y, x+y
			fmt.Printf("after x: %d, y: %d\n", x, y)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}


// タイプアサーション（タイプスイッチ）インタフェースの色々？
func doAI(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// フェボナッチ数列計算をmapで実施した例
func Memorize(n int, memo map[int]int) int {
    if n < 2 {
        return n
    }
    if _, ok := memo[n]; !ok {
        memo[n] = Memorize(n-2, memo) + Memorize(n-1, memo)
    }
    return memo[n]
}

// クロージャー
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
// クロージャー2
func adder2() func(int) int {
	var sum = 10;
	return func(x int) int {
		sum += x
		return sum
	}
}


// channel
func taskChan1(result chan string) {
    // 重い処理を想定
    time.Sleep(time.Second * 2)
    fmt.Println("task1 finished!")

    result <- "task1 result"
}


// goroutine
func task1() {
    // 重い処理を想定
    time.Sleep(time.Second * 2)
    fmt.Println("task1 finished!")

}
func task2() {
    fmt.Println("task2 finished!")
}


// 空のインタフェース型で引数を受け取る
func show(t interface{}) {
	// 型アサーション
	// 2つの値が返る
	_, ok := t.(japanese)
	// okを使って条件分岐
	if ok {
		fmt.Println("i am japanese")
	} else {
		fmt.Println("i am not japan")
	}

	// 型Switch
	switch t.(type) {
	case japanese:
		fmt.Println("僕は日本人だよ")
	default:
		fmt.Println("僕は日本人じゃないよ")
	}
}


// インタフェース
type greeter interface {
    greet()
}
type japanese struct{}
type american struct{}
func (ja japanese) greet() {
    fmt.Println("こんにちは！")
}
func (us american) greet() {
    fmt.Println("Hello")
}


// 構造体
type Vertex struct {
	Lat, Long float64
}
var mVertex111 map[string]Vertex
type user struct {
    // フィールド
    name  string
    score int
}
// レシーバー　コピーで渡す
func (u *user) show() {
	fmt.Printf("name: %s, socre: %d\n", u.name, u.score)
	fmt.Println(u)
}
// 参照渡しパターン
func (u *user) hit() {
    u.score++
}


// 関数
// 引数なし関数
func sayHi() { fmt.Println("hi!") }
// 引数あり関数
func sayName(name string) { fmt.Println(name) }
// return関数
func getMessage(name string) string {
    msg := "hi! my name is " + name
    return msg
}
// 名前付きreturn関数
// 関数内で使った変数名を返す
func getHelloMessage(name string) (msg string) {
    msg = "Hello " + name
    return
}
// 複数の返り値を返す事ができる
func swap(a int, b int) (int, int) { return b, a }
// スライスの値の計算
func barrcChange1(a []int) {
	a[0] += 120;
}




