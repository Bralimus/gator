# gator

## Blog aggregator
* RSS (Really Simple Syndication) feed aggregator
    Allows users and applications to access updates to websites in a standardized, computer-readable format
* Fetch the RSS feed of a website and store its content in a structured format in our database
    So it can be nicely desplayed in CLI

## Uses
1. Add RSS feeds from across the internet to be collected
2. Store the collected posts in a PostgreSQL database
3. Follow and unfollow RSS feeds that other users have added
4. View summaries of the aggregated posts in the terminal, with a link to the full post

## Info
* Use a JSON file for multiplayer functionality
    Who is currently logged in
    The connection credentials for the PostgreSQL database
* No user-based authentication for this app
    If someone has the database credentials, they can act as ANY user
    Focus on SQL, CLIs and long-running services ffor this project
* Postgresql database is "gator"
    User postgres, password: postgres
    User bralimus, password: postgres
* Connection String
    postgres://bralimus:postgres@localhost:5432/gator