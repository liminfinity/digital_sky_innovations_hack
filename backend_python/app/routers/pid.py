from fastapi import APIRouter, Depends, status
from fastapi.responses import JSONResponse
from app.schemas.pid import GetPidsResponse, SavePidsDto
from app.services import PidService
from app.dependencies import get_pid_service

router = APIRouter(prefix="/pids", tags=["Pids"])


@router.get("", response_model=GetPidsResponse)
async def get_pids(service: PidService = Depends(get_pid_service)):
    pids = await service.get_pids()
    return pids


@router.patch("")
async def save_pids(
    pids_dto: SavePidsDto, service: PidService = Depends(get_pid_service)
):
    await service.save_pids(pids_dto)
    return JSONResponse(None, status_code=status.HTTP_204_NO_CONTENT)
