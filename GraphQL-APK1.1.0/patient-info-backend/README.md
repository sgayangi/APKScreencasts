## Building the backend

To build this backend, run the following commands.
Note: Ensure you have [buildx](https://docs.docker.com/build/architecture/#install-buildx) installed.
```
docker buildx create --use
docker buildx build --platform linux/amd64,linux/arm64 -t {your-docker-username}/patient-info-backend:latest --push .
```