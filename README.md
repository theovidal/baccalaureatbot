<div align="center">
    <img src="assets/parcoursup.jpg" alt="parcoursup" max-width="150px">
    <h1>BaccalauréatBot for Telegram</h1>
    <h3>School-related stuff within easy reach</h3>
    <a href="https://t.me/BaccalaureatBot">Add the bot</a>
</div>

## 🌈 Features

- Fetch the homework and timetable for the coming days directly from PRONOTE (French software for school life)
- Do some maths (calculating, plotting...)
- Search data over educational APIs (Parcoursup...)

## 💻 Development

First, check the following requirements:

- Git, for version control
- Golang 1.15 or higher with go-modules for dependencies
- A running instance of [Redis](https://redis.io/) v5 or higher
- A running instance of [pronote-api](https://github.com/Litarvan/pronote-api) for PRONOTE-related commands

Clone the project on your local machine:

```bash
git clone https://github.com/theovidal/bacbot  # HTTP
git clone git@github.com:theovidal/bacbot      # SSH
```

Set up some environment variables described in the [.env.example file](./.env.example), either by adding them in the shell or by creating a .env file at the root of the project.

To run and test the bot, simply use `go run .`. To build an executable, use `go build .`.

## 📜 Credits

- Maintainer: [Théo Vidal](https://github.com/theovidal)
- Libraries: [check go.mod](./go.mod)
- Services: [pronote-api](https://github.com/Litarvan/pronote-api), [OpenData](https://data.enseignementsup-recherche.gouv.fr/explore/dataset/fr-esr-parcoursup/information/?timezone=Europe%2FBerlin&disjunctive.fili=true&sort=tri), [Redis](https://redis.io/)

## 🔐 License

[GNU GPL v3](./LICENSE)
