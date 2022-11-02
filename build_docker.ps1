# docker build -f formatter/Dockerfile -t indosaram/formatter:latest .
# docker push indosaram/formatter:latest

$LANGS = "python", "golang", "java", "javascript"
foreach ($LANG in $LANGS)
{
    docker build -f language_servers/"$LANG"/Dockerfile -t indosaram/lang-servers-"$LANG":latest .
    docker push indosaram/lang-servers-"$LANG":latest
}
