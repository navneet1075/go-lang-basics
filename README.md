 **If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will panic at runtime with Deadlock.**

 **it is possible to convert a bidirectional channel to a send only or receive only channel but not the vice versa.**

 **Senders have the ability to close the channel to notify receivers that no more data will be sent on the channel.**