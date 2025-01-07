# Composition

Composition, not inhertiance

The fields of an embedded struct are promoted to the
level of the embediing structure

    type Pair struct {
        Path string
        Hash string
    }

    type PairWithLength struct {
        Pair            // promoted, i.e fields of Pair appear at the same level as fields of PairWithLength
        Length int
    }

    p1 := PairWithLength{Pair{"/usr", "0xdeadbeef"}, 121}

    fmt.Println(p1.Path, p1.Length)    // not p1.x.Path


*methods for a promoted type are also promoted.*


# Composition with Pointer Types

A struct can embed a pointer to another type; promotion
of its fields and methods works the same way

    type Fizgig struct {
        *PairWithLength
        Broken bool
    }

    fg := Fizgig{
        &PairWithLength{Pair{"/usr", "0xdead"}, 123},
        false
    }

    fmt.Println(fg)

    // Length of /usr is 123 with hash 0xdead

# Sorting

Sortable Intefaces

- sort.Interface is defined as 

    type Interface intrface {
        // Len is the number of elements in collection
        Len() int

        // Less reports weather the element with
        // index i should sort begore the element with index j.
        Less(i, j int) bool

        // Swap swaps the elements with indexes j and j
        Swap(i, j int)
    }

- and sort.Sort as 

    func Sort(data Inteface)

# Sortable builnins

- Slices of strings can be sorted using StringSlice

    // defined in sort package
    // type StringSlice []string

    entries := []strings{"charlie", "dog", "baker","able",}

    sort.Soort(sort.StringSlice(entries))

    fmt.Println(entries)     // [able baker charlie dog]

 
# Sorting example

- Implement sort.Interface to make a type sortable:

        type Organ struct {
            Name string
            Weight int
        }

        type Organs []Organ

        func (s Organs) Len() int { return len(s)}

        fun (s Organs) Swap(i, j int ) { s[i], s[j] = s[j], s[i]}

# Sorting in reverse

- Use sort.Reverse which is defined as

    type reverse struct {
        Interface
    }

    type (r reverse) Less(i, j int) bool {
        return r.Interface.Less(j, i)
    }

    func Reverse(data Interface) Interface {
        return &reverse{data}
    }

*Make the nil value useful : https://youtu.be/0X6AcnwocbM?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&t=1456*

    type StringStack struct {
        data []string   // data , d is small so not accessbile to outside world 
                        // "zero" value ready to use
    }

    func (s *StrigStack) Push(elem string) {
        s = append(s.data, elem)
    }

    func (s *StringStack) Pop() string {

        if l := len(s.data); l != 0 {
            t := s.data[l-1]
            s.data = s.data[:l-1]
            return t
        }

        panic("cannot pop from empty stack")

    }

# nil as a reciever value

* *NOTHING IN GO PREVENTS CALLING A METHOD WITH A NIL RECIEVER* *

    type Intlist struct{    // LINKED-LIST
        Value int
        Tail  *IntList
    }

    func (list *IntList) Sum() int {
        if list == nil {
            return 0
        }
        return list.Value + list.Tail.Sum()
    }

https://youtu.be/0X6AcnwocbM?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6