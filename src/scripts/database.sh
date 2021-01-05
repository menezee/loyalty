COMMAND=${1:-"start"}

{
  if [[ $COMMAND != "start" ]] && [[ $COMMAND != "stop" ]]
  then
    COMMAND="stop"
  fi
  eval "docker $COMMAND my-postgres"
} || {
  echo "Postgres not installed, pulling the image."
  docker run -d -p 5432:5432 --name my-postgres -e POSTGRES_PASSWORD=mysecretpassword postgres
}