## Fabric-Heroes-Service

#Tiền điều kiện:
- Cài đặt docker, docker compose
- Go version 1.9.x hoặc cao hơn

# Cài đặt Fabric SDK cho go 
`go get -u github.com/hyperledger/fabric-sdk-go`
`cd $GOPATH/src/github.com/hyperledger/fabric-sdk-go`
`git checkout 614551a752802488988921a730b172dada7def1d`

#1. Check required dependencies 
    `cd $GOPATH/src/github.com/hyperledger/fabric-sdk-go `&& \`
    `make depend-install ` run command install dependencies

#2. Run test SDK
    `make`
    + If get error: ltdl.h: No such file or directory, run comand `sudo apt install libltdl-dev`

## Create network hero-services 
#1. Prepare environment
`mkdir -p $GOPATH/src/github.com/chainHero/heroes-service`
`cd $GOPATH/src/github.com/chainHero/heroes-service`

#2. Get fixtures from github,usse subversion to get folder (not get all of repository)
    - [dont need because already create on source code]
`sudo apt install -y subversion`
`cd $GOPATH/src/github.com/chainHero/heroes-service `
`svn checkout https://github.com/chainHero/heroes-service/branches/v1.0.5/fixtures`
`rm -rf fixtures/.svn`

#3. Test network
`cd $GOPATH/src/github.com/chainHero/heroes-service/fixtures`
`docker-compose up`

## Use fabric SDK Go: OPEN OTHER TERMINAL
- Edit file chainHero/heroes-service/config.yaml: Define all property of the network indluding:
    + Client: define [Organization],
    + Configuration for peer, event service and orderer timeout
    + Define Root of the MSP(Member Service Provider) directories with keys and certs.
    + BCCSP(Blockchain Cryptographic Service Provider) config for the client. Used by GO SDK.
    + Construcred [Chanel] object chanler:[chainhero:orderers:orderer.hf.chainhero.io]
- Configuration for Network Entity
    + peer0 and peer1.org1.hf.chainhero.io with [url], [eventUrl], [grpcOptions],[tlsCACerts]path and save on [fixtures/crypto-config/peerOrganizations]
    + certificateAuthorities [ca.org.hf.chainhero.io], with [url], [httpOptions:verify], [registrar:enrollId,enrollSecrect], [caName]

#1. Initialise
- file heroes-service/blockchain/setup.go : initialized a [client] communicate to a peer, a [CA] and an [Orderer]

#2. Test 
- heroes-service/main.go 
    + Config all properties of the ../blockchain/setup.go/FabricSetup struct
    + Call [setup.Initialize]
- [dep] : golang package manager (like npm in node)
- create Gopkg.toml (like package.json)
- run command
    `cd $GOPATH/src/github.com/chainHero/heroes-service`
    `dep ensure`
- build and run 
    `go bulid`
    `./heroes-service`

#3. Install & instantiate the chaincode
- Editted [chainHero/heroes-service/chaincode/main.go]

#4. Query the chaincode 
- Chaincode is plugged and  ready to answer 
- Edited file [/heroes-service/blockchain/query.go]
- Call it at [heroes-service/main.go]

#5. Change the ledger state
- Edited [chaincode/main.go] and edited [Invoke] function
- Edited [blockchain/invoke.go]
- Invoke request function call with [value] String arg and this arg will call in arg[3] in [4-arg] need to sent API request

#6. Put on MVC webapp (invoke the invoke, query function on controller.go)


