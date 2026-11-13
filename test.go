// test.go
//
// Go 语言学习实验室。
//
// 运行方式：
//
//	go run test.go
//
// 输入 1-8 选择一节示例，输入 9 依次运行全部示例，输入 q 退出。
// 想改完立即看效果：运行 dev.bat（热重载），保存后会自动重编译、
// 重跑上次选择的小节，并高亮显示改动的源码行。
// 建议先运行一遍，再修改源码中的数字和打印语句，观察输出变化。
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func main() {
	autoDemo := flag.Int("demo", 0, "启动后先运行指定小节（1-8，9=全部）")
	flag.Parse()

	fmt.Println()
	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("  Go 语言学习实验室 · test.go")
	fmt.Println("════════════════════════════════════════════════════════════")
	fmt.Println("  先读代码，再运行；改一改，再运行。")

	if *autoDemo >= 1 && *autoDemo <= 9 {
		fmt.Println()
		fmt.Println("━━━━━━━━━━ 自动运行（上次选择） ━━━━━━━━━━")
		fmt.Printf("  第 %d 节：%s\n", *autoDemo, demoName(*autoDemo))
		runDemo(*autoDemo)
		fmt.Println()
		fmt.Println("━━━━━━━━━━ 自动运行结束 ━━━━━━━━━━")
		fmt.Println("  下面是菜单，输入编号继续；输入 q 退出。")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		printMenu()
		fmt.Print("请选择> ")
		if !scanner.Scan() {
			fmt.Println()
			return
		}

		choice := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if choice == "q" || choice == "quit" || choice == "exit" {
			fmt.Println("再见，继续加油！")
			return
		}

		n, err := strconv.Atoi(choice)
		if err != nil || n < 1 || n > 9 {
			fmt.Println("无效输入，请输入 1-9 或 q。")
			continue
		}

		rememberChoice(n)
		runDemo(n)
	}
}

func demoName(n int) string {
	switch n {
	case 1:
		return "变量、常量与基本类型"
	case 2:
		return "条件、循环与 switch"
	case 3:
		return "函数、多返回值、闭包与 defer"
	case 4:
		return "切片、map 与 range"
	case 5:
		return "结构体、方法与指针"
	case 6:
		return "接口、错误处理与 errors.Is/As"
	case 7:
		return "泛型"
	case 8:
		return "goroutine、channel 与 WaitGroup"
	case 9:
		return "全部示例"
	default:
		return "未知小节"
	}
}

// runDemo 集中处理菜单编号，供手动选择和热重载自动运行共用。
func runDemo(n int) {
	switch n {
	case 1:
		demoBasics()
	case 2:
		demoControlFlow()
	case 3:
		demoFunctions()
	case 4:
		demoCollections()
	case 5:
		demoStructs()
	case 6:
		demoInterfacesAndErrors()
	case 7:
		demoGenerics()
	case 8:
		demoConcurrency()
	case 9:
		runAll()
	}
}

// rememberChoice 记住最后一次选择，热重载重启后由 watcher 读取。
func rememberChoice(n int) {
	_ = os.WriteFile(".last_demo", []byte(strconv.Itoa(n)), 0o644)
}

func printMenu() {
	fmt.Println()
	fmt.Println("──────────────────────── 选择学习内容 ────────────────────────")
	fmt.Println("  1. 变量、常量与基本类型")
	fmt.Println("  2. 条件、循环与 switch")
	fmt.Println("  3. 函数、多返回值、闭包与 defer")
	fmt.Println("  4. 切片、map 与 range")
	fmt.Println("  5. 结构体、方法与指针")
	fmt.Println("  6. 接口、错误处理与 errors.Is/As")
	fmt.Println("  7. 泛型")
	fmt.Println("  8. goroutine、channel 与 WaitGroup")
	fmt.Println("  9. 运行全部示例")
	fmt.Println("  q. 退出")
}

// section 输出统一的小节标题，让每一节的开始更醒目。
func section(n int, title string) {
	fmt.Println()
	fmt.Println(strings.Repeat("─", 60))
	fmt.Printf("  %d. %s\n", n, title)
	fmt.Println(strings.Repeat("─", 60))
}

func runAll() {
	demoBasics()
	demoControlFlow()
	demoFunctions()
	demoCollections()
	demoStructs()
	demoInterfacesAndErrors()
	demoGenerics()
	demoConcurrency()
}

// ---------------------------------------------------------------------------
// 1. 变量、常量与基本类型
// ---------------------------------------------------------------------------

func demoBasics() {
	section(1, "变量、常量与基本类型")

	var name string = "Gopher"
	age := 8
	height := 1.75
	isReady := true
	const pi = 3.14159

	fmt.Println()
	fmt.Println("  变量值：")
	fmt.Printf("    %-10s = %-14v (%T)\n", "name", name, name)
	fmt.Printf("    %-10s = %-14v (%T)\n", "age", age, age)
	fmt.Printf("    %-10s = %-14v (%T)\n", "height", height, height)
	fmt.Printf("    %-10s = %-14v (%T)\n", "ready", isReady, isReady)

	fmt.Println()
	fmt.Println("  常量：")
	fmt.Printf("    pi = %.5f\n", pi)

	fmt.Println()
	fmt.Println("  表达式：")
	fmt.Printf("    age + 1    = %d\n", age+1)
	fmt.Printf("    height * 2 = %.1f\n", height*2)

	// 只声明、不赋值时，变量会得到该类型的“零值”。
	var zeroInt int
	var zeroString string
	var zeroBool bool

	fmt.Println()
	fmt.Println("  零值（只声明、不赋值）：")
	fmt.Printf("    %-10s = %-14v (%T)\n", "int", zeroInt, zeroInt)
	fmt.Printf("    %-10s = %-14q (%T)\n", "string", zeroString, zeroString)
	fmt.Printf("    %-10s = %-14v (%T)\n", "bool", zeroBool, zeroBool)

	fmt.Println()
	fmt.Printf("  提示：%%T 打印类型，%%v 打印默认格式，%%.2f 保留两位小数。\n")

	// 练习：
	//  1. 增加变量 city，并用 %s 打印。
	//  2. 把 age 改成 18，观察表格里的变化。
	//  3. 试试把 height 改成 int，看编译器报什么错。
}

// ---------------------------------------------------------------------------
// 2. 条件、循环与 switch
// ---------------------------------------------------------------------------

func demoControlFlow() {
	section(2, "条件、循环与 switch")

	// FizzBuzz：3 的倍数输出 Fizz，5 的倍数输出 Buzz，15 的倍数输出 FizzBuzz。
	for i := 1; i <= 15; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("FizzBuzz ")
		case i%3 == 0:
			fmt.Print("Fizz ")
		case i%5 == 0:
			fmt.Print("Buzz ")
		default:
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	score := 88
	switch {
	case score >= 90:
		fmt.Println("等级：A")
	case score >= 80:
		fmt.Println("等级：B")
	default:
		fmt.Println("等级：C，继续努力")
	}

	// switch 也可以先执行一条初始化语句。
	switch day := 3; day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("周末")
	}

	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println("0..4 之和 =", sum)

	n := 3
	for n > 0 {
		fmt.Printf("倒计时 %d\n", n)
		n--
	}

	words := []string{"go", "is", "fun"}
	for i, word := range words {
		fmt.Printf("words[%d] = %s\n", i, word)
	}

	// 练习：
	//  1. 打印 1..100 中所有能被 7 整除的数。
	//  2. 用 switch 写一个函数，把 1-7 映射成“星期一”到“星期日”。
}

// ---------------------------------------------------------------------------
// 3. 函数、多返回值、闭包与 defer
// ---------------------------------------------------------------------------

func add(a, b int) int {
	return a + b
}

// Go 的函数可以返回多个值。
func divmod(a, b int) (int, int) {
	return a / b, a % b
}

// nums ...int 是可变参数，调用时可以传任意多个 int。
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// counter 返回一个闭包：它“记住”了 count 变量。
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func withDefer() {
	defer fmt.Println("defer：函数返回前执行")
	fmt.Println("函数体：正常执行")
}

func demoFunctions() {
	section(3, "函数、多返回值、闭包与 defer")

	fmt.Println("add(2, 3) =", add(2, 3))

	q, r := divmod(17, 5)
	fmt.Printf("divmod(17, 5) = 商 %d，余数 %d\n", q, r)

	fmt.Println("sum(1, 2, 3, 4) =", sum(1, 2, 3, 4))

	next := counter()
	fmt.Println("闭包计数：", next(), next(), next())

	withDefer()

	// 练习：
	//  1. 写一个 max(a, b int) int。
	//  2. 写一个 greet(name string) string，返回 "Hello, <name>"。
	//  3. 让 divmod 在除数为 0 时返回错误（需要改成三个返回值）。
}

// ---------------------------------------------------------------------------
// 4. 切片、map 与 range
// ---------------------------------------------------------------------------

func demoCollections() {
	section(4, "切片、map 与 range")

	numbers := []int{5, 2, 8, 1}
	numbers = append(numbers, 9)
	fmt.Printf("numbers=%v len=%d cap=%d\n", numbers, len(numbers), cap(numbers))
	fmt.Println("numbers[1:4] =", numbers[1:4])

	slices.Sort(numbers)
	fmt.Println("排序后 =", numbers)

	total := 0
	for _, n := range numbers {
		total += n
	}
	fmt.Println("总和 =", total)

	scores := map[string]int{
		"小明": 90,
		"小红": 85,
	}
	scores["小刚"] = 78
	scores["小明"] = 95
	delete(scores, "小红")

	// map 的遍历顺序是随机的，所以先取出 key 排序，保证输出稳定。
	names := make([]string, 0, len(scores))
	for name := range scores {
		names = append(names, name)
	}
	slices.Sort(names)
	for _, name := range names {
		fmt.Printf("%s: %d 分\n", name, scores[name])
	}

	// 查 map 时用“值, ok”形式，可以区分“不存在”和“值为零值”。
	if score, ok := scores["小红"]; ok {
		fmt.Println("小红还在：", score)
	} else {
		fmt.Println("小红已被删除，查 map 时 ok =", ok)
	}

	// 注意：向 nil map 写入会 panic。
	// var m map[string]int
	// m["x"] = 1 // panic: assignment to entry in nil map

	// 练习：
	//  1. 统计一段文字中每个单词出现的次数。
	//  2. 删除切片中所有偶数，返回新切片。
}

// ---------------------------------------------------------------------------
// 5. 结构体、方法与指针
// ---------------------------------------------------------------------------

type Student struct {
	Name  string
	Score int
}

// 指针接收者：方法内修改的是原对象。
func (s *Student) AddScore(delta int) {
	s.Score += delta
}

// 值接收者：方法拿到的是副本，适合只读操作。
func (s Student) Passed() bool {
	return s.Score >= 60
}

func demoStructs() {
	section(5, "结构体、方法与指针")

	s := Student{Name: "小明", Score: 80}
	fmt.Printf("创建后：%+v，及格：%v\n", s, s.Passed())

	s.AddScore(8)
	fmt.Printf("AddScore(8) 后：%+v，及格：%v\n", s, s.Passed())

	p := &s
	p.Score += 2
	fmt.Printf("通过指针修改后：%+v\n", *p)

	// 结构体赋值是值拷贝，副本的修改不会影响原值。
	copyOfS := s
	copyOfS.Score = 0
	fmt.Printf("副本修改后：副本=%+v，原值=%+v\n", copyOfS, s)

	// 练习：
	//  1. 增加 Age 字段，并写一个 IsAdult() bool 方法。
	//  2. 把 AddScore 改成值接收者，观察 s 是否还会变化。
	//  3. 给 Student 实现 String() string，让 fmt.Println(s) 更好看。
}

// ---------------------------------------------------------------------------
// 6. 接口、错误处理与 errors.Is/As
// ---------------------------------------------------------------------------

type Shape interface {
	Name() string
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Name() string       { return "矩形" }
func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Circle struct {
	Radius float64
}

func (c Circle) Name() string       { return "圆形" }
func (c Circle) Area() float64      { return 3.14159 * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * 3.14159 * c.Radius }

// 哨兵错误：调用方可以用 errors.Is 判断。
var ErrTooYoung = errors.New("年龄不足 18 岁")

// 自定义错误：调用方可以用 errors.As 取出具体字段。
type ValidationError struct {
	Field string
	Value int
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("字段 %s 的值 %d 不合法", e.Field, e.Value)
}

func register(name string, age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Value: age}
	}
	if age < 18 {
		return fmt.Errorf("注册 %s 失败: %w", name, ErrTooYoung)
	}
	return nil
}

func demoInterfacesAndErrors() {
	section(6, "接口、错误处理与 errors.Is/As")

	// 只要实现了 Shape 的三个方法，就自动满足 Shape 接口，不需要显式声明。
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 2},
	}
	for _, shape := range shapes {
		fmt.Printf("%s：面积 %.2f，周长 %.2f\n",
			shape.Name(), shape.Area(), shape.Perimeter())
	}

	if err := register("小明", -1); err != nil {
		fmt.Println("注册失败：", err)
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("  errors.As 拿到字段：%s=%d\n", ve.Field, ve.Value)
		}
	}

	err := register("小红", 16)
	if err != nil {
		fmt.Println("注册失败：", err)
		if errors.Is(err, ErrTooYoung) {
			fmt.Println("  errors.Is 判断：确实是年龄不足")
		}
	}

	// 练习：
	//  1. 给 Shape 增加 String() string，并让两个图形都实现它。
	//  2. 写一个函数，接收任意 error，打印它的完整错误链。
}

// ---------------------------------------------------------------------------
// 7. 泛型
// ---------------------------------------------------------------------------

// Map 把 []T 的每个元素转换成 U，返回 []U。
func Map[T, U any](items []T, fn func(T) U) []U {
	result := make([]U, 0, len(items))
	for _, item := range items {
		result = append(result, fn(item))
	}
	return result
}

// Filter 保留满足条件的元素。
func Filter[T any](items []T, keep func(T) bool) []T {
	result := make([]T, 0, len(items))
	for _, item := range items {
		if keep(item) {
			result = append(result, item)
		}
	}
	return result
}

// Sum 支持 int 和 float64 两种数值类型。
func Sum[T int | float64](items []T) T {
	var total T
	for _, item := range items {
		total += item
	}
	return total
}

func demoGenerics() {
	section(7, "泛型")

	numbers := []int{1, 2, 3, 4, 5}

	doubled := Map(numbers, func(n int) int { return n * 2 })
	fmt.Println("Map 翻倍：", doubled)

	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Println("Filter 偶数：", evens)

	fmt.Println("Sum int：", Sum(numbers))
	fmt.Println("Sum float64：", Sum([]float64{1.5, 2.5, 3}))

	squares := Map(numbers, func(n int) string {
		return strconv.Itoa(n * n)
	})
	fmt.Println("Map 转字符串：", squares)

	// 练习：
	//  1. 写一个 Reduce。
	//  2. 写一个 Contains[T comparable](items []T, target T) bool。
}

// ---------------------------------------------------------------------------
// 8. goroutine、channel 与 WaitGroup
// ---------------------------------------------------------------------------

func demoConcurrency() {
	section(8, "goroutine、channel 与 WaitGroup")

	var wg sync.WaitGroup
	results := make(chan string)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		// 把 i 作为参数传进去，避免多个 goroutine 共享同一个循环变量。
		go func(id int) {
			defer wg.Done()
			results <- fmt.Sprintf("worker %d 完成任务", id)
		}(i)
	}

	// 等待所有 worker 结束后关闭 channel，range 才能正常结束。
	go func() {
		wg.Wait()
		close(results)
	}()

	for msg := range results {
		fmt.Println(msg)
	}

	fmt.Println("所有 worker 已结束")

	// 练习：
	//  1. 把 results 改成带缓冲的 channel，观察行为。
	//  2. 启动 5 个 worker，计算 1..100 的平方和。
	//  3. 想一想：如果忘记 close(results)，会发生什么？
}
