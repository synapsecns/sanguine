# Gin Helper

Note: gindump utilities are courtesy of [this repo](https://github.com/tpkeeper/gin-dump) and have only been updated/modified here because of dependency issues.
<!--TODO: document non gindump deps-->

* Gin middleware/handler to dump header/body of request and response .

* Very helpful for debugging your applications.

* More beautiful output than httputil.DumpXXX()

## Content-type support / todo

* [x] application/json
* [x] application/x-www-form-urlencoded
* [ ] text/xml
* [ ] application/xml
* [ ] text/plain

## Usage

All:

```go
func main() {
    router := gin.Default()

    //use Dump() default will print on stdout
    router.Use(gindump.Dump())

    //or use DumpWithOptions() with more options
    router.Use(gindump.DumpWithOptions(true, true, false, true, false, func(dumpStr string) {
	    fmt.Println(dumpStr)
    }))

    router.Post("/",myHandler)

    ...

    router.Run()
}
```

Group:

```go
func main() {
    router := gin.Default()

    dumpGroup := router.Group("/group")

    //use Dump() default will print on stdout
    dumpGroup.Use(gindump.Dump())

    //or use DumpWithOptions() with more options
    dumpGroup.Use(gindump.DumpWithOptions(true, true, false, true, false, func(dumpStr string) {
	    fmt.Println(dumpStr)
    }))

    dumpGroup.Post("/",myHandler)

    ...

    router.Run()
}

```

EndPoint:

```go
func main() {
    router := gin.Default()

    //use Dump() default will print on stdout
    router.Post("/",gindump.Dump(),myHandler)

    //or use DumpWithOptions() with more options
    router.Post("/",gindump.DumpWithOptions(true, true, false, true, false, func(dumpStr string) {
	    fmt.Println(dumpStr)
    }),myHandler)

    ...

    router.Run()
}
```


### Output is as follows

```sh
[GIN-dump]:
Request-Header:
{
    "Content-Type": [
        "application/x-www-form-urlencoded"
    ]
}
Request-Body:
{
    "bar": [
        "baz"
    ],
    "foo": [
        "bar",
        "bar2"
    ]
}
Response-Header:
{
    "Content-Type": [
        "application/json; charset=utf-8"
    ]
}
Response-Body:
{
    "data": {
        "addr": "tpkeeper@qq.com",
        "name": "jfise"
    },
    "ok": true
}
```
