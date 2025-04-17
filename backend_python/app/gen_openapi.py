import json
from fastapi.openapi.utils import get_openapi
from .main import app


def generate_openapi_json(output_path: str = "openapi.json") -> None:
    openapi_schema = get_openapi(
        title=app.title,
        version=app.version,
        openapi_version=app.openapi_version,
        description=app.description,
        routes=app.routes,
        tags=getattr(app, "openapi_tags", []),
        servers=getattr(app, "servers", None),
    )

    if not openapi_schema.get("servers"):
        server_url = getattr(app, "openapi_prefix", None) or getattr(
            app, "root_path", None
        )
        if server_url:
            openapi_schema["servers"] = [{"url": server_url}]

    with open(output_path, "w", encoding="utf-8") as f:
        json.dump(openapi_schema, f, indent=2)
    print(f"OpenAPI schema successfully written to '{output_path}'")


if __name__ == "__main__":
    generate_openapi_json()
