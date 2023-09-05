FROM node:18.17.1-alpine3.18

ENV NODE_ENV development

#RUN apk --no-cache add yarn python2 make gcc g++
ADD start.sh /
RUN chmod +x /start.sh

CMD ["/start.sh"]