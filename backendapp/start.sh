# postgresの設定
# 起動待ち
until `psql -h db -U postgres -c '\q'`; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"

#初回であればDBが存在しないため作成する
# データベース一覧を取得し、データベース名のみを出力、その後、空白削除、不要行削除をへて、grep -w(完全一致)で対象のデータベースが存在するか確認している
COMMAND=$(psql -h db -U postgres -l |awk -F'|' '{print $1}' | tr -d ' '| grep -v -e '^\s*$' -e '^Name' -e '^Listofdatabases' -e '^-----------+'| grep -w 'app')
if [ ${COMMAND} ] ; then
  echo 'プロダクションのデータベースは存在します。'
  goose -env production up
else
  echo 'DBが存在しないため、新しくDBを作成します。'
  psql -h db -U postgres -c 'create database app;'
  goose -env production up
fi

# 起動
echo 'create backend service start!'
backendapp
