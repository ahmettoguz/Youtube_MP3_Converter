technologies
vue
nodejs
ajax
---
dependencies nodejs
---
go backend and frontned file and download dependencies

npm i

edit .env.txt file as .env and configure host name of the backend.

start backend with npm run start

start frontend with npm run serve

---

line 69 node_modules\yt-converter\src\utils\convertAudio.js:
        const pathname = path.resolve(process.cwd(), directoryDownload, `${title}.mp3`)
convert that line as following:
        const pathname = path.resolve(process.cwd(), directoryDownload, `a.mp3`)
---
cd ~

git clone https://github.com/ahmettoguz/Youtube_MP3_Converter


# Backend ssl configs
cd ~/Youtube_MP3_Converter/backend/src/keys

bash placeKeys.sh

cd ~/Youtube_MP3_Converter/
---

# Frontend connection configs
cd ~/Youtube_MP3_Converter/frontend

mv .env.txt .env

nano .env
---

# nginx ssl configs

cd ~/Youtube_MP3_Converter/nginx/ahmetproje.com.tr/keys

bash placeKeys.sh

---
# Running Deployment

cd ~/Youtube_MP3_Converter/

docker-compose down

git pull

docker-compose up -d --build

docker ps -a

docker exec -it nginx-c /bin/bash

bash link.sh

exit

docker ps -a

docker logs nginx-c