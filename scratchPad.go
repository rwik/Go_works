package main

import (
	"fmt"
	"strconv"
	"time"
)

//init function executes before main
func init() {
	fmt.Println("Inside init")
}
func init() {
	fmt.Println("There can be more than one init")
}

//struct exmple
type superHero struct {
	name   string
	power  string
	energy int
}
type emp struct {
	Name    string // exported as capitalized
	sallery float32
	int     // anonymous field. can be accessed and assigned value
}

func main() {
	fmt.Println("it works")
	// multiple variable declaration with initial value
	var a1, b1 int = 45, 54
	fmt.Println(a1, b1)
	// decalre inside block
	var (
		aa int
		bb float32 = 34.6
		cc string  = "rwik"
	)
	fmt.Printf("%v , %v , %v", aa, bb, cc)
	//Unused vars will be reported as compiler error
	var x = 34.6 // type inference float64
	fmt.Printf("%v ", x)

	//local vs global var
	i := 10
	for i := 0; i < 3; i++ {
		fmt.Println("local ", i)
	}
	fmt.Println("global ", i)
	//constan declaration
	const c = "cyan"
	const (
		c1 = 12
		c2 = "dsfsad"
	)
	// const vars cant be reassigned and can remain unused
	// const declared in inner scope will always get more
	// priority than const declared outside
	// go doesn't support const struct map array or slice
	// There are no ternary operator in go, use if else
	if c == "cyan" {
		fmt.Println("cyan it is")
	} else if c == "blue" {
		fmt.Println("blue it is")
	} else {
		fmt.Println("It is something else")
	}
	// go doesn't need any break in switchcase
	// it exits after 1 condition matches
	// use fallthrough to go through all cases
	switch {
	case c == "cyan":
		fmt.Println("inside first case")
		fallthrough
	case c != "blue":
		fmt.Println("inside second case")
	default:
		fmt.Println("this is default case")
	}
	//every go source file is part of a package
	//main is executable package
	//all other packages are called utility package
	//$GOPATH/src contains all suuporting modules
	// go get gthub.com/package
	// will fetch package and save to src folder
	//go.mod (import path,version,dependency)
	//go.sum (checksum for dependencies)

	//Map iteration
	sample := map[string]string{
		"k1": "val1",
		"k2": "val2",
	}
	for k, v := range sample {
		fmt.Println(k + " " + v)
	}
	//get all keys in a array
	var keys []string
	for k := range sample {
		keys = append(keys, k)
	}
	fmt.Print(keys)
	//add to map
	sample["k3"] = "v3"
	//delete a key
	delete(sample, "k3")
	// an empty map will have nil value
	//Maps are not safe for concurrent use: it's not defined what happens when you read and write to them simultaneously. If you need to read from and write to a map from concurrently executing goroutines, the accesses must be mediated by some kind of synchronization mechanism. (source: https://blog.golang.org/go-maps-in-action)

	//check if key exists
	val, ok := sample["rodi"]
	fmt.Println("Key doesn't exist case "+val, ok)

	//struct example
	s1 := superHero{
		name:   "Superman",
		energy: 99,
	}
	s2 := superHero{"batman", "can talk in heavy tone", 78}
	fmt.Println(s1)
	fmt.Println(s2)

	//create a channel
	ch := make(chan int, 1) // create a buffered channel length 1
	ch <- 12
	ch_val := <-ch // you can receive only once
	//ch_val <- ch//error
	fmt.Println(len(ch), cap(ch)) // length and capacity of channel
	//for unbuffered channel, length and capcity is always zero
	fmt.Println("Data received from chanel ", ch_val)
	close(ch)

	// channels are two types
	//buffered : send works until buffer is full
	//unbuffered : send is blocked until another go routine is there to receive

	//go rotine is lightweight thread
	fmt.Println("before go routine call")
	go rtCall()
	fmt.Println("After go routine call")
	time.Sleep(time.Second * 1)
	// advantages of go routine over thread
	//goroutine starts with 8kb,os threads can be more than 1mb
	//scheduling is done by go runtime, hence it is faster. go runtime considers virtual cpus. 100s of routines can be managed by 2 os threads.
	//communication between routines is done by channel

	//pointer declaration
	aPointer := new(int)
	*aPointer = 989
	fmt.Println(aPointer, *aPointer) // prints address and value
	// no pointer arithmatic in go
	bPointer := &aPointer
	fmt.Println(bPointer, *bPointer, **bPointer) // pointer to pointer

	//String to int
	var str1 string
	str1 = "1234"
	valstr1, err := strconv.Atoi(str1)
	if err == nil {
		fmt.Println("String str1 is a number", valstr1)
	}
}

func rtCall() {
	fmt.Println("inside go routine")
}
