##Installation instructions
In order to run the app you need to have redis installed 

You can change the server `port` in config

Then just need to run code by following command
```sh
go run main.go
```

There are 5 workers as default but you can increase the number of worker by `max_workers` flag
```sh
go run main.go -max_workers=1
```

##Technologies

library I have used for this app are in the below
```sh
github.com/gin-contrib/cors v1.3.1
github.com/gin-gonic/gin v1.6.2
github.com/gomodule/redigo v1.8.0
golang.org/x/net v0.0.0-20200501053045-e0ff5e5a1de5
```

##Requirement

I have implemented a API for user request to be able to send one or more url(s) but for others i have not got the time.
You can see urls title in the terminal which you ran the app

##Test
I haven't got time to write test yet but you can find a postman profile in extra folder which you can import and test the api

