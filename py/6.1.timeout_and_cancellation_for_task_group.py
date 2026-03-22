import asyncio

async def worker(name: str, delay: int, fail: bool = False) -> None:
    try:
        print(f'worker {name} started')
        await asyncio.sleep(delay)
        if fail:
            raise Exception("worker failed")
        print(f'worker {name} done after {delay} seconds')
    except asyncio.CancelledError:
        print(f'worker {name} canceled after timeout')
        raise

async def main():
    try:
        async with asyncio.TaskGroup() as tg:
            tg.create_task(worker('1', 4))
            tg.create_task(worker("2", 1, fail=True))
            tg.create_task(worker("3", 5))
    except asyncio.TimeoutError:
        print('all tasks were canceled')

if __name__ == '__main__':
    asyncio.run(main())