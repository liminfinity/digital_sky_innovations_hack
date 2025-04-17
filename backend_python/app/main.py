from fastapi import FastAPI
from app.routers import auth_router, pid_router
from app.db import init_db
from contextlib import asynccontextmanager


@asynccontextmanager
async def lifespan(app: FastAPI):
    init_db()
    yield


app = FastAPI(root_path="/api/v1", lifespan=lifespan)


app.include_router(auth_router)
app.include_router(pid_router)
