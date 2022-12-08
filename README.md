# GRPC for UfaElcetro

---

## Installation

**For security reasons, all files that can be compromised have been moved to **`.gitignore`** and used through the virtual environment. Therefore, you need to enter your own data and change the code a little.**

**It uses a clean dependency injection architecture. In the project, all the functions are empty and return a nil, if you wish, you can modify it by changing it for yourself.**

---

**To run the Go api use this commands:**
```go
go mod build
go run aapi/cmd/main.go
```
**or**
```go
go mod build
go guild api/cmd/main.go
```
**also you can use makefile**
```makefile
make run
make build
```

**If you change the description in the proto file, you must run the following command:**
```makefile
make gen
```

**Look for its description in the makefile**
___

### To contact with me:

#### [Site link](https://kron-x.ru/)
#### [My Telegram](https://t.me/amirich18)