Strings (Strings in go are unicode)

Types related to stings:
    - byte: a synonym for unit8
    - rune: a synonym for int32 for characters
    - string: an immutable sequence of "characters"
        - physically a squnece of bytes (UTF-8 encoding)
        - logically a sequence of (unicode) runes

Runes are enclosed in single quote 'a'

"Raw" strings use backtick quotes `string with "quotes"`

They also don't evaluate escape charactes such as \n


*The Internal string resprentation is a pointer and a length
*Strings are immutable and share underlying storage


s := "hello, world"
hello := s[:5]
world := s[7:]


[|...|h|e|l|l|o|,| |w|o|r|l|d|...|]


s => [ data | len 12 ]

hello => [ data | len 5 ]
hello => [ world | len 5 ]




Strings overload the addition operator (+ and += )

s := "the quick brown fox"

a := len(s)    // 19
b := s[:3]     // "the"
c := s[4:9]    // "quick"
d := s[:4] + "slow" + s[9:]    // "the quick slow fox" (Points to a new string)

s[5] = 'a'                     // SYNTAX ERROR
s += "es"                      // "the quick brown foxes" 

strings are pased by refrence, thus thet aren't copied


Package strings has many functions on string

s := "a string"

x := len(s)                    // built-in, = 8

string.Contains(s, "g")        // returns true
string.Contains(s, "x")        // returns false
strings.HasPrefix(s, "a")      // return true
strings.Index(s, "string")     // returns 2

s = string.ToUpper(s)           // retruns "A STRING" and assign to s, will point to new block at different from original location (not modified in-place because of immutability)



https://youtu.be/nxWqANttAdA?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6