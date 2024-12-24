#  Functions, Parameters and Defer

functions are "first calss" objects; you can:
    - define them - even inside function
    - create anony ous function literals
    - Pass them as function parameters / return values
    - Store in variables
    - Store them in slices and maps (but not as keys).
    - Store them as fileds of a struture type.
    - Send and recieve them in channels
    - Write methjod against a function tyope
    - Compare a function var against nil.


# Function Scope
Almost anythin can be defined inside a function


func Do() error {
    const a = 21

    type b struct {
        ...
    }

    var c int

    func reallyDoIt() {
        ...
    }
}

Methods cannot be defined in a function (only at package scope.)


The signatire of a function is the order & type of its parameters and return values.

It does not depends on the names of those parameters or returns

var try func(string, int) string

func Do(a string, b int) string {

}

func notDo(x string, y int) (a string) {}
    ...
}


These functions have the same structural type.

# Parameter Terms

a function declaration lists formal parameters

func do(a, b int) int { ... }

a function call has actual parameters (a/k/a "arguments")

return := do(1, 2)

A parameter is passed "by value" if the function gets a copy
the caller cannot see the changes to copy

A parameter is passed "by refrence" if the function can modify the actual paremters
such that caller can see the changes

By Value:
     - numbers
     - bool
     - arrays
     - structs

By Refrence:
     - things passed by pointer (&x)
     - strings (but they are immutable) 
     - slices
     - maps
     - channel


**Parameters are passed by value or by refrence

Actually, all parameters are passed by copying something (i.e, by value)

If the thing copied is a pointer or descriptor, then the shared
backing store (arry, hash table, etc.) can be changed through it

Thus we think of its as "by refrence".


# Return values

Functions can have mul;tiple return values (single return can be direct but mutiple return definition should have paratheses around)

Every return statement must have all values specified.

# Defer

How do we make sure someting gets done ?
     - close a file we opened
     - close socket / HTTP request we made.
     - unlock a mutex we locked
     - make sure something was saved before we're done

Defer statement captures a function call to run later.


e.g

func main() {

    f, err := os.Open(fname)

    if err != nil {
        ...
    }
    defer f.Close()
    ...
}


The call to Close is guranted to run at function exit
(don't defer the closing file until we know it really opened.)

*multiple defer will exit in LIFO order
*defer works at function scope i.e even if defer is called at a sub-cope,
 i.e for-loop block or if block the defer will be closed at function exit.

# e.g Defer Gotcha #1

func main() {
    for i := 1; i < len(os.Args); i++ {
        f, err := os.Open(os.Args[i])
        ...
        defer f.Close() // use f.Close() instead els the all files will remain open until the main function exists.
        ...
    }
}

The deffered calls to Close must wait until the function exit
(we might run out for file descriptors before that!)

# e.g Defer Gotcha #2

A defer statement run before return is "done"

func do() (a int){
    defer func() {
        a = 2
    }()

    a = 1

    return 
} // returns 2

We have named return value and a "naked" return
The deferred anonymous function can update that variable.



https://youtu.be/wj0hUjRHkPs?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6