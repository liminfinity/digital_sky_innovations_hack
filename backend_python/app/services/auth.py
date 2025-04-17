from app.db import get_connection
from app.schemas.auth import LoginUserDto, LoginUserResponse
from app.core.security import verify_password


class AuthService:
    async def login(self, login_user_dto: LoginUserDto) -> LoginUserResponse | None:
        username, password = login_user_dto.username, login_user_dto.password
        with get_connection() as conn:
            row = conn.execute(
                "SELECT id, username, password FROM users WHERE username = ?",
                (username,),
            ).fetchone()

        if not row:
            return None

        if not verify_password(password, row["password"]):
            return None

        user = {"id": row["id"], "username": row["username"]}

        return user
