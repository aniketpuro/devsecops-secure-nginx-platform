FROM python:3.10-slim

WORKDIR /app

# DevSecOps Best Practice: Create a non-root user
RUN groupadd -r appgroup && useradd -r -g appgroup appuser

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY main.py .

# Switch to non-root user before running the app
USER appuser

CMD ["python", "-u", "main.py"]