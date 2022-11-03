# moonlay-todo-api

Set Up project

1. go mod tidy
2. Create Database in postgreSQL
3. setup .env
   Example:
   HOST=localhost
   PORTDB=5432
   DBUSER=postgres
   DBPASSWORD=postgres
   DBNAME=moonlay-todo-api
4. create uploads folder (for upload pdf or txt file)
5. go run main.go or can run with make nrun (nodemon)

# API DOCUMENTATION

recomended test with postman

Get all todo
http://localhost:8080/todo?limit=2&page=1 (GET)
param:
limit
page

Get all todo with sub todo
http://localhost:8080/todo/subtodo (GET)

Get Sub todo with todo id
http://localhost:8080/subtodo (GET)
Body:
form-data:
todo_id = 1

Create Todo
http://localhost:8080/todo (POST)
Body:
Form-data:
title
description
file (file upload)

Create Subtodo
http://localhost:8080/subtodo (POST)
Body:
Form-data:
title
description
file (file upload)
todo_id

Delete Todo
http://localhost:8080/todo/1 (DELETE)

Delete Subtodo
http://localhost:8080/subtodo/5 (DELETE)

Update Todo
http://localhost:8080/todo/1 (PATCH)
Body:
Form-data:
title
description
file (file upload)

Update subtodo
http://localhost:8080/subtodo/6 (PATCH)
Body:
Form-data:
title
description
file (file upload)
