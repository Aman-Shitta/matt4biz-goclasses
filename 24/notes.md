# Context


- Cancellation and timeouts

The Context package offers a common method to cancel requests.

    - explicit cancellation
    - implicit cancellation based on timeout or deadline

A context may also carry request-specific values, such as traceID

Many network of database requests, for example, take a context for cancellation.


A context offers two controls
    - a channel that closes when the cancellation occurs
    - an error that's readable once the channel closes.

The error value tells you weather the request cancelled or timed out

We often use the channel from Done() in a select block.



Context form an immutable tree structure
(goroutine-safe; changes to a context doesn't affect its ancestors)


Cancellation or timeout applies to the current context and it's subtree


Ditto for a value


A subtree ma be created with a shorter timeout (but no longer)

# Context example

- The context value should always be the first parameter


    // First runs a set of queries and returns the result from the
    // the first to respond, cancelling the other

    func First(ctx context.Context, urls []string) (*Result, error) {
        c := make(chanResult, len(urls))        // buffered to avoid oprphans
        ctx, cancel := context.WithCancel(ctx)

        defer cancel()     // cancel the other queries when we're done

        search := func(url string) {
            c <-runQuesry(ctx, url)
        }

        ...

    }


# Values

Context values should be data specific to request, such as:
    - a trace ID or start time (for latency calculation)
    - security or authorization data


Avoid using context to carry optional parameters


Use package-specific, private context key type (not string) to avoid collisons



    type contextKey int

    const TraceKey contextKey = 1

    // AddTrace is a HTTP middleware to insert a trace ID into the request

    func AddTrace(next http.Handler) http.Handler {

        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)) {
            ctx := r.Context()

            if traceID := r.Header.Get("X-Cloud-Trace-Cnntext"); traceID != "" {
                ctx = context.WithValue(ctx, TraceKey, traceID)
            }
        next.ServeHTTP(w, r.WithContext(ctx))
        }
    }


    Example:

    type contextKey int

    const TraceKey contextKey = 1

    // ContextLog makes a log with the trace ID as prefix.
    func ContextLog(ctx context.Context, f string, args ...interface{}) {
        //reflection -- TBD

        traceID, ok := ctx.Value(TraceKey).(string)

        if ok && traceID != "" {
            f = traceID +  ": " + f
        }

        log.Printf(f, args)
    }