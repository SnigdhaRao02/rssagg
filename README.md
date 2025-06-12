# GoRSSAggregator

## About The Project

This project is a self-hosted RSS (Really Simple Syndication) feed aggregator built with Go. It allows users to subscribe to various RSS feeds, automatically fetches and stores their content in a database, and provides a clean API to access aggregated posts, empowering users to control their information flow outside of algorithmic feeds and ephemeral social media platforms.

### Key Features

* **API-Driven:** All interactions are handled via a RESTful API, making it easy to integrate with custom front-ends or other applications.
* **Feed Management:** Add, list, and delete RSS feed subscriptions.
* **Automatic Polling:** Periodically fetches new posts from subscribed feeds.
* **Content Storage:** Stores articles in a PostgreSQL database for persistent access.
* **User Authentication:** API Keys to secure access to user-specific feeds and posts.
