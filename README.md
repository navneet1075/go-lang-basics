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


**Sync.Pool internal design**

**Get method** 

	1. request internal pool :

		a.if first usage : then create a pool (one poolLocal per processor (GOMAXPROCS))
		Note : Each poolLocal has 2 attributes : private and shared (type interface{})
		b.if not return the pool used by the local processor


	2. Does the private attibute of the internal pool  has a item : 
		a. if yes -> remove the object from internal pool local return it.
		b. if no -> does the shared attribute of any other processors pool local have a item ?.
		 	1. if yes ->  steal it
		 	2. if no -> then create a new item using New function and store in the private attribute of the localPool.

**Put method**
```
1. internal pool local has a item in private attribute ?
	yes : -> stored in shaed attribute of the internal pool local of the processor on which the goroutine is running.
	no : -> stored in the private atribute of the internal pool local of the processor on which the goroutine is running.
```

**go maps**

1. maps iteration order is not guaranteed.
2. maps are not goroutine safe . use a struct with sync.RWMutex or some other concurrent safe primitive to have the goroutine behavior. 
	
	type struct Constructor {
		
		sync.RWMutex
		myMap map[int]string
	
	}
	cons := Constructor {
		
		myMap = make(map[int]string)
	
	}	
	
	to read :
	cons.RLock()
	defer cons.RUnlock()
	// read the data here
	
	to write :
	cons.Lock()
	defer cons.Unlock()
	// write the data here
	
	

Memory Management in Go:

1. checking the memory allocation in program with memory flags
	```go build -gcflags "--m -m" program-name.go
	```
2. go tool compile "-m" program-name.go 

generate assembly and see the calls 
3. go tool compile -S program-name.go
	
	
scheduler tracing can be done using :
GODEBUG=scheddetail=1,schedtrace=1000 ./'program-name'

important articles on go memory management and scheduler:

1. https://blog.learngoprogramming.com/a-visual-guide-to-golang-memory-allocator-from-ground-up-e132258453ed
2. https://povilasv.me/go-scheduler/
3. https://www.ardanlabs.com/blog/2015/02/scheduler-tracing-in-go.html
4. http://www.cs.columbia.edu/~aho/cs6998/reports/12-12-11_DeshpandeSponslerWeiss_GO.pdf
5. https://morsmachine.dk/go-scheduler


Use Cgo with Python :

https://www.datadoghq.com/blog/engineering/cgo-and-python/

