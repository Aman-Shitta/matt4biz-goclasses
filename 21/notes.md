# What is concurrency

"Execution happens in some non-determisinstic order"

"Undefined out of order execution"

"Non-sequential execution"

"Parts of program execute out-of-order or in particular error"


# Parallelism

- Parts of program execute independently at the same time

- You can have concurrency with a sing;e-core processor
   (think interrupt hadnlind in OS)

- Parallelism can happen only in multicore processor

- Concurrency doen't necessarily makes the program faster parallelism does.

# Concurrency vs parallelism

Concurrency is about dealing with things happening out-of-order
Parallelism is about things actually happening at the same time.

A single program won't have parallelism without concurrency
We need concurrency to allow parts of program to exiecute independently

And that's where the fun begins...


# Race Condition

"System behaviour depends on the (non-determistic) sequence
or timing of parts of the program executing independently,
where some possible behaiors (orders of execution) produce invalid results"

