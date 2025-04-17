from pydantic import BaseModel


class PID(BaseModel):
    name: str
    Kp: float
    Ki: float
    Kd: float
    integral_min: float
    integral_max: float
    inp_rise_deriative: float
    inp_fall_deriative: float
    min: float
    max: float
    preset_allowed_at_low: float
    preset_allowed_at_high: float


class GetPidsResponse(BaseModel):
    data: list[PID]
