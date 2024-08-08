## I am unable to upgrade the modules anymore in Go 1.22.6 

I was able to upgrade the modules using the following steps in older version of golang but not working in newer version

```sh
cd hello
go get -u ./...
```

ERROR:

```sh
➜  hello go get -u ./...
go: hello imports
	foo/bark: package foo/bark is not in std (/usr/local/go/src/foo/bark)
```