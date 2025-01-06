# Methods and Interfaces

Why have methods ?

An interfcae specifies abstract behaviour in terms of methods

    type Stringer interface { // in fmt
        String()    string
    }

"Concrete" types offer methods that satisfy the interface.


# Methods are type-boud function

A method is a special type of function (syntax from Oberon-2)

Ih has a reciever parameter before the function name parameter

    type IntSlice []int

    func (is IntSlice) String() string {
        strs []string

        for _, v := range is {
            strs = append(strs, strconv.Itoa(v))
        }

        return "[" +strings.Join(strs) + "]"
    }


# Why Intrefaces

Without interfaces, we'd have to write (many) functions
for (many) concrete types, possibly coupled to them

    func OutputToFile(f *File, ...) { ... }
    func OutputToBuffer(b *Buffer, ...) { ... }
    func OutputToSocket(s *Socket, ...) { ... }

Better - we want to define our function in terms of abstract behavior

    type Writer interface {
        Write([]bytes) (int, error)
    }

    func OutputTo(w *io.Writer, ...) { ... }

An interface specifies required behavior as a mtheod set

Any type that implements that method set satisfies the inteerface:

    type Stringer interfact {
        String() string
    }

    func (is IntSlice) String() string {
        ...
    }

This is known as structural typing ("duck" typing)

No type will declare itself to implement ReadWrite Explicitly.


*A method may be defined on any user-declared (named) type*

That menas metyhods can't be decalred on int, but
     
    type MyInt int

    func (m MyInt) String() string { ... }

The same method name may be bound to different types

*Some rules and restriction apply, see packege insert for details

# Recievers

A method may take a pointer or value reciever, but not both

    type Point struct {
        X, Y float64
    }

    func (p Point) Offset(x, y flaot64) Point {
        return Point{p.X+x, p.Y+y}
    }

    func (p *Point) Move(x, y flaot64) {
        p.X += x
        p.Y += y
    }

*Taking a pointer allows a method to change the reciever (original object)*


# Interfaces and Substitution

All the methods must be rpesnt to satisfy the interface

    var w io.Writer
    var rwc io.ReadWriteCloser
    
    w = os.Stdout                   // OK: *os.File has Write method 
    w = new(bytes.Buffer)           // OK: *bytes.Buffer has Write Method
    w = time.Second                 // Error: no Write Method

    rwc = os.Stdout                 // OK: *os.File has all three methods
    rwc = new(bytes.Buffer)         // Error: no Close Method

    w = rwc                         // OK: io.ReadWriteCloser has Write
    rwc = w                         // Error: No close method

That's why it pays to keep interfaces small.


*The reciever must be of right type (pointer or value)*

    type IntSet struct { ... }

    func (is *IntSet) String() string { ... }

    var _ = IntSet{}.String()       // Error: String need *Intset(location-value)
    var s IntSet

    var _ = s.Strng()               // OK: s is a variable; &s used automatically

    var _ fmt.Stringer = &s         // OK
    var _ fmt.Stringer = s          // Error: no String Method

We'll come back and talk about pointer vs value recievers in more detail


# Interface Composition

io.ReadWriter is actually defined by Go as two interfaces

    type Reader interface() {
        Read(p []byte) (n int, err error)
    }

    type Writer interface() {
        Write(p []byte) (n int, err error)
    }

    type ReadWriter interface {
        Reader
        Writer
    }

Small interfaces with composition where needed are more felxible

# Interface Declarations

ALl methods for a given type must be deeclared in the same package
where the type is declared.

This allows a package, importing the type, to know all the methods at compile time.

But we can always extend a type in a new package through embedding

    type Bigger struct {
        my.Big          // get all Big methoids via promotion
    }

    func (b Bigger) DoIt() { ... }  // ass one more method here


https://youtu.be/W3ZWbhQF6wg?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6
