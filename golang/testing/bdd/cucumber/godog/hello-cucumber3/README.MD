// this scenario is to show how to run the cucumber tests using go test when the e2e features are in a separate folder.
// compared to "hello-cucumber2" example, the features in a e2e folder.



unlike in "hello-cucumber2" example, the generated code cannot be in _test file. otherwise they cannot be imported in the main test. otherwise we will get error as below

```
s@mulys-new-mbp hello-cucumber3 % go test
go build github.com/muly/howto/golang/testing/bdd/cucumber/godog/hello-cucumber3/e2e: no non-test Go files in /Users/s/go/src/github.com/muly/howto/golang/testing/bdd/cucumber/godog/hello-cucumber3/e2e
FAIL    github.com/muly/howto/golang/testing/bdd/cucumber/godog/hello-cucumber3 [build failed]
```

however removing the _test from the  ile name will not allow running godog using cli

```
s@mulys-new-mbp e2e % godog
failed to compile testmain package: exit status 1 - output: compile: -importcfg: open /var/folders/dy/cs1csg6d0kvbby7cxz5rqphr0000gp/T/go-build1927167850/b001/importcfg.link: no such file or directory
```