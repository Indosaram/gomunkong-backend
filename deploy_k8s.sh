kubectl apply -f k8s/formatter.yaml

for LANG in "python" "golang" "java" "javascript"; do
    kubectl apply -f k8s/lang_servers-"$LANG".yaml
done
