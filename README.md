# Test Api documentation

To run the project, first you need to run migrations in your postgreSql database. Just copy and paste 1_create_tables.up.sql file (in folder migrations) on your db.

To run the project in _development mode_:

    $ go run cmd/main.go

Here is two endpoints:
POST: ```/insert```

Structure must be like:

``` 
{
  "first_name": "John",
  "last_name": "Doyle",
  "addresses": [
    {
      "address": "Miami Beach, Local street 1"
    },
    {
      "address": "Miami Beach, Local street 2"
    }
  ]
 } 
 ```

GET: ```/get```

To get user by id you should in header put:

```X-User-Id: {user_id}
```