Note: This is a fork of [schollz/sqlitedump](https://github.com/schollz/sqlite3dump). We could not use the original package because it uses am old version of `mattn/go-sqlite3`.

<p align="center">
<strong>sqlitedump</strong>
<br>
<img src="https://img.shields.io/badge/coverage-79%25-green.svg?style=flat-square" alt="Code Coverage">
<a href="https://godoc.org/github.com/schollz/sqlitedump"><img src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square" alt="Code Coverage"></a>
</p>


This is a Golang port of [Python's `sqlite3 .iterdump()`](https://github.com/python/cpython/blob/3.6/Lib/sqlite3/dump.py) command. This is written to use [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3). Both are supposed to be functionally equivalent to `sqlite3 DATABASE .dump`.

There is also a command-line tool that you can use.

```
$ go get github.com/schollz/sqlite3dump/...
$ sqlite3dump database.db > database.sql
```

# License

MIT
