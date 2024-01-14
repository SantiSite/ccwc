# CCWC

Version in Go of the Unix command line tool **wc**.

It was carried out in order to overcome the [Coding Challenges](https://codingchallenges.fyi/) challenge.

## Install

You must have Go installed and configured on your machine.

### Step 0
```shell
git clone git@github.com:SantiSite/ccwc.git
```

### Step 1
```shell
cd ccwc
```

### Step 2
```shell
go install
```

## Usage

How to use:
```shell
ccwc --help
```

Print the word counts:
```
$ ccwc -w <path/to/your/file>
```

### Examples:
```
$ ccwc main.go
output:
      newlines           words           bytes           chars
            135             382            2549            2549    main.go
```

```
$ cat main.go | ccwc
output:
       newlines           words           bytes           chars
            135             382            2549            2549    main.go
```

