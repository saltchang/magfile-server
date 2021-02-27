if [ "$( docker container inspect -f '{{.State.Status}}' postgres_magfile )" == "running" ];
then
    docker stop postgres_magfile;
fi

docker rm postgres_magfile;
make run-postgres;
