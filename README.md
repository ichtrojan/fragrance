# Fragrance

![hero](https://res.cloudinary.com/ichtrojan/image/upload/c_scale,w_1233/v1582069174/Screenshot_2020-02-18_at_00.59.01_hx9giw.png)

## Introduction

This is a test project to demonstrate template rendering with Go.

## Requirements

* [Go](https://golang.org) -  v1.11 above
* [MySQL](https://mysql.com) - v5.7 minimum

## Installation

>**NOTE**</br>
> * Before the fourth step, you should have created a database for this application
> * The supported database is MySQL

* Clone this repo

  ```bash
  git clone https://github.com/ichtrojan/fragrance.git
  ```

* Change directory to project directory

  ```bash
  cd fragrance
  ```

* Copy `.env` template

  ```bash
  cp .env.example .env
  ```

* Add correct database credentials to the `.env` file, credentials include:
  - `PORT`: This is the port the application will be served on
  - `DB_HOST`: This is your database host name/IP address
  - `DB_NAME`: This is the name of the database created for the application
  - `DB_NAME`: This is your database user
  - `DB_PASS`: This is your database password if any, it should be left blank if no password is configured (localhost)

* Run application

  ```bash
  go run main.go
  ```

  This will serve this application on the port you specified in the `.env` file.

## Usage

>**NOTE**</br>
> #### Dashboard credentials </br>
> **E-mail** - `cleopatra@gofragrance.xyz` </br>
> **Password** - `iloveegypt`

|  Page Name	|  Description 	|  Route 	|   Screenshot	|
|---	|---	|---	|---	|
|  Homepage 🏚	|  This is the first page, it links to all other pages and contains a nice paragraph 	|  `/` 	|   ![homepage](https://res.cloudinary.com/ichtrojan/image/upload/v1582069174/c_scale,w_1233/Screenshot_2020-02-18_at_00.59.01_hx9giw.png)	|
|   Categories ⛓	|   This page shows all the available fragrance Category, can be accessed by clicking the `Get Started` button on the homepage	|   `/categories`	|   ![categories](https://res.cloudinary.com/ichtrojan/image/upload/c_scale,w_1305/v1582069267/Screenshot_2020-02-18_at_01.01.29_cackyr.png)	|
|  Scent 👗	|  This shows the available scent available in a Category, can be accessed by clicking on any of the available categories in the categories page 	|  `/{category}` 	|   ![category](https://res.cloudinary.com/ichtrojan/image/upload/v1582069150/Screenshot_2020-02-18_at_01.01.37_qclxlc.png)	|
|   Bottle 🍶	|   This page shows the available bottle type available for a specified scent, it can be accessed by clicking on any of the scent available on the Scent page	|   `/{category}/{scent}`	|  ![bottle](https://res.cloudinary.com/ichtrojan/image/upload/v1582071428/Screenshot_2020-02-19_at_01.16.19_rrbdwj.png) 	|
|  Checkout 💳	|  This page allows you to select the bottle size and quantity of you chosen fragrance based on your previous selection. It can be accessed by clicking on any of the available bottles shown in the Bottle page	|  `/{category}/{scent}/{bottle}/checkout` 	|   ![checkout](https://res.cloudinary.com/ichtrojan/image/upload/v1582069057/Screenshot_2020-02-18_at_01.01.47_lbquea.png)	|
|   Dashboard Login	🔐|   We wanted to build a Dashboard but we didn't have the time to do so; so we made a functional Auth system	|   `/login`	|  ![login](https://res.cloudinary.com/ichtrojan/image/upload/v1582072615/Screenshot_2020-02-19_at_01.34.23_k3qauq.png) 	|
|   Dashboard Home 🏡	|   The dashboard home is the page that comes up immediately you login; the only function available is the logout functionality... I wish we had time to do more 🤕	|  `/dashboard` 	|  ![dahboard](https://res.cloudinary.com/ichtrojan/image/upload/v1582072629/Screenshot_2020-02-19_at_01.34.35_be1wpj.png) 	|

## Contributors

|   Contributor Name	| Role  	|  Tool 	| Language(s)  	|
|---	|---	|---	|---	|
|   [Ifeoluwasimi Olusola](https://twitter.com/o_ifeoluwasimi)	|  Design 	|   Adobe XD	|   nil	|
|   [Muheez Jimoh](https://twitter.com/Kng_maaj)	|  Developer 	|   VS Code	|  HTML, CSS, JavaScript & Go 	|
|  [Michael Trojan Okoh](https://twitter.com/ichtrojan) 	|   Developer	|   Goland	|   HTML & Go	|
