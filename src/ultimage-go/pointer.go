package main

/*
Pointer - value that contains the address of a variable. It is the location at which a value is stored. Not every value has an address, but every variable does.

「
var x int
&x - address of x - yields a pointer to integer variable, that is a value of type *int
*int - pointer to int
p := &x
fmt.Println(*p) // “1”
*p = 2 // equivalent to x = 2
fmt.Println(x) // “2”
」

& operator yields the address of a variable, and * operator retrieves the variable that the pointer refers to.
Zero value of a pointer is nil.
Pointer serves only 1 purpose: sharing. Pointer shares values across the program boundary (ex.: function calls).

*/



// Pointer serves 1 purpose: sharing
// Pointer shares values across the program boundary (ex.: function calls)

// When program starts, the runtime creates a Goroutine.
// This program has only one goroutine - main goroutine.

// Every Goroutine is given a block of memory - stack.
// The block starts at 2K, very small, and can change over time.
// Everytime a function is called, a piece of stack is used to help that function run.
// Stack grows downward.

// Every function is given a stack frame, memory execution of a function.
// The size of every stack frame is known at compile time. No value can be
// placed on a stack unless the compiler knows its size ahead of time.
// If we don't know the size of something at compile time, it has to be on the heap.

// Zero value enables us to initialize every stack frame that we take.
// Stack are self cleaning. We clean our stack on the way down.

type user struct {
	name 	string
	email string
}

func main() {
	// --------------
	// Pass by value
	// --------------

	// This value is put on a stack with a value of 10.
	count := 10

	// To get the address of a value, we use &
	println("count:\tValue of[", count, "]\tAddr of[", &count, "]")

	// Pass the "value of" count
	increment1(count)

	// Print out the value of count - nothing has changed.
	println("count:\tValue of[", count, "]\tAddr of[", &count, "]")

	// Pass addres of the count variable
	// This is still pass by value, not by reference because the address itself is a value
	increment2(&count)

	// This time count is updated.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	stayOnStack()
}

func increment1(value int) {
	value++
	println("inc1:\tValue of[", value, "]\tAddr of[", &value, "]")
}

// pointer - address to a variable
// * here is a part of the type name *int, read: pointer to int
func increment2(pointer *int) {
	// Increment the "value of" count that the "pointer" points to.
	*pointer++	// Gets the value of the variable it points to
	println("pointer2:\tValue Of[", pointer, "]\tAddr Of[", &pointer, "]\tValue Points To[", *pointer, "]")	
}


// Variable does not escape stack, it is passed back to the caller's stack
func stayOnStack() user {
	// In the stayOnStack stack frame, create a value and initialize it.
	u := user{
		name: "Chi",
		email: "chi@here.com",
	}
	// Take the value and return it, pass back to main stack frame
	return u
}


// in stayOnStack function, we are passing the copy of the value itself, it is safe
// to keep these things on the stack. Bu when we SHARE something above the call stack
// like this, escape analysis said this memory is no longer valid when we get back
// to main, we must put it on the heap. main will end up having a pointer to the heap.
func escapeToHeap() *user {
	u := user{
		name: "Chi",
		email: "chi@here.com",
	}
	return &u
}

// ----------------------------------
// What if we run out of stack space?
// ----------------------------------

// What happen next is during that function call, there is a little preamble that asks "Do we have
// enough stack space? for this frame?". If yes then no problem because at complied time we know
// the size of every frame. If no, we have to have bigger frame and these values need to be copy
// over. The memory on that stack move. It is a trade off. We have to take the cost of this copy
// because it doesn't happen a lot. The benefit of using less memory any Goroutine outrace the
// cost.

// Because stack can grow, no Goroutine can have a pointer to some other Goroutine stack.
// There would be too much overhead for complier to keep track of every pointer. The latency will
// be insane.
// -> The stack for a Goroutine is only for that Goroutine only. It cannot be shared between
// Goroutine.

// ------------------
// Garbage collection
// ------------------

// Once something is moved to the heap, Garbage Collection has to get in.
// The most important thing about the Garbage Collector (GC) is the pacing algorithm.
// It determines the frequency/pace that the GC has to run in order to maintain the smallest t as
// possible.

// Image a program where you have a 4 MB heap. GC is trying to maintain a live heap of 2 MB.
// If the live heap grow pass 4 MB we have allocate a larger heap.
// Depending how fast the heap grow, we determine the pace that the GC has to run. The smaller the
// pace, the less impact it is going to have. The goal is to get the live heap back down.

// When the GC is running, we have to take a performance cost so all Goroutine can keep running
// concurrently. The GC also have a group of Goroutine that perform the garbage collection work.
// It takes 25% of our available CPU capacity to itself.
// More details about GC and pacing algorithm can be find at:
// https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/README.md