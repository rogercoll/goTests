## Update on patch version

```
	cat go.mod
	go get -u 
	cat go.mod
```

## Update on minor version

```
	cat go.mod
	go get -u "github.com/xxx/yyy@v1.X.1 
	cat go.mod
```

## Update on major version

If you don't want to allow multiversion(create v2/ with a different go.mod inside) you must rename the actual go.mod too
```
	cat go.mod
	go get -u "github.com/xxx/yyy@vX.0.0
	cat go.mod
```
