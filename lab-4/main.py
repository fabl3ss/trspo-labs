import numpy as np
import pprint
import time
from mpi4py import MPI


class Montecarlo:
    @staticmethod
    def is_in_circle(points):
        return np.square(points).sum(axis=1) <= 1

    def estimate_pi(self, comm, points_amt):
        data = None

        if rank == 0:
            cnt, reminder = divmod(points_amt, processes_amt)
            subset = [cnt * i for i in range(1, processes_amt + 1)]
            subset[-1] += reminder
            data = subset

        data = comm.scatter(data, root=0)
        start_t = time.time()
        points = np.random.rand(data, 2)
        in_circle = self.is_in_circle(points).astype(float).sum()
        pi = 4 * in_circle / data

        local_res = {
            'Process ID': comm.Get_rank(),
            'Points': data,
            'PI:': pi,
            'Time elapsed': time.time() - start_t,
        }

        global_res = comm.gather(local_res, root=0)
        if rank == 0:
            return global_res


if __name__ == '__main__':
    mpi = MPI.COMM_WORLD
    rank = mpi.Get_rank()
    processes_amt = mpi.Get_size()

    montecarlo = Montecarlo()
    results = montecarlo.estimate_pi(mpi, 1000000)

    if rank == 0:
        for result in results:
            print(50 * chr(95))
            pprint.pprint(result)
