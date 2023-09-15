

# Problems with swagger 


Swagger and Swaggo defeated  me for now. Unfortunately you can’t update cookie token like in the picture below
![image](https://github.com/adam-pawelek/go_exercise/assets/45467141/be591687-3c8f-4083-a847-ab3921e65d5c)

## How to use token cookie authorization?.
1. Create user if you don’t have one 
2. Use login request (swagger/postman will automatically save cookie)
3. Use validate request to check if you are log in 

# App info 

### Database 
You can connect to website via adminer app (localhost:8080)

### Swagger 
http://localhost:4000/swagger/index.html#

### App
To run the application, all you need to do is enter this command (the .env file is in the repository)

'''
docker compose -f docker-compose-qa.yaml up
'''

<br>
<br>
<br>
<br>
<br>
<br>
<br>




# Commands for me 

Occurred problems swago 
Problems with swag init 
export PATH=$(go env GOPATH)/bin:$PATH

---

CLI commands 
**Swagger**
```
swag init
```

**Docker** 
```
docker compose -f docker-compose-dev.yaml up  -> setting up only databse
docker compose -f docker-compose-prod.yaml up -> setting up whole application 
```

