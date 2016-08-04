# HTTPMock - Examples

In this folder, you'l find examples showing how to use this lib. Although the examples are complete and cover most of 
its functionality, there may be some missing features. Feel free to ask about them in the Bug tracker and I'l add 
examples here :)

I also highly recommend to look at the source code... I put a lot of comments there to explain what I am doing.

To run these examples, run the `/runnable/main.go` file.

## Quick explanation

| Filename          | Description |
|-------------------|-------------|
| serve_generic.go  | The most basic use case. All requests return the same response, no matter what |
| serve_multiple.go | The most common use case. Define pages and map them to addresses. When the client requests one of those addresses, the server will serve the related page |