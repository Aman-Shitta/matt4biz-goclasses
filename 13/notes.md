# Refrences and Value Semantics (When to se pointers or values)

| Pointers           |      Values       |
| Shared, not copied | Copied not shared |

Valuse semantics lead to higher integrity particularly with concurrency (don't share)

Pointer semantics mayt be more efficient.

# Pointers VS values

Common use of pointers
    - some objects can't be copied safely (mutex)
    - some objects are too large to copy effeciently
        (consider pointers when size > 64 bytes)
    - some methods need to change (mutate) the reciever [later!]
    - when decoding protocol data into an object
        (JSON etc, often in variable argument list)


        var r Response

        err := json.Unmarshal(j, &r)
    
    - when using pointer to signal a "null" object.

*MUST NOT COPY e.g

    - Any struct with a mutex `must` be passed by refrence:

    type Employee struct{
        mu sync.Mutex
        Name string
        ...
    }

    func do(emp *Employee) {
        emp.mu.Lock()

        defer emp.mu,Unlock()
        ...
    }

*MAY COPY
    - A\ny small struct under 64bytes probaly should be copied:

type Widget  struct{
    ID int
    Count int
}

func Expend(w Widget) Widget {

    w.Count--

    return w
}

**Note that go rountienly copies string & slice descriptiors

# Semantic consistency

If a thing is to be shared, then always pass a pointer.

        [ |Employee rolaction| ]
                   |
            f1(emp *Employee)
                   |
            f2(emp *Employee)
                   |
            f3(emp *Employee)
                   |
            f4(emp Employee)   // Passes a Copy
                   |
            f5(emp *Employee)  // Chnages are lost, as nowe changes will be made on copy


# Stack allocation

 - slack allocation is more efficient

 - Accessing a variable directly isd more efficient than following pointer

 - Accesing a dense sequence of data is more efficient that sparse data
    (an array is faster tah a linked list, etc)


# Heap allocation

Go would prefer to allocate on the stack, but sometimes can't

 - a function returns a pointer to local object.
 - a local object is captured in function closure.
 - a pointer to a local object is sent via a channel
 - any object assigned into an interface
 - any object whose size is variable at runtime (slices)


# for loops

The value returned by range is always a copy

    for i, value := range things {
        // value is a copy
    }

Use the index if you need to mutate the element.

    for i := range things{
        things[i].which = whatever
        ...
    }

# Slice safety

Anytime a function mutates a slice that's passed in, we must return a copy

func update(things []thing) []thing {
    ...
    things = append(things, x)

    return things
}

That's because the slice's backing array may be reallocated to grow.

Keeping a pointer to an elemt of slice is risky

    type user struct {name string; count int}

    func addTo(u *User) { u.count++ }

    func main() {
        users := []user{{"alice": 0}, {"bob": 1}}

        alice := &users[0]     // risky (0xc0000a8120)

        amy := user{"amy", 2}

        users = append(users, amy)   // reallocation, alice points to @0xc0000ac120

        addTo(alice)          // alice is likely a stale pointer
        fmt.Println(users)    // so alice's count will be 0
    }


# Capturing the loop variable, again
 - Taking the address of a mutattin loop variable is wrong (*before go 1.22)
    
    func (r OfferResolver) Chnages() []ChangeResolver {
        var result []ChangeResolver 

        for _, change := range r.d.Status.Changes {
            result = append(result, ChangeResolver{&change})    // WRONG
        }

        return result
    }

    WRONG: all the returned resolvers pointes to the last change in the list (*before go 1.22, after 1.22 works as expected)

    func (r OfferResolver) Chnages() []ChangeResolver {
        var result []ChangeResolver 

        for _, c := range r.d.Status.Changes {
            change := c                     // FIX
            result = append(result, ChangeResolver{&change})
        }

        return result
    }


```
    Before Go 1.22

        The Issue with Capturing Loop Variables:
            In versions prior to Go 1.22, the loop variable (in your example, change) is reused across iterations of the loop.
            When you take the address of the loop variable (e.g., &change), all iterations of the loop will point to the same memory location — the loop variable itself, not the actual array elements.
            As a result, all appended ChangeResolver{&change} entries in result end up pointing to the last element of r.d.Status.Changes.

        for _, change := range r.d.Status.Changes {
            result = append(result, ChangeResolver{&change}) // All point to the same 'change'
        }

        The Outcome:
            All the resolvers in the result slice will hold the address of the loop variable change, which is overwritten in each iteration. After the loop finishes, all resolvers will point to the last element of r.d.Status.Changes.

    Go 1.22 and Beyond

    In Go 1.22, the behavior of loop variables was fixed. Each iteration now creates a new instance of the loop variable, so taking the address of the loop variable (e.g., &change) will now correctly refer to the current element of the iteration.

        The code works as expected, and ChangeResolver{&change} will point to the correct element in r.d.Status.Changes for each iteration.

    Best Practices

    **IMP: Even though Go 1.22 fixes this behavior, it’s still considered better practice to avoid taking the address of a loop variable directly. Instead, you can use an explicit reference to the array element:

    for i := range r.d.Status.Changes {
        change := r.d.Status.Changes[i] // Create a separate variable
        result = append(result, ChangeResolver{&change})
    }

    This pattern works consistently across all Go versions and avoids relying on compiler-specific fixes.
```