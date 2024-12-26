# Closures. (Functions inside functions)

- Scope vs Lifetime
    -> Sope is static, based on the code at compile time.
    -> Lifetimne depends on program execution (runtime).

package xyz

func doIt() *int {
    var b int
    ...
    return &b  // In languages like C, this is not valid as once doIt finishes the memory coccuption occurs as &b can now point to some other data
}

Variable b can only beseen inside "doit", but its value will live on

The value (object) will live so long as part of the program keepos a pointer to it.


*Closure is  when a function inside another function "closes over"
one or more local variables of the outer function.

func fib() func() int {
    a, b := 0, 1

    return func() int {
        a, b = b, a + b
        return b
    }
}

The inner function gets refrence to the outer function's vars

Those variable may end up with a much longer lifetime than expected, as long as there's a refrence to the inner function.

- How Closure work: https://youtu.be/US3TGA-Dpqo?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&t=421



type kv struct {
    key string
    val int
}

var ss []kv

for k, v := range words {
    ss = append(ss, kv{k, v})
}

sort.Slice(ss, func(i, j int) boobl {
    return ss[i].val > ss[j].val
})

ss = ss[:3]

for _, s := range ss {
    fmt.Println(s.key, "appears", s.val, "times)
}




https://youtu.be/US3TGA-Dpqo?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6