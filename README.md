# In progress...

## Executing with Docker

Building and running the Docker image:

    ```
    $ docker build -t my-golang-app
    $ docker run -it --rm --name my-running-app my-golang-app
    ```

Database (Just for testing, without volumes)

    ```
    $ docker run --name post -e POSTGRES_DB=sqlc -e POSTGRES_USER=sqlcgabrielteiga -e POSTGRES_PASSWORD=123abc -p 5432:5432 -d postgres:latest
    ```