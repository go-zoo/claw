claw [![GoDoc](https://godoc.org/github.com/squiidz/claw?status.png)](http://godoc.org/github.com/squiidz/claw)
=======

## What is claw ?

Claw is a Middleware chaining module, compatible with
every mux who respects the ` http.Handler ` interface. Claw allows you
to create shcemas for specific tasks.

![alt tag](http://upload.wikimedia.org/wikipedia/commons/thumb/7/7e/Claw.jpg/640px-Claw.jpg)

## Features

- Uses Standard ` func (http.ResponseWriter, *http.Request) ` as Middleware.
- Also uses ` func(http.Handler) http.Handler ` as Middleware.
- Middleware Chaining.
- Global Middleware.
- Middleware Schema.
- Compatible with every mux that implements ` http.Handler ` interface.

## Example
```go
package main

import "github.com/squiidz/claw"

func main() {
	// Create a new Claw instance, and set some Global Middleware.
	c := claw.New(GlobalMiddleWare)

	// You can also, create a Schema(), which is a stack
	// of MiddleWare for a specific task
	auth := c.NewSchema(CheckUser, CheckToken, ValidSession)

	// Wrap your global middleware with your handler
	http.Handle("/home", c.Use(YourHandler))

	// Add some middleware on a specific handler.
	http.Handle("/", c.Use(YourOtherHandler).Add(OtherMiddle)) 

	// Add a Schema to the route.
	http.Handle("/", c.Use(YourOtherHandler).Schema(auth)) 

	// Start Listening
	http.ListenAndServe(":8080", nil)
}
```

## TODO
- Refactoring
- Debugging

## Contributing

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Write Tests!
4. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request

## License
MIT
