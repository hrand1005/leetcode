from threading import Event

class Foo:
    def __init__(self):
        self.first_event = Event()
        self.second_event = Event()


    def first(self, printFirst: 'Callable[[], None]') -> None:
        printFirst()
        self.first_event.set()


    def second(self, printSecond: 'Callable[[], None]') -> None:
        self.first_event.wait()
        printSecond()
        self.second_event.set()


    def third(self, printThird: 'Callable[[], None]') -> None:
        self.second_event.wait()
        printThird()