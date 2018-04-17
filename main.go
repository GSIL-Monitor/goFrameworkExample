package main

//
//import (
//	"flag"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//	_ "net/http/pprof"
//	"sync"
//	"time"
//)
//
//func counter() {
//	list := []int{1}
//	c := 1
//	for i := 0; i < 10000000; i++ {
//		httpGet()
//		c = i + 1 + 2 + 3 + 4 - 5
//		list = append(list, c)
//	}
//	fmt.Println(c)
//	fmt.Println(list[0])
//}
//
//func work(wg *sync.WaitGroup) {
//	for {
//		counter()
//		time.Sleep(1 * time.Second)
//	}
//	wg.Done()
//}
//
//func httpGet() int {
//	queue := []string{"start..."}
//	resp, err := http.Get("http://www.163.com")
//	if err != nil {
//		// handle error
//	}
//
//	//defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		// handle error
//	}
//	queue = append(queue, string(body))
//	return len(queue)
//}
//
//func main() {
//	flag.Parse()
//
//	//这里实现了远程获取pprof数据的接口
//	go func() {
//		log.Println(http.ListenAndServe("localhost:7777", nil))
//	}()
//
//	var wg sync.WaitGroup
//	wg.Add(10)
//	for i := 0; i < 100; i++ {
//		go work(&wg)
//	}
//
//	wg.Wait()
//	time.Sleep(3 * time.Second)
//}
import (
	"bufio"
	"fmt"
	"github.com/google/gops/agent"
	"ht.com/go-my/test/reflect-example"
	"ht.com/go-my/test/sync-example"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

const (
	i = 1 << iota
	k
	j = iota
)

func main() {
	fmt.Println("i=", i)
	fmt.Println("k=", k)
	fmt.Println("j=", j)
	reflect_example.AddressableExample()

	fmt.Printf("%d", os.Getpid())
	if err := agent.Listen(agent.Options{ShutdownCleanup: true}); err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < 100; i++ {
		fibonacci(45)
		time.Sleep(10 * time.Second)
	}
	select {}
}

func testSignal() {
	sync_example.Example()
}

func abcd() {
	fmt.Printf("|%#10.v|", "asdfs")
	fmt.Printf("|%#10|", "asdfs")
	var a, b string
	fmt.Scanln(&a, &b)
	fmt.Printf("Hi %s %s!\n", a, b)
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err4 := inputReader.ReadString('\n')
	//func (b *Reader) ReadString(delim byte) (line string, err error) ,‘S’ 这个例子里使用S表示结束符，也可以用其它，如'\n'
	if err4 == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}

func testPprof() {
	f, err := os.Create("cpu-profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	fmt.Println(fibonacci(45))
	pprof.StopCPUProfile()
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
