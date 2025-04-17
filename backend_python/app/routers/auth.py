from fastapi import APIRouter, Depends, HTTPException
from app.schemas.auth import LoginUserDto, LoginUserResponse
from app.services.auth import AuthService
from app.dependencies import get_auth_service

router = APIRouter(prefix="/auth", tags=["Auth"])


@router.post("/login", response_model=LoginUserResponse)
async def login(
    login_user_dto: LoginUserDto, service: AuthService = Depends(get_auth_service)
):
    user = await service.login(login_user_dto)
    if not user:
        raise HTTPException(status_code=404, detail="User not found")
    return user
