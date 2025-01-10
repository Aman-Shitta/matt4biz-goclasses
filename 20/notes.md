# Interfaces in http

    type Handler interface{
        ServeHTTP(ResponseWriter, *Requet)
    }

    type HandlerFunc func(ResponseWriter, *Request)

    type (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
        f(w, r)
    }

    // handler matched type handlerFunc and so interface Handler
    // so tht HTTP framework can call ServeHTTP on it

    func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, world! from %s\n", r.URL.Path[1:])
    }



# Assignment #4 (GOPL 7.11)

Exercis e 7.11: Add addition al handlers so that clients can create, read, update, and delete
database ent ries. For example, a request of the form /update?item=socks&price=6 will
up date the price of an item in the inventory and report an error if the item does not exist or if
the price is invalid. (Warning: this change int roduces conc urrent var iable updates.)
