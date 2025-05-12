# quack-stack
An **overkill** stack for managing read operations on *ducks*. If you ever needed an API that returns ducks, **this is it**.

### 🌐 Overview

Simple GET endpoint deployed to Kubernetes. Coded in Go.

---

### 🚀 Run the project:

Prerequisites: this tutorial assumes that you run on macOS and have minikube as your local Kubernetes cluster.

1. Clone the repository

```bash
git clone https://github.com/Robert076/quack-stack
```

2. Start your local minikube cluster

```bash
minikube start
```

3. Apply the Kubernetes files

```bash
kubectl apply -f kubernetes/
```

4. Port forward inside the Minikube VM

```bash
kubectl port-forward service/api-service 8080:8080
```

5. Visit [localhost:8080/ducks](http://localhost:8080/ducks)

---

### ⚙️ Project structure

```bash
.
├── 🐳 Dockerfile                          # Defines the Docker image for the application.
├── 📜 LICENSE                            
├── 📖 README.md                           
├── cmd                                     
│   └── 🦆 quack-stack                     # Main application package directory.
│       └── 🎯 main.go                     
├── go.mod                                  
├── 🔐 go.sum                              
├── internal                               
│   ├── 🗃️ database                        
│   │   └── 💾 db.go                       # Handles the database connection and querying logic.
│   └── 🦆 duck                            
│       └── 🦢 duck.go                     # Defines the Duck model.
└── kubernetes                           
    ├── 🚀 api-deployment.yml              # Kubernetes deployment configuration for the API service.
    ├── 🌐 api-service.yml                 # Kubernetes service configuration to expose the API to other pods or externally.
    ├── 🗄️ db-deployment.yml               # Kubernetes deployment configuration for the database service (PostgreSQL).
    ├── 💻 db-service.yml                  # Kubernetes service configuration to expose the database service to other pods.
    ├── 🛠️ init-db.yml                     # Kubernetes Job configuration to initialize the database schema and tables.
    └── ⚙️ quack-config.yml                # ConfigMap configuration for the application, storing environment variables.

```
