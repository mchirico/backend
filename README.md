


## Don't forget golint

```

golint -set_exit_status $(go list ./... | grep -v /vendor/)

```


