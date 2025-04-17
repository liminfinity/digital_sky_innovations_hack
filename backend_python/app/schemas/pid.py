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


class PIDFile(BaseModel):
    id: int
    filename: str
    created_at: str


class GetPidsData(BaseModel):
    pids: list[PID]
    changes: list[PIDFile]


class GetPidsResponse(BaseModel):
    data: GetPidsData


class GetPidsByIdResponse(BaseModel):
    data: GetPidsData


class SavePidsDto(BaseModel):
    data: list[PID]


class SavePidsResponse(PIDFile):
    pass
