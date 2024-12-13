basic types in GO


- go has 25 keywords which cannot be used as names


Constants in go
 - true, false, iota, nil

Types:
 - int (8,16,32,64)
 - unint (8, 16, 32, 64), uintptr
 - float32, float64, comploex64, complex128
 - bool, byte, rune, string, error

Functions
    - make, len, cap, new, append, copy, close, delete
      complex, real , imag
      panic, recover


machine-native anad interpreted
-------------------------------

a := 2 -> memory location at RAM. (go)

a = 2 -> an Object a is created, which interpreter converts and places into memory.


- `int` is the default type for integers in Go.
- non-integeers are represnted in floating point.
    float32, float64(default)
- complex (imaginary) floating point numbers:
    complex64, complex128


*Don't use floating point for monetary calculations! (Try package like Go money)


Declarations

- var a int

- var (
    b = 2
    f = 2.01
)


- c := 1 [Only inside functions]


% in in formatted string indicates verb

%T - Type of variable
%v - Value of variable


Both the statements below outputs the same reuslt as %[n] indicates to re-use the n positioned variable for the verb.
and Indexing starts from 1 for the arguments in formatted statement(printf)
- fmt.Printf("a: %8T %v \n", a, a)
- fmt.Printf("a: %8T %[1]v \n", a)


- Special Types

bool (true false)
    - *these values are not convertible to/from integers

error (a special type with one function, Error())
    - an error may be nil or non-nil

pointers (physicall address, logically opaque)
    - a pointer may be nil or non-nil
    - `no pointer maniupulation` except through package unsafe 


--> Go initliazes the variables to "zero" by default if we don't
- numericals - 0
- bool - false
- strings - "" (empty string, length 0)
- everythin else - nil

Constants (absolutely Immutable)

const (
    a = 1              // int
    b = 2 * 1024       // 2048
    c = b << 3         // 16384

    g unit8 = 0x07     // 7
    h unit8 = g & 0x03 //3

    s = "a string"
    t = len(s)         // 8
    u = s[2:]          // SYNTAX ERROR
)


https://youtu.be/NNLpEPb2ddE?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6