# Interfaces and Methods in detail

- Details, Details

# Nil Interfaces

    An interface variable is nil until Initialized

    It really has two parts:
        - a value or pointer of some type.
        - a pointer to type information so the corrcet actual method can be identified


        var r io.Reader        // nil until initialized
        var b *bytes.Buffer    // nil until initialized

        r = b                  // r is nop longer nil
                               // but it has nil pointer to buffer
    
    This may confuse; an interface is nil only if both parts are

# Error is really an interface

We called error a special type, but it's really and interface

    type error Interface{
        Error() string
    }

We can compare it to nil unless we make a mistake

The mistake is to store a nil pointer to a concrete type in error variable

# Pointers vs value recievers

- Pointers method may be called on non-pointers and vice-versa

 - Go will automatrically use * or & as needed

    p1 := new(Point)        // *Point, at 0,0
    p2 := Pouint{1,2}

    p2.OffsetOf(p2)         // same as (*p1).OffsetOf(p2)

    p2.Add(3, 4)            // same as (& p2).Add(3, 4)


Except & may be only appied to object that are addresable

- Compatability between objects and reciever types

                                 | Pointer  | L-Value   |  R-value  |
                pointer reciever |   OK     |   OK &    |   Not OK  |
                value   reciever |   OK *   |   OK      |    OK     |

- A method requiring a pointer reciver may only be called on addressabe object

    var p Point

    p.Add(1, 2)             // OK, &p
    Point{1, 2}.Add(3, 4)   // Not OK, can't take address of Point literal.


# Consistency in reciever types

- If One method of type takes a pointer reciever, then all
  its methods should take pointers(except when they shouldn't for other reasons)


And in gerebal objects of tyoe are probably not safe to copy

    type Buffer struct {
        buf []byte
        off int
    }

    func (b *Buffer) ReadString(delim byte) (string, error) {
        ...
    }


# Currying functions

- Currying takes a functions and reduces its argument count by one (one argument gets bound, and new function is returned.)

    func Add(a, b int) int {
        return a+b
    }

    func AddtoA(a int) func(int) int {
        return func(b int) int {
            return Add(a, b)
        }
    }

    addTo1 := AddToA(1)

    fmt.Println(Add(1, 2) == addTo1(2)) // true

# Method values

- A selected method may be passed similar to a closur;
  the reciever is colsed over at that point


    func (p Point) Distance(q Point) float64 {
        return math.Hypot(q.X-p.X, q.Y - p.Y)
    }


    p := Point{1, 2}
    q := Point{4, 6}

    distanceFromP := p.Distance       // This is a mthod value 

    fmt.Println(distanceFromP(q))     // can be called later


# Refrences and value semantics

A method value with a value reciever copies the reciever
If It has a pointer reciever, it copies a pointer to the reciever.


    func (p *Point) Distance(q Point) float64 {
        return math.Hypot(q.X-p.X, q.Y - p.Y)
    }

    p := Point{1, 1}
    q := Point{4, 6}

    distanceFromP := p.Distance
    p := Point{3, 4}

    fmt.Println(distanceFromP(q))     // takes new value of p


# Intefaces in practice
    - Lets consumers define interfaces
        (what minimal behavior do they require)
    
    - Re-use standard intrfaces whenever possoble
    
    - Keep Interfaces declrations small
        ("The bigger the Interface, the weaker the abstraction")
    
    - Compose one-method interfaces into larger interfaces
    
    - Avoid xoupling interfaces to particular types/implementations.
    
    - Accepet interfaces, but return concrete types
        (let consumers of the return type decide how to use it)

# Interfaces vs concrete values
    "Be libeal in what you accept, be conservative in what you return"

    - Put the least restriction on what parametrs you accept
        (minimal interface)
        
        Don't require ReadWriteCloser if you only need to read.
    
    - Avoid restricting the use of your return type (the concrete value you return
      might fit into many interfaces!)

      Returning *os.File is less restrictive than eturning i0.ReadWriteClose
      becaue files have other useful methods.
    
    Returning error is a good example of an exception to this rule.

# Empty Interfaces

    The intereface{} type has no methods, so it is satisfied by anything!

    Empty interfaces are commnly used; they're how the formatted I/O routines
    can print any type

        func fmt.Printf(f string, args ...interface{})

    Reflection is needed to determine what concrete type is



https://youtu.be/AXCIEiebVfI?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6