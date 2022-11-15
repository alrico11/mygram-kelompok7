# Hacktiv Final Project 2
My Gram is a rest api that functions to manage users, save a photo or make comments on other people's photos, etc

# Installation
Requires [Golang](https://go.dev/dl/) and [MySQL](https://dev.mysql.com/downloads/installer/)

Edit .env first to sql configuration and server port

- **Clone repository**
```
git clone https://github.com/alrico11/mygram-kelompok7.git
```
- **Change directory**
```
cd Final-Project-2-Hacktiv8-KampusMerdeka
```
- **Run "main.go" file**
```
go run main.go
```

# Endpoint
## User
**`POST`**	  - https://mygram-kelompok7-production.up.railway.app/users/register	      - **Create User Account**

**`POST`**	  - https://mygram-kelompok7-production.up.railway.app/users/login	          - **Login User**

**`PUT`**	    - https://mygram-kelompok7-production.up.railway.app/users/		            - **Update Data User**

**`DELETE`**	- https://mygram-kelompok7-production.up.railway.app/users/		            - **Delete User Account**

## Photos
**`POST`**  	- https://mygram-kelompok7-production.up.railway.app/photos/		            - **Post a Photo**

**`GET`**	    - https://mygram-kelompok7-production.up.railway.app/photos/		            - **Get All Photo**

**`PUT`**	    - https://mygram-kelompok7-production.up.railway.app/photos/:photoId	      - **Update Photo Posted**

**`DELETE`**	- https://mygram-kelompok7-production.up.railway.app/photos/:photoId	      - **Delete Photo Posted**

## Comments

**`POST`**  	- https://mygram-kelompok7-production.up.railway.app/comments/		            - **Post a Comment**

**`GET`**	    - https://mygram-kelompok7-production.up.railway.app/comments/		            - **Get Comment User**

**`PUT`**	    - https://mygram-kelompok7-production.up.railway.app/comments/:commentId	    - **Update User Comment Posted**

**`DELETE`**	- https://mygram-kelompok7-production.up.railway.app/comments/:commentId	    - **Delete User Comment Posted**

## Social Media

**`POST`**	  - https://mygram-kelompok7-production.up.railway.app/socialmedias/			          - **Post a Social Media info**

**`GET`**	    - https://mygram-kelompok7-production.up.railway.app/socialmedias/			          - **Get User Social Media**

**`PUT`**	    - https://mygram-kelompok7-production.up.railway.app/socialmedias/:socialMediaId	- **Update Social Media info User**

**`DELETE`**	- https://mygram-kelompok7-production.up.railway.app/socialmedias/:socialMediaId	- **Delete Social Media info User**

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
- ``Endpoint`` : Deploy API to Railway App
- ``Helper`` : Generate Password, Verify Password, jwt, dan Validator

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
- ``Documentation``	: Add README.md
