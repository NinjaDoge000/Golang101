
# go routine

A goroutine is a lightweight thread managed by the Go runtime. It allows you to run functions concurrently (at the same time) with other code.

 go keyword before a function call

 Goroutines are much cheaper than system threads because they are managed by the Go runtime, not the operating system. Here’s why:

 + Smaller initial memory: Each goroutine starts with a small stack (about 2 KB), while OS threads usually need much more (often 1 MB).

 + Growable stacks: Goroutine stacks grow and shrink as needed, saving memory.

 + User-space scheduling: The Go runtime schedules goroutines in user space, multiplexing many goroutines onto a smaller number of OS threads. This reduces context-switching overhead.

 + Fast creation and destruction: Creating and destroying goroutines is much faster than OS threads.

 ## WaitGroups

 A WaitGroup in Go (from the sync package) is a synchronization primitive that lets you wait for a group of goroutines to finish.

You use it to make sure your main program waits until all goroutines are done before exiting.

 How it works:

+ `wg.Add(n)`: Set the number of goroutines to wait for.
+ Each goroutine calls `wg.Done()` when finished.
+ `wg.Wait()`: Blocks until all goroutines have called Done().

Never try to get them below Zero!

## Mutex

A mutex (short for "mutual exclusion") in Go is a synchronization primitive from the sync package. It is used to protect shared data from being accessed by multiple goroutines at the same time, preventing race conditions.

**How to use a mutex in Go** 
+ Lock: mtx.Lock() — Only one goroutine can hold the lock at a time.
+ Unlock: mtx.Unlock() — Releases the lock so other goroutines can acquire it.