- Gator

A small RSS aggregation CLI written in Go.

Requirements
• Go installed
• PostgreSQL running locally

Install
go install

Make sure ~/go/bin is in your PATH.

- Setup
  1.  Create a Postgres database (e.g. gator)
  2.  Run:
      gator reset

This creates the config file and prepares the database.

Usage

- A few basic commands:
  gator register <username>
  gator login <username>
  gator addfeed "<name>" "<url>"
  gator follow "<url>"
  gator following
  gator feeds
