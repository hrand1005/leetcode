from threading import Barrier, Semaphore

class H2O:
    def __init__(self):
        self.group = Barrier(3)
        self.h = Semaphore(2)
        self.o = Semaphore(1)


    def hydrogen(self, releaseHydrogen: 'Callable[[], None]') -> None:
        with self.h:
            self.group.wait()
            releaseHydrogen()


    def oxygen(self, releaseOxygen: 'Callable[[], None]') -> None:
        with self.o:
            self.group.wait()
            releaseOxygen()