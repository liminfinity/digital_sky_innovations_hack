from fastapi import APIRouter, Depends
from app.schemas.pid import GetPidsResponse
from app.services import PidService
from app.dependencies import get_pid_service

router = APIRouter(prefix="/pids", tags=["Pids"])


@router.get("", response_model=GetPidsResponse)
async def get_pids(service: PidService = Depends(get_pid_service)):
    pids = await service.get_pids()
    return pids
