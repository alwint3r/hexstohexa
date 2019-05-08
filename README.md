hexstohexa
==========

Convert hex string to hex array for C/C++ code.

Why? Because I'm so bad at naming things.

### Installing

```sh
go get github.com/alwint3r/hexstohexa/cmd/...
```

### Usage

```sh
$ hexstohexa aa bb cc dd [...hex char]
```

Output:

```
{ 0xAA, 0xBB, 0xCC, 0xDD }
```