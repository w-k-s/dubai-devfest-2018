# Audience Questions

**1. Wasn't there a controversy that Go would not add generics?**

The authors of Go were never against generics; they felt that current implementations of Generics (such as those in C++) add a great deal of complexity. The code is difficult to read, and the error messages even more so.

Contracts offer a clean and simple way to implement Generics, which is why the Go team is considering it.

**2. Has dependency management improved in Go?**

Context: 

Go packages can be added to `GOPATH` using the `go get` command e.g. `go get github.com/SermoDigital/jose/jwt`.

The problem was that `go get` would only fetch the latest commit in the master branch. This made it difficult to pin your source code to a specific version of a third-party library.

Answer:

A new dependancy management tool has been introduced since v1.9 called [`dep`](https://golang.github.io/dep/docs/introduction.html). This tool allows you to specify dependencies in a `.toml` file mentioning repo url, branch and/or commit.

**3.a. Are they are dependancy injection libraries for Go?**

[dig by uber](https://github.com/uber-go/dig)

**3.b. Are they are annotation-based dependency injection frameworks for Go (similar to Java Beans)**

Firstly, 

The way beans work in Java is that a class is loaded and instantiated from the class path using its fully qualified package name.

As far as I am aware, this is not possible in Go.

Secondly,

`Beans` and `BeanFactories` are very much the opposite of Go's philosophy. The idea is to avoid such complex abstractions and to make an application that's lightweight ans straight-forward.


