from app.schemas.pid import (
    GetPidsResponse,
    PID,
    SavePidsDto,
    PIDFile,
    SavePidsResponse,
    GetPidsData,
    GetPidsByIdResponse,
)
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
        changes: list[PIDFile] = []

        with get_connection() as conn:
            rows = conn.execute(
                "SELECT id, filename, created_at FROM pids ORDER BY created_at DESC"
            ).fetchall()

        if rows:
            latest_file = rows[0]["filename"]
            candidate_path = PIDS_PATH / latest_file
            if candidate_path.exists():
                latest_file_path = candidate_path

            changes = [
                PIDFile(
                    id=row["id"], filename=row["filename"], created_at=row["created_at"]
                )
                for row in rows
            ]

        if latest_file_path is None:
            latest_file_path = PID_ORIGIN_PATH

        with open(latest_file_path, encoding="utf-8") as f:
            json_data = json.load(f)

        pids = [PID(**item) for item in json_data.get("data", [])]

        return GetPidsResponse(data=GetPidsData(pids=pids, changes=changes))

    async def get_pids_by_id(self, pid_id: int) -> GetPidsByIdResponse:
        json_data: dict[str, Any] = {}
        selected_file_path: Path | None = None
        changes: list[PIDFile] = []

        with get_connection() as conn:
            rows = conn.execute(
                "SELECT id, filename, created_at FROM pids ORDER BY created_at DESC"
            ).fetchall()

            changes = [
                PIDFile(
                    id=row["id"], filename=row["filename"], created_at=row["created_at"]
                )
                for row in rows
            ]

            selected_row = next((row for row in rows if row["id"] == pid_id), None)

            if selected_row:
                candidate_path = PIDS_PATH / selected_row["filename"]
                if candidate_path.exists():
                    selected_file_path = candidate_path

            if selected_file_path is None:
                selected_file_path = PID_ORIGIN_PATH

            with open(selected_file_path, encoding="utf-8") as f:
                json_data = json.load(f)

            pids = [PID(**item) for item in json_data.get("data", [])]

            return GetPidsByIdResponse(data=GetPidsData(pids=pids, changes=changes))

    async def save_pids(self, pids_dto: SavePidsDto) -> SavePidsResponse:
        time = datetime.now()
        timestamp = int(time.timestamp())
        created_at = time.isoformat()

        filename = f"pid_{timestamp}.json"

        with open(PIDS_PATH / filename, "w", encoding="utf-8") as f:
            json.dump(pids_dto.model_dump(), f, indent=2)

        with get_connection() as conn:
            cursor = conn.execute(
                "INSERT INTO pids (filename, created_at) VALUES (?, ?)",
                (filename, created_at),
            )
            pid_id = cursor.lastrowid

        update_pid_xml(pids_dto)

        pidfile = PIDFile(id=pid_id, filename=filename, created_at=created_at)

        return SavePidsResponse(**pidfile.model_dump())
