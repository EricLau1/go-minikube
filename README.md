Env:

    export GO111MODULE=on

Run:

    go mod init

    go build main.go


Docker:

    sudo docker build -t goapi:v1 .

    sudo docker run -it --rm -p 9000:9000 goapi:v1

Kubernetes:

    minikube start
    
    eval $(minikube docker-env)

    sudo docker build -t goapi:v1 .

    kubectl create -f app.yaml

    kubectl create -f deployment.yaml

    kubectl create -f service.yam

    Url service:
    
    minikube service list

    minikube service goapi --url

Delete images:

    kubectl delete services goapi

    kubectl delete deployments goapi
    
    kubectl delete pods goapi
    
    docker rmi -f $(docker images -qa)

    docker rm -f $(docker ps -qa)