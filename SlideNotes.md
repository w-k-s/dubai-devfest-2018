# Slide Notes

## Hello World

- Every Go program is made up of packages.
- The starting point of an app is the main function in the main package
- Here we’re importing the “fmt” package for formatted I/O.
- When calling an exported method from an external package, the method name must be prefixed with the package name.
- You can run short go scripts using `go run`. 

## Arrays Slices & For Range

**What are the advantages of fixed-size arrays?**

- Size of array is known so the OS can allocate exactly as many bytes as are needed (time-efficient)
- Avoids cost associated with resizing arrays (resource efficient)
- Compiler optimizations such as loop unrolling

## Goroutines

**What's the difference between a process and a thread**

- A process is an execution environment e.g. a program or application. It has a seperate memory space. Communication between processes is done using IPC (e.g. pipes, sockets)
- Threads is an independent sequence of execution. It exists within a process and it can be thought of a lightweight process. They have their own execution environment and their own sandboxed memory space. The memory space of a thread is part of the memory space of the process (i.e. it’s shared memory). Threads can communicate with each directly
- [Diagram](https://i.stack.imgur.com/NVNge.jpg)
