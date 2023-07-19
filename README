# Go and wasm

> Warning!
> This is an example of using wasm in go. This example is buggy and should only be used as an early taste. It should not be used in production

## Use

> Only the Jebt series products are recommended.
> Other tools can be configured by themselves

1. set wasm env and build go file to wasm

```shell
GOOS=js GOARCH=wasm go build -o static/main.wasm
```

2. Copy wasm exec file to root fold

```shell
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```


3. Run web server in root fold

> You can use any tool that can start the web service to run it

if you installed `goexec`, not installed? `go get -u github.com/shurcooL/goexec`
```shell
goexec 'http.ListenAndServe(`:9999`, http.FileServer(http.Dir(`.`)))'
```

Or

if you installed `http-server`, not installed? `npm install -g http-server`
```shell
http-server
```

4. Open a browser and type `localhost:8080`

## Tips

1. Each change requires recompilation

> Makefile used `goexec` tool, If you don't like it, feel free to replace it in the Makefile
```shell
make
```
