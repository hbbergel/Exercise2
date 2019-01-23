# Python 3.3.3 and 2.7.6
# python foo.py
#! /usr/bin/env python

from threading import Lock, Thread

# Potentially useful thing:
#   In Python you "import" a global variable, instead of "export"ing it when you declare it
#   (This is probably an effort to make you feel bad about typing the word "global")
lock = Lock()

i = 0

def incrementingFunction():
    global i
    # TODO: increment i 1_000_000 times
    for t in range (0, 1000000):
        lock.acquire()
        i = i + 1
        lock.release()
    

def decrementingFunction():
    global i
    # TODO: decrement i 1_000_000 times
    for t in range (0, 1000000):
        lock.acquire()
        i = i - 1
        lock.release()



def main():
    global i

    incrementing = Thread(target = incrementingFunction, args = (),)
    decrementing = Thread(target = decrementingFunction, args = (),)
    
    # TODO: Start both threads
    incrementing.start()
    decrementing.start()
    
    incrementing.join()
    decrementing.join()
    
    print("The magic number is %d" % (i))


main()
