import asyncio
from random import random

async def producer(queue: asyncio.Queue):
    for i in range(10):
        await asyncio.sleep(random())
        await queue.put(i)


async def filter_stage(in_q: asyncio.Queue, out_q: asyncio.Queue):
    try:
        while True:
            value = await in_q.get()
            if value % 2 == 0:
                await asyncio.sleep(random())
                await out_q.put(value)
            in_q.task_done()
    except asyncio.CancelledError:
        return


async def square_stage(in_q: asyncio.Queue, out_q: asyncio.Queue):
    try:
        while True:
            value = await in_q.get()
            await asyncio.sleep(random())
            await out_q.put(value**2)
            in_q.task_done()
    except asyncio.CancelledError:
        return

async def consumer(queue: asyncio.Queue, result: list[int]):
    try:
        while True:
            value = await queue.get()
            result.append(value)
            queue.task_done()
    except asyncio.CancelledError:
        return


async def run_stage(workers):
    for w in workers:
        w.cancel()
    await asyncio.gather(*workers, return_exceptions=True)


async def main():
    result = []

    filter_workers_num = 3
    squares_workers_num = 2
    consumer_workers_num = 4

    producer_queue = asyncio.Queue()
    filter_queue = asyncio.Queue()
    square_queue = asyncio.Queue()

    filters = [asyncio.create_task(filter_stage(producer_queue, filter_queue)) for _ in range(filter_workers_num)]
    squares = [asyncio.create_task(square_stage(filter_queue, square_queue)) for _ in range(squares_workers_num)]
    consumers = [asyncio.create_task(consumer(square_queue, result)) for _ in range(consumer_workers_num)]

    await producer(producer_queue)

    await producer_queue.join()
    await run_stage(filters)

    await filter_queue.join()
    await run_stage(squares)

    await square_queue.join()
    await run_stage(consumers)

    print(result)

if __name__ == '__main__':
    asyncio.run(main())