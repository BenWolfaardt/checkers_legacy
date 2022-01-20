# Checkers
**Checkers** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

> Please note that the logic has been implemented for this project but that there are still a few things todo (pease see below).

## TODO Roadmap

- [x] Implement checkers on the Cosmos SDK
- [ ] Update README to include how to interact with the chain in the CLI.
- [ ] Write unit tests.
- [ ] Create a basic UI (even if the game is only visualised in the CLI).
- [ ] Figure out `starport network` to run the chain on multiple nodes.

## Get started

> To perform your first transaction on this chain please see the below section [Transfering Tokens with Frontend](#transfering-tokens-with-frontend) after having performed the below setup.

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

<!-- ## How to play

First create the binary for the `CLI`

```
    make build-local-linux
    cd build
    ./checkersd
```

> In this directory `./checkersd -h` will give you an overview of all of the commands.

We're interesting in `./checkersd tx checkers -h` which returns: 

```
    Available Commands:
        create-game Broadcast message createGame
        play-move   Broadcast message playMove
        reject-game Broadcast message rejectGame
```

Firstly we'll start a new game:

```
    # Checkersd tx checkers create-game [red] [black] [wager] [token] [flags]
    
    ./checkersd tx checkers create-game \
        cosmos17zx4vs4fhy7c2yv4gzq5u9uqutlzqv7xnt7035 \
        ...
        ...
        ...
        --from cosmos17zx4vs4fhy7c2yv4gzq5u9uqutlzqv7xnt7035 \
        --node http://0.0.0.0:26657 \
```

Then we'll make a move:

```
    # Checkersd tx checkers play-move [idValue] [fromX] [fromY] [toX] [toY] [flags]

    ./checkersd tx checkers play-move \
        1 \
        ...
        ...
        ...
        --from cosmos17zx4vs4fhy7c2yv4gzq5u9uqutlzqv7xnt7035 \
        --node http://0.0.0.0:26657 \
```

And then try rejecting a game: 

```
    # Checkersd tx checkers reject-game [idValue] [flags]

    ./checkersd tx checkers reject-game \
        1 \
        --from cosmos17zx4vs4fhy7c2yv4gzq5u9uqutlzqv7xnt7035 \
        --node http://0.0.0.0:26657 \
        -y
```

> Noting that the above should fail as the game has already commenced   -->

### Configure

Please see the `config.yml` file in the root directory for the configuration of the blockchain. To learn more, see the [Starport docs](https://docs.starport.network).

> Please note that the mnemonics have only been added here for reproduceability and to make life easier and that in production that this is bad practice. 

### Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

> Please note that the Frontend can be used to execute "transfer" transactions on the chain and that it isn't there to interact with the Checkers game.

#### Transfering Tokens with Frontend

* Open the Frontend on [http://localhost:8081/](http://localhost:8081/).
* In the top right-hand corner click on the dropdown and select "Add New Wallet".
* Select "Import existing wallet" and then type in a menonic from the `config.yml` file in the root directory.
* Do this for at least 2x users, and then fill in the necessary field to perform a transaction noting the user addresses in the form `cosmos17zx4vs4fhy7c2yv4gzq5u9uqutlzqv7xnt7035`.

> You have succesfully performed your first transaction on the Cosmos Chain 🥳

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

### Launch

To launch your blockchain live on multiple nodes, use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

<!-- ## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.network/BenWolfaardt/Checkers@latest! | sudo bash
```
`BenWolfaardt/Checkers` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).  -->

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/W8trcGV)
