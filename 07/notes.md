Formatted I/O


 - Standard inut
 - Standard ouput
 - Standard error


fmt.Println("printing to satndard output")
fmt.Fprintln(os.Stderr, "printing to error output")


The fmt package uses reflection and can print anything;
some of the function take a format string


// always os.Stdout

fmt.Println(...interface{}) (int, error)
fmt.Printf(string, ...interface{}) (int, error) 

// print to anything that has the correct Write() methof

fmt.Fprintln(io.Writer, ...interface{}) (int, error)
fmt.Fprintf(io.Writer, string, ...interface{}) (int, error)

// return a string

fmt.Sprintln(...interface{}) string
fmt.Sprintf(string, ...interface{}) string


# File I/O

Package os has functions to open or create files,
list directories, etc. and hosts the File Type.

Package io has utilities to read and wrote; bufio provides the buffered I/O scanners, etc.

Packagae io/ioutil has extra utilities such as reading an entire file to memory, or writing it out all at once.

package strconv has utilities to convert to/from string representations.



https://youtu.be/dqEtGT-dxoY?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6