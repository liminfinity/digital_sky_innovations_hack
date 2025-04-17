from sqlite3 import connect, Connection, Row
from app.core.config import DATABASE_URL
from app.core.security import hash_password


def get_connection() -> Connection:
    conn = connect(DATABASE_URL)
    conn.row_factory = Row
    return conn


def init_db() -> None:
    with get_connection() as conn:

        conn.execute(
            """
            CREATE TABLE IF NOT EXISTS users (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                username TEXT UNIQUE NOT NULL,
                password TEXT NOT NULL
            )
            """
        )

        users = conn.execute("SELECT username FROM users").fetchall()
        existing_users = {user[0] for user in users}

        if "user1" not in existing_users:
            conn.execute(
                "INSERT INTO users (username, password) VALUES (?, ?)",
                ("user1", hash_password("password1")),
            )

        if "user2" not in existing_users:
            conn.execute(
                "INSERT INTO users (username, password) VALUES (?, ?)",
                ("user2", hash_password("password")),
            )
