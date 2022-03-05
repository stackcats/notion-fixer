# notion-fixer

Synchronize currency rates to Notion using [Fix.io](https://fixer.io/) and Github action

### Step 0

* Fork this repo.

### Step 1

* Create integration in Notion settings and set integration token to Github repo secret `NOTION_TOKEN`.
* Create a database and add a column called `Rates` with type `number`.
* Copy the link of the database to get the database id and set the id to Github repo secret `NOTION_DATABASE_ID`.
* Share the database to the integration which we created previously.

### Step 2

* Sign up [Fix.io](https://fixer.io/) to get a free API key.
* Set the API key to Github repo secret `FIXER_ACCESS_KEY`
