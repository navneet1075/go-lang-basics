 **If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will panic at runtime with Deadlock.**

 **it is possible to convert a bidirectional channel to a send only or receive only channel but not the vice versa.**

 **Senders have the ability to close the channel to notify receivers that no more data will be sent on the channel.**

 A problem with reentrancy (Reader Writer Locks)
 Recursive mutexes solve the problem of non-reentrancy with regular mutexes: if a function that takes a lock and executes a callback is itself called by the callback, deadlock ensues.[1] In pseudocode, that is the following situation:

var m : Mutex  // A non-recursive mutex, initially unlocked.

function lock_and_call(i : Integer)
    m.lock()
    callback(i)
    m.unlock()

function callback(i : Integer)
    if i > 0
        lock_and_call(i - 1)
Given these definitions, the function call lock_and_call(1) will cause the following sequence of events:

m.lock() — mutex locked
callback(1)
lock_and_call(0) — because i > 0
m.lock() — deadlock, because m is already locked, so the executing thread will block, waiting for itself.
Replacing the mutex with a recursive one solves the problem, because the final m.lock() will succeed without blocking.

important concepts with links:
**https://github.com/golang/go/wiki/MethodSets**

**defer statments**
1. A deferred function's arguments are evaluated when the defer statement is evaluated.
2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
3.Deferred functions may read and assign to the returning function's named return values.


links: https://medium.com/rungo/the-anatomy-of-maps-in-go-79b82836838b


Sync.Pool design :

GoLang internal pool design :


GoLang internal pool design :


Get method : 

	```1. request internal pool :

	if first usage : then create a pool (one poolLocal per processor (GOMAXPROCS))

	if not return the pool used by the local processor

	each poolLocal has 2 attributes : private and shared```


	2. Does the private attibute of the internal pool  has a item : 
		a. if yes -> remove the object from internal pool local return it.
		b. if no -> does the shared attribute of any other processors pool local have a item ?.
		 1. if yes ->  steal it
		 2. if no -> then create a new item using New function and store in the private attribute of the localPool.



Put method :

1. if the internal pool local has a item in private attribute ?

	```yes : -> stored in shaed attribute of the internal pool local of the processor on which the goroutine is running.
	
	   no : -> stored in the private atribute of the internal pool local of the processor on which the goroutine is running.```
