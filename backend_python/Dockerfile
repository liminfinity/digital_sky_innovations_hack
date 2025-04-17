FROM python:3.13-slim

RUN apt-get update && apt-get install -y \
    curl \
    build-essential \
    make \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

RUN curl -sSL https://install.python-poetry.org | python3 - && \
    ln -s /root/.local/bin/poetry /usr/local/bin/poetry

WORKDIR /app

COPY . .

RUN poetry config virtualenvs.create false \
    && poetry install --no-interaction

RUN mkdir -p /app/pids

EXPOSE 8000

CMD ["make", "serve"]
