> # Practica 1
> ### Elian Saúl Estrada Urbina
> #### 201806838

# Go

```go
package main

import (
	"server/routes"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	app := gin.New()

	app.Use(cors.Middleware(cors.Config{
		Origins: "*",
		Methods: "GET, PUT, POST, DELETE",
		//RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		Credentials:     false,
		ValidateHeaders: false,
	}))

	//app.GET("/vehicle", controllers.GetVehicles)

	router := app.Group("")
	routes.Routes(router)

	app.Run(":4000")
}

```

# React

### Index 

```js
import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';

import "primereact/resources/themes/md-dark-indigo/theme.css";
import "primereact/resources/primereact.min.css";
import 'primeflex/primeflex.css';
import "primeicons/primeicons.css";

import "./index.css";

import App from './App';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </React.StrictMode>
);

```

### Router

### Components

# Mongo
# Docker

### DockerFile - Server

```dockerfile
FROM golang:1.17.6
WORKDIR /home/server
COPY . .
RUN go download
CMD ["go", "run", "main.go"]
```

### DockerFile - Client

```dockerfile
FROM node:16.17.0
WORKDIR /home/client
COPY .  .
RUN npm install
CMD [ "npm", "start" ]
```

### DockerCompose

```yml
version: "3"

services:
  database:
    container_name: mongo-sopes
    restart: always
    image: mongo
    ports:
      - "27017:27017"
    volumes:
      - ./data:/data/db
    networks:
      - sopes1_network

  backend:
    container_name: go-server
    restart: always
    build: ./server
    image: elian_estrada/backend_p1_201806838
    ports:
      - "4000:4000"
    depends_on:
      - database
    links:
      - database
    networks:
      - sopes1_network

  frontend:
    container_name: react-client
    restart: always
    build: ./client
    image: elian_estrada/frontend_p1_201806838
    ports:
      - "3000:3000"
    depends_on:
      - backend
    links:
      - backend
    networks:
      - sopes1_network

networks:
  sopes1_network:
    driver: bridge 
```

# KVM