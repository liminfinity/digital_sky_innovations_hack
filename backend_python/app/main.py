from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from app.routers import auth_router, pid_router
from app.db import init_db
from contextlib import asynccontextmanager
from app.lib.pid import pid_xml_to_json


@asynccontextmanager
async def lifespan(app: FastAPI):
    init_db()
    pid_xml_to_json()
    yield


app = FastAPI(
    root_path="/api/v1", lifespan=lifespan, description="PID Api", version="1.0.0"
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)
app.include_router(auth_router)
app.include_router(pid_router)
