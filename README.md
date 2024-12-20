# RSSGo

RSSGo is a simple cli tool to fetch and aggregate rss feeds built with Go.

## Get Started

- Clone repo
```bash
git clone https://github.com/mohits-git/rss-go.git
```

- Add `rssgoconfig.json` file at the root of your machine (`~/.rssgoconfig.json`) with the following structure:
```json
{
  "db_url": "postgres://<username>:<password>@localhost:5432/<db_name>?sslmode=disable"
}
```
> Note: Replace `<username>`, `<password>`, `<db_name>` with your postgres credentials.

- Build/Install the project
```bash
go build
go install
```

- Run the Project
  - Development
  ```bash
  go run . [commands...]
  ```
  - Production
  ```bash
    rss-go [commands...]
  ```

## Commands
- register
  - Register a new user
  - Usage:
  ```bash
  rss-go register <username>
  ```
- login
  - Login to the application
  - Usage:
  ```bash
  rss-go login <username>
  ```
- users
  - List all the registered users
  - Usage:
  ```bash
  rss-go users
  ```
- addfeed
  - Add a new feed to the db
  - Protected command, user should be logged in
  - URL as a identifier, should be unique
  - Usage:
  ```bash
  rss-go addfeed <feed_name> <feed_url>
  ```
- feeds
  - List all the feeds in the db
  - Usage:
  ```bash
  rss-go feeds
  ```
- follow
  - Follow a feed
  - Protected command, user should be logged in
  - URL as a identifier
  - Usage:
  ```bash
  rss-go follow <feed_url>
  ```
- unfollow
  - Unfollow a feed
  - Protected command, user should be logged in
  - URL as a identifier
  - Usage:
  ```bash
  rss-go unfollow <feed_url>
  ```
- following
  - List all the feeds followed by the user
  - Protected command, user should be logged in
  - Usage:
  ```bash
  rss-go following
  ```
- agg
  - Aggregate the feeds in the db and store the data in the db. Runs on a time interval specified by the user. Fetches one feed at a time.
  - Usage:
  ```bash
  rss-go agg <time_interval>
  ```
  > time_interval: 1s, 1m, 1h, ... etc <br/>
  > Make sure you specify the correct time interval do not DOS the servers by specifying a very low time interval. You can stop the aggregation by pressing `Ctrl + C`
- browse
  - Browse the aggregated feeds, follwed by the user
  - Protected command, user should be logged in
  - Limit defaults to 5 and offset defaults to 0
  - Ordered by published date. Recent first.
  - Usage:
  ```bash
  rss-go browse [limit] [offset]
  ```
