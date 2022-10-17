# bin/sh

url=http://localhost:40400/

echo "start web server test"
sleep 1
max=1000
for i in `seq 1 $max` 
do
    curl ${url} &
done
sleep 1
echo "stop web server test"