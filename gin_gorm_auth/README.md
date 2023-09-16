# App info 

### Database 
You can connect to database via adminer app (localhost:8080) -> Docker

### Swagger 
http://localhost:4000/swagger/index.html#  -> Docker

### App
You can run this application either with docker or kubernetes.

#### Docker 
```
docker compose -f docker-compose-qa.yaml up
```

#### Kubernetes
```
1. 
kubectl apply -f postgres-deploy.yaml 
kubectl apply -f postgres-service.yaml 

2. 
kubectl apply -f go-app-deploy.yaml
kubectl apply -f go-app-service.yaml 

3. For checking database
kubectl apply -f db-amdiner-deploy.yaml 
kubectl apply -f db-amdiner-service.yaml 
``` 

<br>

# Problems with swagger 
### Problem with kubernetes
For know you can access swagger from kubernetes deployment but you can't make request via Swagger. 


Swagger and Swaggo defeated  me for now. Unfortunately you can’t update cookie token like in the picture below
![image](https://github.com/adam-pawelek/go_exercise/assets/45467141/be591687-3c8f-4083-a847-ab3921e65d5c)

## How to use token cookie authorization?.
1. Create user if you don’t have one 
2. Use login request (swagger/postman will automatically save cookie)
3. Use validate request to check if you are log in 



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
```
export PATH=$(go env GOPATH)/bin:$PATH
```

---

CLI commands 
**Swagger**
```
swag init
```

**Docker** 
```
docker compose -f docker-compose-dev.yaml up  -> setting up only database
docker compose -f docker-compose-prod.yaml up -> setting up whole application 
```

