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
$ hexstohexa aabbccdd ddeeff0102 [...hex strings]
```

Output:

```
{ 0xAA, 0xBB, 0xCC, 0xDD }
{ 0xDD, 0xEE, 0xFF, 0x01, 0x02 }
```