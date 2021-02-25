if [ "$( docker container inspect -f '{{.State.Status}}' postgres13.2 )" == "running" ];
then
    docker stop postgres13.2;
fi

docker rm postgres13.2;
make run-postgres;
