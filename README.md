我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

# Fragrance

![hero](https://res.cloudinary.com/ichtrojan/image/upload/c_scale,w_1233/v1582069174/Screenshot_2020-02-18_at_00.59.01_hx9giw.png)

## Introduction

This is a Fragrance showcase built with Go. We hope this becomes an inspiration to you; feel free to use components of this codebase in your future projects. We built this to make ourselves happy, We hope you're happy looking at this project right now and you're making other people happy. ❤️

## Requirements

* [Go](https://golang.org) -  v1.11 above
* [MySQL](https://mysql.com) - v5.7 above

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

## Performance

![Performance](https://res.cloudinary.com/ichtrojan/image/upload/v1582076107/Screenshot_2020-02-19_at_02.34.44_ch4bpu.png)

## Contributors

|   Contributor Name	| Role  	|  Tool 	| Language(s)  	|
|---	|---	|---	|---	|
|   [Ifeoluwasimi Olusola](https://twitter.com/o_ifeoluwasimi)	|  Designer 	|   [Adobe XD](https://www.adobe.com/products/xd.html)	|   ---	|
|   [Muheez Jimoh](https://twitter.com/Kng_maaj)	|  Developer 	|   [VSCode](https://code.visualstudio.com)	|  HTML, CSS, JavaScript & Go 	|
|  [Michael Trojan Okoh](https://twitter.com/ichtrojan) 	|   Developer	|   [Goland](https://www.jetbrains.com/go/)	|   HTML, CSS & Go	|
