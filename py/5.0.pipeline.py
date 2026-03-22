import asyncio
from random import random

async def producer(queue: asyncio.Queue, output_workers_num: int):
    for i in range(10):
        await asyncio.sleep(random())
        await queue.put(i)

    for i in range(output_workers_num):
        await queue.put(None)


async def filter_stage(in_q: asyncio.Queue, out_q: asyncio.Queue, input_workers: int, output_workers: int):
    none_count = 0
    while True:
        value = await in_q.get()
        if value is None:
            in_q.task_done()
            none_count += 1
            if none_count == input_workers:
                for _ in range(output_workers):
                    await out_q.put(None)
                return
            continue
        if value % 2 == 0:
            await asyncio.sleep(random())
            await out_q.put(value)
        in_q.task_done()


async def square_stage(in_q: asyncio.Queue, out_q: asyncio.Queue, input_workers: int, output_workers: int):
    none_count = 0
    while True:
        value = await in_q.get()
        if value is None:
            in_q.task_done()
            none_count += 1
            if none_count == input_workers:
                for _ in range(output_workers):
                    await out_q.put(None)
                return
            continue
        await asyncio.sleep(random())
        await out_q.put(value**2)
        in_q.task_done()


async def consumer(in_q: asyncio.Queue, out_q: asyncio.Queue):
    result = []
    while True:
        value = await in_q.get()
        if value is None:
            in_q.task_done()
            await out_q.put(None)
            break
        result.append(value)
        in_q.task_done()
        await out_q.put(value)
    print(result)


async def get_results(in_q: asyncio.Queue, workers_num: int):
    result = []
    workers_finished = 0
    while True:
        value = await in_q.get()
        if value is None:
            in_q.task_done()
            workers_finished += 1
            if workers_finished == workers_num:
                print(result)
                return result
        else:
            result.append(value)
            in_q.task_done()


async def main():
    producer_workers_num = 1
    filter_workers_num = 3
    squares_workers_num = 2
    consumer_workers_num = 4

    producer_queue = asyncio.Queue()
    filter_queue = asyncio.Queue()
    square_queue = asyncio.Queue()
    result_queue = asyncio.Queue()

    producers = []
    for _ in range(producer_workers_num):
        producers.append(producer(producer_queue, filter_workers_num))

    filters = []
    for _ in range(filter_workers_num):
        filters.append(filter_stage(producer_queue, filter_queue, producer_workers_num, squares_workers_num))

    squares = []
    for _ in range(squares_workers_num):
        squares.append(square_stage(filter_queue, square_queue, squares_workers_num, consumer_workers_num))

    workers = []
    for _ in range(consumer_workers_num):
        workers.append(consumer(square_queue, result_queue))

    results = [get_results(result_queue, consumer_workers_num)]

    await asyncio.gather(*filters, *squares, *workers, *producers, *results)

if __name__ == '__main__':
    asyncio.run(main())