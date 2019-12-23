itfmc: count the number of an interface methods and output the statistics

The example below outputs the interfaces method statistics under \$GOROOT：

```
$ find $GOROOT -name '*.go'|grep -v "_test.go" |xargs cat| grep -v ^$|grep -v "^[[:space:]]*\/\/"|itfmc 2>/dev/null
a[0] = 13
a[1] = 315
a[2] = 113
a[3] = 36
a[4] = 14
a[5] = 13
a[6] = 11
a[7] = 6
a[8] = 4
a[9] = 1
a[11] = 1
a[12] = 2
a[13] = 1
a[16] = 1
a[17] = 1
a[31] = 1
interfaces in total = 533
interfaces in total by name= 314
```

* ```a[1] = 315``` means there are **315** interfaces which only has **one** method.
* ```a[2] = 113``` means there are **113** interfaces which only has **two** method.
* ... ...

The input data passed to itfmc should be follow these rules:
* no blank line: ```grep -v ^$```
* no comment line: ```grep -v "^[[:space:]]*\/\/"```
