# movie
**movie** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/movie@latest! | sudo bash
```
`username/movie` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).


### How To Run

```
ignite chain serve
```
#### Add Wallet
```
movied keys add [wallet-name]
```
#### Movie
1. Create Movie
```
movied tx movie create-movie [title] [plot] [year] [genre] [language] [is-published] [flags]
```
For more information about flags, run this following command 
```
movied tx movie create-movie --help
```
2. Update Movie
```
movied tx movie update-movie [id] [title] [plot] [year] [genre] [language] [is-published] [flags]
```
3. Delete Movie
```
movied tx movie delete-movie [id] [flags]
```

#### Review
1. Create Review
```
movied tx movie create-review [movie-id] [star] [comment] [flags]
```

2. Update Review
```
movied tx movie update-review [id] [movie-id] [star] [comment] [flags]
```

3. Delete Review
```
movied tx movie delete-review [id] [flags]
```



## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)

