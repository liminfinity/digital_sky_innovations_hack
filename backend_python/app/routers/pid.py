from fastapi import APIRouter, Depends
from app.schemas.pid import (
    GetPidsResponse,
    SavePidsDto,
    SavePidsResponse,
    GetPidsByIdResponse,
)
from app.services import PidService
from app.dependencies import get_pid_service

router = APIRouter(prefix="/pids", tags=["Pids"])


@router.get("", response_model=GetPidsResponse)
async def get_pids(service: PidService = Depends(get_pid_service)):
    pids = await service.get_pids()
    return pids


@router.get("/{pid_id}", response_model=GetPidsByIdResponse)
async def get_pids_by_id(pid_id: int, service: PidService = Depends(get_pid_service)):
    pids = await service.get_pids_by_id(pid_id)
    return pids


@router.patch("", response_model=SavePidsResponse)
async def save_pids(
    pids_dto: SavePidsDto,
    service: PidService = Depends(get_pid_service),
):
    change = await service.save_pids(pids_dto)
    return change
