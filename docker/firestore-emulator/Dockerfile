FROM node:18.17.1-alpine3.18

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache openjdk17-jre
RUN apk add --no-cache bash

RUN npm config set registry https://registry.npm.taobao.org/
RUN npm install -g firebase-tools@12.5.2

COPY start.sh .
RUN chmod +x start.sh
COPY firebase.json .
COPY .firebaserc .

ENTRYPOINT ["./start.sh"]
