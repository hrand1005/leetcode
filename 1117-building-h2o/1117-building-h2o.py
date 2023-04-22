from threading import Barrier, Semaphore

class H2O:
    def __init__(self):
        self.group = Barrier(3)
        self.h = Semaphore(2)
        self.o = Semaphore(1)


    def hydrogen(self, releaseHydrogen: 'Callable[[], None]') -> None:
        self.h.acquire()
        self.group.wait()
        releaseHydrogen()
        self.h.release()


    def oxygen(self, releaseOxygen: 'Callable[[], None]') -> None:
        self.o.acquire()
        self.group.wait()
        releaseOxygen()
        self.o.release()