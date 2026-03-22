import asyncio
from datetime import datetime


async def worker(i: int):
    await asyncio.sleep(1)


async def main():
    workers = []

    dt1 = datetime.now()
    for i in range(3):
        workers.append(worker(i))
    await asyncio.gather(*workers)
    dt2 = datetime.now()
    print(f'Параллельное выполнение: {dt2 - dt1}')

    dt3 = datetime.now()
    for i in range(3):
        await worker(i)
    dt4 = datetime.now()
    print(f'Последовательное выполнение: {dt4 - dt3}')


if __name__ == '__main__':
    asyncio.run(main())
