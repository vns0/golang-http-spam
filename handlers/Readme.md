# SPAM BOT API Documentation

This documentation provides an overview of the SPAM BOT API, which is responsible for handling user authentication and managing user-related data. The API is built using the Gin framework in the Go programming language.

## Getting Started

To set up and run the SPAM BOT API:
    
    See in main Readme.md
    The API will start running on `http://localhost:8080`.

```go
// start api function, change port if u need
func InitAPI() {
    fmt.Print("api and client start")
    router := gin.Default()
    router.Use(cors.Default())
    
    apiRoutes := router.Group("/api/v1/auth")
    {
        apiRoutes.POST("/login", api.AuthLogin)
    }
    authedApiRoutes := router.Group("/api/v1").Use(middleware.AuthMiddleware())
	{
        authedApiRoutes.GET("/users", api.GetUsers)
        authedApiRoutes.POST("/users", api.CreateUser)
    }
    router.Run(":8080")
}
```

## API Endpoints

- **POST /api/v1/auth/login**
    - **Description:** Authenticates a user using an authorization code.
    - **Request:**
      ```json
      {
        "code": "your_authorization_code"
      }
      ```
    - **Response:**
      ```json
      {
        "access_token": "your_access_token",
        "refresh_token": "your_refresh_token"
      }
      ```

- **GET /api/v1/users**
    - **Description:** Retrieves a list of users (requires authentication).
    - **Response:**
      ```json
      {
        "users": [
          {
            "ID": 1,
            "CreatedAt": "2024-02-17T00:31:31.3624404+07:00",
            "UpdatedAt": "2024-02-17T00:31:31.3624404+07:00",
            "DeletedAt": null,
            "UserID": 1
          },
          {
            "ID": 2,
            "CreatedAt": "2024-02-17T00:31:31.3624404+07:00",
            "UpdatedAt": "2024-02-17T00:31:31.3624404+07:00",
            "DeletedAt": null,
            "UserID": 2
          }
        ]
      }
      ```

- **POST /api/v1/users**
    - **Description:** Creates a new user (requires authentication).
    - **Request:**
      ```json
      {
        "UserID": "1"
      }
      ```
    - **Response:**
      ```json
      {
        "user": {
           "ID": 1,
           "CreatedAt": "2024-02-17T00:31:31.3624404+07:00",
           "UpdatedAt": "2024-02-17T00:31:31.3624404+07:00",
           "DeletedAt": null,
           "UserID": 1
        }
      }
      ```
