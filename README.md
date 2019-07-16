# klyngen/jsend



**NOT ANYMORE**

This really basic JSend implementation does so much for you. The package basically just exports one Function. It is that easy

# Installation

**The usual way** `go get github.com/klyngen/jsend`

# Usage

This is how simple we can make the response...

```
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    jsend.FormatResponse(w, "Time to read the documentation?", jsend.NotFound)
}
```

> Fisty comments in the message body is optional

## Use with gorrila/mux

[Mux](https://github.com/gorilla/mux) is a great router. What is really nice, is that mux allows us to replace its NotFound-handler and MethodNotAllowed-handler; like so:



```go
	router.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowedHandler)
```



## The results speak for themselves

By using this library, you don't need to pay attention to the structure of the Jsend-formatted message. The library does that for you. 



```
{
    "status": "client-error",
    "message": "We do not support that method for this endpoint"
}
```

Since the data was a string and the type was an error, we know that data should not be present.

Not-mandatory fields are omitted, as they should



## Abbreviations from the standard

Becauce I like it that way, the status-object can be: `success`, `server-error` and `client-error`. 


