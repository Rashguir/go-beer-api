Go Beer API
===========

A simple API based on a json file containing [those data](https://data.opendatasoft.com/explore/dataset/open-beer-database%40public-us/table/).

This simple API only provides these routes as of today :

- `GET /login` to authenticate
- `GET /beers` allows you to get all the beers without pagination
- `GET /beers/:id` allows you to get one beer per ID
- `GET /beersCount` allows you to count how many beers will be returned
- `GET /db/beers` allows you to get all the beers without pagination from the database
- `GET /import` allows you to import data from the JSON file into the database

Authentication
--------------

You must authenticate with the user `toto` and password `qwerty`.

To authenticate yourself, use this route:

- `GET /login`

It's response will contain a `token`.

To call other routes in this API, use this header:

```
Token: {{the token you got from /login}}
```

Installation
------------

To have data in db, use `/import` with the token, as explained in the previous section.
