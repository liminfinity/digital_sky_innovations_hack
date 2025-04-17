from app.services import AuthService, PidService


def get_auth_service() -> AuthService:
    return AuthService()


def get_pid_service() -> PidService:
    return PidService()
