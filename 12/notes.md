# Structs, Structs tags & JSON

Structs 

# Struct compatability

Two structs are compatible if :
    - the fileds have the same type and names
    - in the same order
    - and the same tags(*)

- A struct may be copied or passed as parameter in its entirety.
- A struct is comparable if all its fileds are comaprable.
- The zero valuue of struct is "zero" for each field in turn.


Structs of anynymous types with same varaibels and type and order can be copeid to each other 
because of structured compatibility e.g

func main() {
    album1 = struct{
        title string
    } {"The White album,}

    album2 = struct{
        title string
    } {"The Black album,}

    album1 = album2 // can be copied
}

but named struct with same varibale and type order cannot because of named compatibilty issue.

struct album1 {
    title string
}

struct album2 {
    title string
}

func main() {

    var a1 = album1 {
        "The white album"
    }
    var a2 = album1 {
        "The black album"
    }

    a1 = a2 // cannot use a2 (type ablbum2) as type album1 in assignemnt 
    a1 = album1(a2) // Conversion is possible
}

Anonymous structs are compatible if they follow

func main() {
    v1 := struct{
        X int `json:"foo"` // `json:"foo"` ---> struct tag for JSON
    }{1}

    v2 := struct{
        X int `json:"foo"`
    }{2}

    v1 = v2 
    fmt.Println(v1) // prints {2}
}

Named struct are convertible if they are compatible (i.e same fields and types)

*From Go 1.8 tag difference don't prevet type conversion, before that the tag difference makes struct in-cpompatible

func main() {
    v1 := struct{
        X int `json:"foo"`
    }{1}

    v2 := struct{
        X int `json:"bar"`
    }{2}

    v1 = v2  // before 1.8 error, after 1.8 supported
    fmt.Println(v1) // prints {2}
}

Structs are passed by value unless a pointer is used.

var while album

func soldAnother(a album) {
    a.copies++
}

func soldAnother(a *album) {
    // orignal album
    a.copies++
}

**Note that "dot" notaion works in pointers too, equivalent to (*a).copies

# Struct Tags and JSON

type Response struct{
    Page int `json:"page"`
    Words []sting `json:"words,omitempty"`
}

func main() {
    r := &Response{Page:1, Words: []string{"up, "in", "out}}
    
    j, _ := json.Marshall(r)
    fmt.Println(string(j))
    fmt.Printf("%#v\n", r)

    var r2 Response

    _ = json.Unmarshall(j, &r2)
    fmt.Printf("%#v\n", r2)
}