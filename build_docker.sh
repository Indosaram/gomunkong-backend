docker build -f formatter/Dockerfile -t indosaram/formatter:latest .
docker push indosaram/formatter:latest

for LANG in "python" "golang" "java" "javascript"; do
    docker build -f language_servers/"$LANG"/Dockerfile -t indosaram/lang-servers-"$LANG":latest .
    docker push indosaram/lang-servers-"$LANG":latest
done