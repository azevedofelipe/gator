# CLI RSS Aggregator

## Install Requirements:
- Postgres
- Go

## Install Gator CLI:
`go install github.com/azevedofelipe/gator`

## Initial Setup 
### Create .gatorconfig in home
`vim ~/.gatorconfig.json`
It should contain the following structure:
`{"db_url":<your_database_connection_string>,"current_user_name":<current_user_name>}`

## Command
### Register User
`register <user>`

### Add RSS Feed
`addfeed <feed name> <feed url>`

### Unfollow feed
`unfollow <feed url>`

### Fetch RSS Items from followed feeds
`agg <time interval (10s)>`

### Browse Fetched Posts
browse <number of posts (10)>
