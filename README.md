[![Build Status](https://dev.azure.com/mchirico/backend/_apis/build/status/mchirico.backend?branchName=master)](https://dev.azure.com/mchirico/backend/_build/latest?definitionId=32&branchName=master)



## Don't forget golint

```

golint -set_exit_status $(go list ./... | grep -v /vendor/)

```


