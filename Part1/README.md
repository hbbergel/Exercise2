# Mutex and Channel basics

### What is an atomic operation?
> In concurrent programming - meaning program operations that run completely independently of any other process. Used in many modern operating systems and parallel processing systems. 

### What is a semaphore?
> A programming concept that is frequently used to solve multi-threading problems. A way to limit the number of consumers for a specific resource. 

### What is a mutex?
> *Provides mutual exclusion, either producer or consumer can have the key (mutex) and proceed with their work. As long as the buffer is filled by producer, the consumer needs to wait, and vice versa. 

### What is the difference between a mutex and a binary semaphore?
> Binary semaphore is signaling mechanism (listening to songs while recieving a call) and does not protect a recourse from access. Mutex is a locking mechanism used to synchronize access to a resource. Used to protect shared resources. 

### What is a critical section?
> Part of a program that accesses shared resources. Only when a process is in its critical section can it be in a position to disrupt other processes. 

### What is the difference between race conditions and data races?
 > Race condition is a flaw that occurs when the timing or ordering of events affects a program's correctness. Data race happens when there are two memory accesses in a program where both target the same location. 

### List some advantages of using message passing over lock-based synchronization primitives.
> 

### List some advantages of using lock-based synchronization primitives over message passing.
> *Your answer here*
