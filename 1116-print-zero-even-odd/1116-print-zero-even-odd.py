from threading import Event, Lock

class ZeroEvenOdd:
    def __init__(self, n):
        self.n = n
        self.num = 0
        self.zero_event = Event()
        self.even_event = Event()
        self.odd_event = Event()
        self.zero_event.set()
        
	# printNumber(x) outputs "x", where x is an integer.
    def zero(self, printNumber: 'Callable[[int], None]') -> None:
        for _ in range(self.n):
            self.zero_event.wait()
            self.zero_event.clear()
            printNumber(0)
            self.num += 1
            if self.num % 2 == 0:
                self.even_event.set()
            else:
                self.odd_event.set()
        
        
    def even(self, printNumber: 'Callable[[int], None]') -> None:
        for _ in range(self.n//2):
            self.even_event.wait()
            self.even_event.clear()
            printNumber(self.num)
            self.zero_event.set()
        
        
    def odd(self, printNumber: 'Callable[[int], None]') -> None:
        for _ in range((self.n+1)//2):
            self.odd_event.wait()
            self.odd_event.clear()
            printNumber(self.num)
            self.zero_event.set()