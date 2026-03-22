import asyncio
from random import random

async def producer(queue: asyncio.Queue, done_event: asyncio.Event):
    for i in range(10):
        await asyncio.sleep(random())
        await queue.put(i)
    done_event.set()

async def filter_stage(
    in_q: asyncio.Queue,
    out_q: asyncio.Queue,
    in_done_event: asyncio.Event,
    out_done_event: asyncio.Event,
):
    while True:
        try:
            value = await asyncio.wait_for(in_q.get(), timeout=1)
            if value % 2 == 0:
                await asyncio.sleep(random())
                await out_q.put(value)
            in_q.task_done()
        except TimeoutError:
            if in_done_event.is_set() and in_q.empty():
                out_done_event.set()
            return

async def square_stage(
    in_q: asyncio.Queue,
    out_q: asyncio.Queue,
    in_done_event: asyncio.Event,
    out_done_event: asyncio.Event,
):
    while True:
        try:
            value = await asyncio.wait_for(in_q.get(), timeout=1)
            await asyncio.sleep(random())
            await out_q.put(value**2)
            in_q.task_done()
        except TimeoutError:
            if in_done_event.is_set() and in_q.empty():
                out_done_event.set()
                return

async def consumer(
    queue: asyncio.Queue,
    result: list[int],
    done_event: asyncio.Event,
):
    while True:
        try:
            value = await asyncio.wait_for(queue.get(), timeout=1)
            result.append(value)
            queue.task_done()
        except TimeoutError:
            if done_event.is_set() and queue.empty():
                return


async def run_stage(workers):
    await asyncio.gather(*workers, return_exceptions=True)


async def main():
    result = []

    filter_workers_num = 3
    squares_workers_num = 2
    consumer_workers_num = 4

    producer_queue = asyncio.Queue()
    filter_queue = asyncio.Queue()
    square_queue = asyncio.Queue()

    producer_done_event = asyncio.Event()
    filter_done_event = asyncio.Event()
    square_done_event = asyncio.Event()

    filters = [asyncio.create_task(filter_stage(producer_queue, filter_queue, producer_done_event, filter_done_event)) for _ in range(filter_workers_num)]
    squares = [asyncio.create_task(square_stage(filter_queue, square_queue, filter_done_event, square_done_event)) for _ in range(squares_workers_num)]
    consumers = [asyncio.create_task(consumer(square_queue, result, square_done_event)) for _ in range(consumer_workers_num)]

    await producer(producer_queue, producer_done_event)

    await producer_queue.join()
    await asyncio.gather(*filters, return_exceptions=True)

    await filter_queue.join()
    await asyncio.gather(*squares, return_exceptions=True)

    await square_queue.join()
    await asyncio.gather(*consumers, return_exceptions=True)

    print(result)

if __name__ == '__main__':
    asyncio.run(main())