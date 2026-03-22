import asyncio
import random
from datetime import datetime


async def produce(queue: asyncio.Queue, workers: int):
    for i in range(10):
        await asyncio.sleep(random.random())
        await queue.put(i)
    for _ in range(workers):
        await queue.put(None)


async def consume(queue: asyncio.Queue):
    while True:
        value = await queue.get()
        if value is None:
            queue.task_done()
            return
        print(value)
        queue.task_done()


async def main():
    queue = asyncio.Queue()
    workers_num = 3
    workers = [produce(queue, workers_num)]
    for i in range(workers_num):
        workers.append(consume(queue))

    await asyncio.gather(*workers)

    await queue.join()
    print('tasks done')

if __name__ == '__main__':
    asyncio.run(main())
