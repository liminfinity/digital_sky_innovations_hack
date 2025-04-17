from app.schemas.pid import GetPidsResponse, PID, SavePidsDto
from app.core.settings import PID_XML_PATH
import xmltodict as xml
from typing import Any


class PidService:
    async def get_pids(self) -> GetPidsResponse:
        with open(PID_XML_PATH, encoding="utf-8") as f:
            pids_dict = xml.parse(f.read())

        pid_structure: dict = pids_dict.get("fsigmodule_structure", {})
        pid_modules: list[dict] = pid_structure.get("fsigmodule", [])

        pids = [
            self.__parse_pid_module(module)
            for module in pid_modules
            if module.get("@type") == "pid"
        ]

        return GetPidsResponse(data=pids)

    def __parse_pid_module(self, module: dict[str, Any]) -> PID:
        raw_params: list[dict] = module.get("param", [])

        formatted_params = {
            param.get("@name"): param.get("#text") for param in raw_params
        }

        return PID(name=module.get("@name"), **formatted_params)

    async def save_pids(self, pids_dto: SavePidsDto) -> None:
        with open("xml_test.xml", "w", encoding="utf-8") as f:
            f.write(xml.unparse(pids_dto.model_dump()))
