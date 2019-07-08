Go Beer API
===========

A simple API based on a json file containing [those data](https://data.opendatasoft.com/explore/dataset/open-beer-database%40public-us/table/).

This simple API only provides these routes as of today :

- `GET /beers` allows you to get all the beers without pagination
- `GET /beers/:id` allows you to get one beer per ID
- `GET /beersCount` allows you to count how many beers will be returned

Authentication
--------------

To authenticate yourself, use this route:

- `GET /login`

It's response will contain a `token`.

To call other routes in this API, use this header:

```
Token: {{the token you got from /login}}
```
