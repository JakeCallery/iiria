# iiria

### To run locally for development:
change to the `tools` directory
run `docker compose -f ./redisDev.docker-compose.yaml`
in a new terminal change to the `worker` directory
create a .env file that looks like:
```bash
apikey=<api key from tomorrow.io>
latlong=30.267123, -97.743162
baseurl=https://api.tomorrow.io/v4/timelines?
localonly=false
```
run: `go run .`
in a new terminal change to the `apiServer` directory
run `go run .`
in a new terminal change to the `client` directory
run `yarn start`

### To run everything under docker compose
build each piece first:
* `cd worker`
* `go build .`
* `cd ../apiServer`
* `go build .`
* `cd ../client`
* `yarn build`

change to the `worker` directory
create a .env file that looks like:
```bash
apikey=<api key from tomorrow.io>
latlong=30.267123, -97.743162
baseurl=https://api.tomorrow.io/v4/timelines?
localonly=false
```
change the root of the project directory
run `docker compose up --build`
Open a brownser and go to `http://localhost`


