FROM nginx:1.13.12

# nodeのインストール
RUN apt-get update
RUN apt-get install -y curl
RUN apt-get install -my wget gnupg
RUN curl -sL https://deb.nodesource.com/setup_9.x | bash -
RUN apt-get install -y nodejs
RUN npm i npm@latest -g

WORKDIR /app
COPY . .
RUN rm -rf node_modules/
RUN npm install
RUN npm run build

# 配置処理
WORKDIR /usr/share/nginx/html
RUN cp -r /app/dist/* .

CMD nginx -g "daemon off;"
