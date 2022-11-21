# jatis-technical-test

Dependencies:
* Mysql 
* Go

use DDL in `script/sql/ddl.sql` to create Database and Table if not exists yet

change the configuration in `config/config.yml`

to run csv loader you can use this command: 
```
go run ./cmd run-csv
```

to run api server you can use this command: 
```
go run ./cmd run-api
```
