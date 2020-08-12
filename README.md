# microgen

A simple project template generator for https://github.com/pthethanh/micro

Install: `go install github.com/pthethanh/microgen`

Usage: `microgen -name <name> -module <module> -heroku_app_name <app_name>`

Backend App: `microgen -name usersrv -module github.com/pthethanh/usersrv -heroku_app_name usersrv`

Web App: `microgen -name usersrv -module github.com/pthethanh/usersrv -heroku_app_name usersrv -web`
