# **Project**

**Dependencies**
- used google library for the uuid
- using golang migration library for schema migrations
- using pqx as a postgres driver for the database/sql standard library
- using viper for configuration file management


## **Migrations**
---
</br>

**Migration steps:** </br>
1. ) Create a database with **psql**
    - [guide here](https://www.guru99.com/postgresql-create-database.html)

2. ) Create Migration </br>

    - **seq:** Marks the migration with sequence of numbers 1,2...
    - **dir:** Folder in which to place the migration
    - **ext:** The extension of the folder, **sql** in this example

    ```powershell
    migrate create -ext sql -dir db/migrations/v1                                   -seq {migration_name}
    ```

</br>

3. ) Run migrations
   1. ) Run migrations with [jq](https://stedolan.github.io/jq/) </br>
      1. For up

           ```powershell
           migrate -database 
           "$(cat config.json | jq '.Databases.PostgresConnection')"              -path db/migrations up 
           ```

      2. For down
           ```powershell
           migrate -database 
           "$(cat config.json | jq '.Databases.PostgresConnection')"              -path db/migrations down
           ```

    2. ) Run migration with raw connection string
       1. For up
           ```powershell
           migrate -database 
           "postgres://postgres:rootPassword@localhost:5432/example?sslmode=disable"                                                  -path db/migrations up 
           ```

         2. For down
              ```powershell
              migrate -database 
            "postgres://postgres:rootPassword@localhost:5432/example?sslmode=disable"                                                   -path db/migrations down
              ```

</br>

**Docs for migration:** </br>
>Migration: </br>
>https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md </br>
>Postgres docs tutorial: </br>
>  https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md


</br>
</br>
</br>


## **Logging**
---
- Use loggers Printf method

 #### **Format**

>method {methodName} | {error} </br>
```go
Example:

  logger.Printf("method GetUserById | %s", err)
```

</br>
</br>
</br>


## **Naming Conventions**
---


fk_book



















