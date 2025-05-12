# quack-stack
An overkill stack for managing read operations on ducks. If you ever needed an API that returns ducks this is it.

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
