from fastapi import FastAPI

from email_sender.route_emailer import router

app = FastAPI()

app.include_router(router, prefix="/email", tags=["email"])

if __name__ == "__main__":
    import asyncio
    from uvicorn import Config, Server

    loop = asyncio.new_event_loop()

    config = Config(app=app, loop=loop, port=9999, host="0.0.0.0", env_file="python_backend.env")
    server = Server(config)

    loop.run_until_complete(server.serve())
    loop.run_forever()

