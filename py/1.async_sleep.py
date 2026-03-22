import asyncio


async def worker(i: int):
    print(f'worker {i} started')
    await asyncio.sleep(1)
    print(f'worker {i} finished')


async def main():
    workers = []

    for i in range(3):
        workers.append(worker(i))
    await asyncio.gather(*workers)


if __name__ == '__main__':
    asyncio.run(main())
