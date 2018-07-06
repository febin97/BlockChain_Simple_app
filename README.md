# Blockchain Sample (KSlyrn)

This is a sample blockchain example implemented in Golang. The blockhain network is simulated using docker.

## Requirements
* Docker
* Docker Compose
* Golang (version > 1.9)
* Fabric-SDK-Go
* Make

## Usage
Run the make command in terminal. A series of steps of environment clean up, building the executable file, 
setting up the environment for the blockchain network and execution of the executable file happen.

Alternatively the following commands can be used.
* make clean
* make build
* make env-up
* make-run

If some problem occurs while 'make clean', some other virtual nodes created by docker may be present.
Then cd to the fixtures file of that code and run docker-compose down. If nothing works try restarting.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Roadmap
https://chainhero.io/2018/03/tutorial-build-blockchain-app-2/
