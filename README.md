# Hacktiv Final Project 2
My Gram is a rest api that functions to manage users, save a photo or make comments on other people's photos, etc

# Installation
Requires [Golang](https://go.dev/dl/) and [MySQL](https://dev.mysql.com/downloads/installer/)

Config the .env first to connect into database

- **Clone repository**
```
git clone https://github.com/alrico11/mygram-kelompok7.git
```
- **Change directory**
```
cd mygram-kelompok7
```
- **Run "main.go" file**
```
go run main.go
```

# Endpoint
## User
**`POST`**	  - http://localhost:8080/users/register	      - **Create User Account**

**`POST`**	  - http://localhost:8080/users/login	          - **Login User**

**`PUT`**	    - http://localhost:8080/users/		            - **Update Data User**

**`DELETE`**	- http://localhost:8080/users/		            - **Delete User Account**

## Photos
**`POST`**  	- http://localhost:8080/photos/		            - **Post a Photo**

**`GET`**	    - http://localhost:8080/photos/		            - **Get All Photo**

**`PUT`**	    - http://localhost:8080/photos/:photoId	      - **Update Photo Posted**

**`DELETE`**	- http://localhost:8080/photos/:photoId	      - **Delete Photo Posted**

## Comments
**`POST`**  	- http://localhost:8080/comments/		            - **Post a Comment**

**`GET`**	    - http://localhost:8080/comments/		            - **Get Comment User**

**`PUT`**	    - http://localhost:8080/comments/:commentId	    - **Update User Comment Posted**

**`DELETE`**	- http://localhost:8080/comments/:commentId	    - **Delete User Comment Posted**

## Social Media
**`POST`**	  - http://localhost:8080/socialmedias/			          - **Post a Social Media info**

**`GET`**	    - http://localhost:8080/socialmedias/			          - **Get User Social Media**

**`PUT`**	    - http://localhost:8080/socialmedias/:socialMediaId	- **Update Social Media info User**

**`DELETE`**	- http://localhost:8080/socialmedias/:socialMediaId	- **Delete Social Media info User**




# Group 7
1. **[Alrico Rizki Wibowo](https://github.com/alrico11)**   - GLNG-KS04-017
2. **[Ricky Khairul Faza](https://github.com/rickyfazaa)**  - GLNG-KS04-022
3. **[Muhammad Rafid](https://github.com/mrafid01)**  - GLNG-KS04-024

## Pembagian Tugas
### Alrico Rizki Wibowo
Alrico Rizki Wibowo mengerjakan beberapa hal berikut :
- ``Endpoint`` : POST /users/register
- ``Endpoint`` : POST /users/login
- ``Endpoint`` : PUT /users
- ``Endpoint`` : DELETE /users
- ``Endpoint`` : POST /comments
- ``Endpoint`` : GET /comments
- ``Endpoint`` : PUT /comments/:commentId
- ``Endpoint`` : DELETE /comments/:commentId
- ``Helper`` : Generate Password, Verify Password, jwt, dan Validator
- ``Additional`` : Deploy API to **Railway App**

### Muhammad Rafid
Muhammad Rafid mengerjakan beberapa hal berikut :
- ``Endpoint``	: POST /photos
- ``Endpoint``	: GET /photos
- ``Endpoint``	: PUT /photos/:photoId
- ``Endpoint``	: DELETE /photos/:photoId
- ``Endpoint``	: POST /comments
- ``Endpoint``	: GET /comments
- ``Endpoint``	: PUT /comments/:commentId
- ``Endpoint``	: DELETE /comments/:commentId
- ``Helper``	: Unit Test dan .env
- ``Additional``	: help Alrico and Ricky for fixing bug

### Ricky Khairul Faza
Ricky Khairul Faza mengerjakan beberapa hal berikut :
- ``Endpoint``	: POST /socialmedias
- ``Endpoint``	: GET /socialmedias
- ``Endpoint``	: PUT /socialmedias/:socialMediaId
- ``Endpoint``	: DELETE /socialmedias/:socialMediaId
- ``Endpoint``	: POST /comments
- ``Endpoint``	: GET /comments
- ``Endpoint``	: PUT /comments/:commentId
- ``Endpoint``	: DELETE /comments/:commentId
- ``Documentation``	: Postman Collection
- ``Additional``	: Add README.md
