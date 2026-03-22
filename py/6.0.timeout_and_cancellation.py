import asyncio

async def long_task():
    try:
        await asyncio.sleep(5)
    except asyncio.CancelledError:
        print('canceled after timeout')
        raise

async def main():
    try:
        async with asyncio.timeout(2):
            await long_task()
    except asyncio.TimeoutError:
        print('timed out')

if __name__ == '__main__':
    asyncio.run(main())