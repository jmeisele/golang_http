# Go lang server using net/http package

Simple HTTP backend with Go using the net/http package and adding 3 endpoints serving 2 GET requests and 1 form.

To run locally:

- `go run github.com/jmeisele/golang_http`

This will spin up the webserver on port 8081. Open a browser window and go to `http://localhost:8081/` to see our simple home page with index.html being served, or `http://localhost:8081/second` to see our second page without an html page, finally `http://localhost:8081/form.html` to fill out our simple form.
