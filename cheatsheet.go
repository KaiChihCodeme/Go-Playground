// goroutine
go func() {

} ()

go function11()

// channel
intChan := make(chan int, 1)
intChan <- 1

for {
	select {
	case res := <-intChan:
		fmt.Println(res)
	}
}

// traverse multi chan
intMChan := make(chan int, 5)
go func() {
	for i := range 5{
		intMChan <- random.Intn(10)
	}
	close(intMChan)
} ()
time.Sleep(time.Second * 5)
for i := range intMChan {
	fmt.Println(i)
}


// waitgroup
wg := &sync.WaitGroup{}
wg.Add(1)
defer wg.Done()
wg.Wait()

// mutex
lock := sync.Mutex
lock.Lock()
// write
lock.Unlock()

// do not stop by panic
if err := recover(); err != nil {
	// print error
}

// context cancel
ctx, cancel := context.WithCancel(context.Background())
go func(ctx context.Context) {
	for {
		select {
		case ctx.Done():
			fmt.Println("Over")
		}
	}
} (ctx)
time.Sleep(time.Second * 5)
cancel()

// context withTimeout, WithDeadline
ctx, _ := context.WithTimeout(context.Background(), time.Second * 5)
ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))

// context WithValue
ctx := context.WithValue(context.Background(), "key", "val")
key := ctx.Value("key").(string)

// reflection
var f float64
f = 35.2

t1 := reflect.TypeOf(f)
v1 := reflect.ValueOf(f)
k1 := reflect.TypeOf(f).Kind()

// check struct info
type User struct {
	Name string
	Age int
	Married bool
}

u := User {
	Name: "aa",
	Age: 1,
	Married: true,
}
vu := reflect.ValueOf(u)
for i:0; i<vu.NumField(); i++{
	field := i.Field()
	fmt.Println(field.Kind())
	if field.Kind() == reflect.Int {
		fmt.Println(field.Type().Name(), field.Int())
	}
}

// gin 
r := gin.Default()
v1 := r.Group("api/v1")
{
	v1.GET("/path/:id", GetPath)
	v1.POST("/path", PostPath)
}
r.Run(":8080")