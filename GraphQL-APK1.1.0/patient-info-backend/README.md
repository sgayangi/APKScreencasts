## Building the backend

```
docker buildx create --use
docker buildx build --platform linux/amd64,linux/arm64 -t {your-docker-username}/patient-info-backend:latest --push .
```