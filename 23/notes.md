# Select

select is a control structure that allows any "ready" alternative to proceed among
    - a channel we can read from
    - a channel we can write to
    - a default action that's always ready

Most often select runs in loop so we keep trying

We can put a timeout or "done" channel into select

    - We can compose chanels as synchronization primitives.
    - Traditional primitive (mutex, conditio variable) can't be composed.


# Select Example

In select block, the default case is always ready and will be choses if no other case is:

    func sendOrDrop(data []byte) {
        slect {
            case ch <-data;
                // sent ok; do nothing
            default:
                fmt.Printf("overflow: drop %d bytes", len(daat)) 
        }
    }

Don't use default inside a loop; the select will busy wait and waste CPU 