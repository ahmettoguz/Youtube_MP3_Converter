# API

```
cd api
```

```
go mod init github.com/ahmettoguz/Youtube_MP3_Converter
```

```
go get github.com/kkdai/youtube/v2
go mod tidy
```

```
go run cmd/app/main.go
```

docker compose up -d --build
docker rm -f youtube-api
rm -rf mp3
mkdir mp3


curl "http://localhost:8080/health"
curl "http://localhost:8080/download?url=https://youtu.be/z1HUvadz_ZY?list=LL"
curl "http://localhost:8080/download?url=https://youtu.be/FIgkb38qINA?list=LL"
