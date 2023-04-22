from threading import Event

class FooBar:
    def __init__(self, n):
        self.n = n
        self.event_foo = Event()
        self.event_bar = Event()
        self.event_foo.set()

    def foo(self, printFoo: 'Callable[[], None]') -> None:
        for i in range(self.n):
            self.event_foo.wait()
            printFoo()
            self.event_foo.clear()
            self.event_bar.set()


    def bar(self, printBar: 'Callable[[], None]') -> None:
        for i in range(self.n):
            self.event_bar.wait()
            printBar()
            self.event_bar.clear()
            self.event_foo.set()