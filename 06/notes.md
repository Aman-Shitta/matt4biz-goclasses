Control statements; Declarations & Types.

Control structure

if-then-else

All if-then statements require braces

if a == b {
    ....
} else {
    ...
}

if err := doSomething(); err != nil {
    ....
}


For Loops

There is only for (no do or while) but with options


for i := 0; i < 10; i++ {
    fmt.Printf("%d, %d\n", i, i*i)
}

// prints (0,0) upto (9, 81)

Three parts, all optinal (initialize, check, incement)

The loop ends when the explicit checks fails (e.g., 1==10)



Imoplicit control through the "range" operator for
array and slices

// var i is an index 0, 1, 2, 3...
for i := range myArray {
    ....
}

// var i is index, v is value
for i, v := range myArray {
    ....
}

The for loop ends when the range is exhausted.


In case of map with one variable the iteration is on keys
with two variables keys and value are iterated.


An infinite loop 

for {
    ...

    if certainCOndition {
        break
    } else {
        // move to next iteration using continue
        continue

    }
}


_ (undersocre) is a blank identifier which is untypes, resuable "variable" placeholder



sometimes we need to break or continue the ouer loop
(nested loop for quadratic search)

outer:
    for k := range arr1 {
        for _, v := range arr2 {
            if k == v.ID {
                continue outer
            }
        }
    }


Switch

switch a := f.Get(); a {
    case 0,1,2:
        ...
    case 3, 4, 5: // intended skip of statement allow

    default:
        ....
}


in swicth statements in go the sattement do not fall through i.e break is not required.

The scope of a variable is overshadwoed by the same varible name defined using short declaration

func foo() {

    var err error

    for n, err := doSome() { // the err declared here is in scope of for loop and oveshadows outer err inside the block
        if err !=nil {
            break
        }
        doSomethingElse()
    }

    return err // ALWAYS nil, outer scope erro will be returned always
}

Structural Typing
 - Its the same type if it has the same structure and base type
    - arrays of same size and base type.
    - slices with the same base types.
    - maps of the same key and value types.
    - structs with sthe same sequence of field name/types.
    - functions with same parameters and return types.

Named Typing

It's only the same tyype is it has same declared type name

type x int

func main() {

    var a x         // types x, base int

    b := 10         // int

    a = b           // TYPE MISMATCH

    a = 12          // OK, untyped literal

    a = x(b)        // OK, type conversion
}

Go uses named typing  for non-function user-delacred types.

https://youtu.be/qpHLhmoV3BY?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6