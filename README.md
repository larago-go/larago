<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">


  <h3 align="center">LARAGO</h3>

  <p align="center">
    Structure-inspired laravel written in langue Go
    <br />
    <a href="https://github.com/larago-go/larago"><strong>Useful links and documents »</strong></a>
    <br />
    <br />
    <a href="https://github.com/gin-gonic/gin">Gin framework</a>
    ·
    <a href="https://gorm.io/docs/">ORM GORM</a>
    ·
    <a href="https://docs.mongodb.com/manual/crud/">MongoDB Crud</a>
    ·
    <a href="https://github.com/casbin/casbin">Casbin Role</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
* [Specification](#specification)
* [Getting Started](#getting-started)
  * [Installation](#installation)
* [License](#license)
* [Contact](#contact)



<!-- ABOUT THE PROJECT -->
## About The Larago

Structure-inspired laravel written in langue Go. The classic structure of the MVC with the implementation of basic authorization and a role management system
 

<!-- Specification -->
## Specification

Supports data bases mysql, postgres, sqlite, sqlserver and mongoDB. Session storage cookie, memcache and redis.

Under the hood:

* Gin framework - heart of the project (route, middleware, html template and other) 
* ORM GORM - the fantastic ORM library for Golang (Supports data bases mysql, postgres, sqlite, sqlserver and crud)
* MongoDB - MongoDB (Supports data bases mongoDB and crud)
* Casbin - An authorization library that supports access control models like ACL, RBAC, ABAC for Golang



<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.



### Installation

1. Clone the repository to the src/ folder of your directory $GOPATH

```sh
git clone https://github.com/larago-go/larago.git
```
2. Install NPM packages
```sh
npm install
```
3. Rename the file .env.example
```sh
mv .env.example .env
```

4. Сreate a database by default this is mysql(when creating use utf8mb4), you can change it by uncomplexing the necessary values in the files /config/Database.go and /config/CasbinRole.go

5. inside your project run the command
```sh
go run main.go
```

6. go to address
```sh
http://localhost:8080/
```
enjoy!


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact


Project Link: [https://github.com/larago-go/larago](https://github.com/larago-go/larago)
