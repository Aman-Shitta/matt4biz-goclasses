# CSP, Goroutines and Channels

# Channels

A channel is one-way communication pipe

    - things go in one end, come out of order
    - in the same order they went in
    - until the channel is closed
    - multiple readers & writers can share it safely

# Sequential process

Looking at a single indepenedent part of program its appears to be sequential


    for {
        read()
        process()
        write()
    }

This is perfectly natural if we think of reading and writing files or network sockets.

# Communicating sequential processes(CSP)

Now we put the parts together with channels to communicate

    - each part is independent
    - all they share are the channels between them
    - the parts can run in parallel as hardware allows

- Why is CSP valuable ?

    - Concurrency is always hard
    - CSP provides a model for thingking about it that makes it less hard
        (take the program apart and make the pieces talk to each other)
    
    " Go deosn't enforce eveloper to embrace the asynchronous way of event drived programming... That 
      lets you *write asynchronous code in synchronous style*. As people, we're much better suited to
      writeing about things in synchronous style" -- Andrew Gerrand
# Goroutines

A goroutine is a unit of independent execution (coroutine)

It's easy to start a goroutine: put go keyword in front of function call

The trick is knowing how goroutine will stop:
    - you have a well-defined loop terminating condition, ort
    - you signal completion through a channel or context, or
    - you let it run until the program stops

But you need to make sur its doesn't get blocked by mistake

# Channels

A channel is like a one-way socket ot Unix pipe
(except it allows mutiple readers & writers)

It's a method of syncronization as well as communication

We know that a send(write) always happend before a recieve(read)

It's also a vehichle fro transferring ownership of data,
so that only one goroutin at a time is writing the data (avoid race conditions)

"Don't communicate by sharing memory, instead share memory by communicating"