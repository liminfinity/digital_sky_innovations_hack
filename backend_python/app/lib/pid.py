import xmltodict as xml
import json
from typing import Any
from app.core.router import PID_XML_PATH, PIDS_PATH
from app.schemas.pid import PID, SavePidsDto


def pid_xml_to_json() -> None:
    with open(PID_XML_PATH, encoding="utf-8") as f:
        pids_dict = xml.parse(f.read())

    pid_structure: dict = pids_dict.get("fsigmodule_structure", {})
    pid_modules: list[dict] = pid_structure.get("fsigmodule", [])

    pids = [
        __parse_pid_module(module)
        for module in pid_modules
        if module.get("@type") == "pid"
    ]

    json_data = {"data": [pid.model_dump() for pid in pids]}

    with open(PIDS_PATH / "pid_origin.json", "w", encoding="utf-8") as f:
        json.dump(json_data, f, indent=2, ensure_ascii=False)


def __parse_pid_module(module: dict[str, Any]) -> PID:
    raw_params: list[dict] = module.get("param", [])

    formatted_params = {param.get("@name"): param.get("#text") for param in raw_params}

    return PID(name=module.get("@name"), **formatted_params)


def update_pid_xml(pids_dto: SavePidsDto) -> None:
    with open(PID_XML_PATH, encoding="utf-8") as f:
        xml_data = xml.parse(f.read())

    structure: dict = xml_data.get("fsigmodule_structure", {})
    modules: list[dict] = structure.get("fsigmodule", [])

    pid_map = {pid.name: pid.model_dump(exclude={"name"}) for pid in pids_dto.data}

    for module in modules:
        if module.get("@type") != "pid":
            continue

        name = module.get("@name")
        if name not in pid_map:
            continue

        params: list[dict] = module.get("param", [])
        for param in params:
            key = param.get("@name")
            if key in pid_map[name]:
                param["#text"] = str(pid_map[name][key])

    xml_data["fsigmodule_structure"]["fsigmodule"] = modules

    with open(PID_XML_PATH, "w", encoding="utf-8") as f:
        f.write(xml.unparse(xml_data, pretty=True))
