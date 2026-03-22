import asyncio
from datetime import datetime


async def worker(i: int) -> int:
    await asyncio.sleep(1)
    return i * 2


async def main():
    workers = []

    for i in range(5):
        workers.append(worker(i))
    results = await asyncio.gather(*workers)
    print(results)

if __name__ == '__main__':
    asyncio.run(main())
