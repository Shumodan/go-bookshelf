# go-bookshelf
## About
The main purpose of the project is to be a dummy app for assessing the level of qa engineers

## Services
This sample provides 3 services: book management, account management, and master management.

### Book Management
There are the following services in the book management.

|Service Name|HTTP Method|URL|Parameter|Summary|
|:---|:---:|:---|:---|:---|
|Get Service|GET|``/api/books/[BOOK_ID]``|Book ID|Get a book data.|
|List/Search Service|GET|``/api/books?query=[KEYWORD]&page=[PAGE_NUMBER]&size=[PAGE_SIZE]``|Page, Keyword(Optional)|Get a list of books.|
|Regist Service|POST|``/api/books``|Book|Regist a book data.|
|Edit Service|PUT|``/api/books``|Book|Edit a book data.|
|Delete Service|DELETE|``/api/books``|Book|Delete a book data.|

### Account Management
There are the following services in the Account management.

|Service Name|HTTP Method|URL|Parameter|Summary|
|:---|:---:|:---|:---|:---|
|Login Service|POST|``/api/auth/login``|Session ID, User Name, Password|Session authentication with username and password.|
|Logout Service|POST|``/api/auth/logout``|Session ID|Logout a user.|
|Login Status Check Service|GET|``/api/auth/loginStatus``|Session ID|Check if the user is logged in.|
|Login Username Service|GET|``/api/auth/loginAccount``|Session ID|Get the login user's username.|

### Master Management
There are the following services in the Master management.

|Service Name|HTTP Method|URL|Parameter|Summary|
|:---|:---:|:---|:---|:---|
|Category List Service|GET|``/api/categories``|Nothing|Get a list of categories.|
|Format List Service|GET|``/api/formats``|Nothing|Get a list of formats.|

## Libraries
This sample uses the following libraries.

|Library Name|Version|
|:---|:---:|
|echo|4.3.0|
|gorm|1.21.11|
|go-playground/validator.v9|9.31.0|
|zap|1.18.1|

## License
The License of this sample is *MIT License*.
