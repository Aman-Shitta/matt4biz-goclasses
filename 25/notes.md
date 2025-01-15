# Channels in detail

# Channels and Synchronization

- Channel state

Channel blocks unless ready to read or write


A channel is ready to write if
    - it has buffer space, or
    - atleast one reader is ready to read (rendezvous)

A channel is ready to read if
    - it has unread data in its buffer, or
    - atleast one writer is ready to write (rendezvous), or
    - it is closed.

Channels are unidirectional, but have two ends
(which can be passed seperately as parameters)

- An end for writing & closing

    func get(url string, ch chan<- result) {    // write only end
        ...
    }

- An end for reading

    func collect(ch <-chan result) map[string]int {   // read only end
        ...
    }


- Closed Channels

Closign channel cause its to return the "zero" value

We can recieve a second value: is the channel close ?


    func main() {
        ch := make(chan int, 1)

        ch <- 1
        b, ok := <-ch
        fmt.Println(b, ok)   // 1 true

        close(ch)

        b, ok := <-ch
        fmt.Println(b, ok)  // 0 false
    }

A channel can only be closed once, else it will panic

One of the main issues working with goroutines is ending them
    - An unbuffered channel requires a reader and writer
     (a writer blocked on a channel with no reader will "leak")

    - Closing channel is often a signal that work is done

    - Only goroutine can close a channel (not many)

    - We may need to some way to coordinate closing channel or stopping goroutines
        (beyond the channel itself)

- Nil channels

Reading or writing a channel that's nil always blocks

But a nil channel in a select block is ignored

This can be powerful tool:
    - Use a channel to get input
    - Supspend it by changing the channel varable to nil
    - You can even un-suspend it again
    - But close the channel if there really is no more input (EOF)

* select ignores a nil channel since it would always block
** reading a closed channel returns (<default-value>, !ok)


# Buffering

Buffering allows the sender to send without waiting

    func main() {
        // make a channel wioth buffer that holds two items
        ch := make(chan int, 2)

        // now we send twice without getting blocked
        ch <- "buffered"
        ch <- "channels"

        // and then recieve both as usual
        fmt.Println(<-ch)
        fmt.Println(<-ch)
    }

With Buffer size 1 (or nu buffer at all), it will deadlock

Common uses of buffered channels
    - avoid goroutines leaks (from an abandoned channel)
    - avoid rendezvous pauses (performance improvement)

Don't buffer until it's needed; buffering may hide a race condition

Some testing may be required to find the right number of slots

Special uses of buffered channels:
    - counting semaphore pattern