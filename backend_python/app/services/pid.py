from app.schemas.pid import GetPidsResponse, PID, SavePidsDto
from app.core.router import PIDS_PATH, PID_ORIGIN_PATH
from app.db import get_connection
from datetime import datetime
import json
from pathlib import Path
from typing import Any
from app.lib.pid import update_pid_xml


class PidService:
    async def get_pids(self) -> GetPidsResponse:
        json_data: dict[str, Any] = {}

        latest_file_path: Path | None = None

        with get_connection() as conn:
            row = conn.execute(
                "SELECT filename FROM pids ORDER BY created_at DESC LIMIT 1"
            ).fetchone()

        if row:
            candidate_path: Path = PIDS_PATH / row["filename"]
            if candidate_path.exists():
                latest_file_path = candidate_path

        if latest_file_path is None:
            latest_file_path = PID_ORIGIN_PATH

        with open(latest_file_path, encoding="utf-8") as f:
            json_data = json.load(f)

        pids = [PID(**item) for item in json_data.get("data", [])]
        return GetPidsResponse(data=pids)

    async def save_pids(self, pids_dto: SavePidsDto) -> None:
        time = datetime.now()
        timestamp = int(time.timestamp())
        created_at = time.isoformat()

        filename = f"pid_{timestamp}.json"

        with open(PIDS_PATH / filename, "w", encoding="utf-8") as f:
            json.dump(pids_dto.model_dump(), f, indent=2)

        with get_connection() as conn:
            conn.execute(
                "INSERT INTO pids (filename, created_at) VALUES (?, ?)",
                (filename, created_at),
            )

        update_pid_xml(pids_dto)
