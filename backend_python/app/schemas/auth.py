from .user import UserBase


class LoginUserDto(UserBase):
    password: str


class LoginUserResponse(UserBase):
    id: int
