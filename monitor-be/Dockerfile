FROM alpine:latest

RUN apk add --update nodejs npm wget git sed
RUN node -v

WORKDIR "/opt"

RUN git clone https://github.com/cubedro/eth-net-intelligence-api.git

WORKDIR "/opt/eth-net-intelligence-api"

RUN sed -i 's/var interv = {};/var interv = [];/' lib/node.js
RUN sed -i 's/results = falses;/results = [];/' lib/node.js
RUN npm install -g pm2
ADD ./app.json ./app.json
RUN npm install
CMD exec pm2-runtime start app.json

