# build処理
npm run build

# 配置処理
cd /usr/share/nginx/html
cp -r /app/dist/* .

nginx -g "daemon off;"
