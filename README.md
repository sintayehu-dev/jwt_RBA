# Go JWT Authentication System 🔐

[![Go Version](https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](LICENSE)
[![Framework](https://img.shields.io/badge/Framework-Gin-00ADD8?style=for-the-badge&logo=go)](https://github.com/gin-gonic/gin)
[![Database](https://img.shields.io/badge/Database-PostgreSQL-336791?style=for-the-badge&logo=postgresql)](https://www.postgresql.org/)

A production-ready JWT authentication system built with Go, Gin, and PostgreSQL. Features role-based access control, secure token management, and clean architecture.

## ✨ Features

- 🔒 **Secure Authentication** - Email/password login with JWT tokens
- 🔄 **Token Refresh** - Built-in token refresh mechanism
- 👮 **Role-Based Access** - Admin and User role separation
- 🛡️ **Middleware Protection** - Secure routes with authentication middleware
- 🧩 **Clean Architecture** - Organized codebase with separation of concerns
- 🗄️ **PostgreSQL** - Reliable database with GORM ORM
- ⚡ **High Performance** - Built with Go for exceptional speed

## 📋 Prerequisites

- Go 1.16+
- PostgreSQL
- Git

## 🚀 Quick Start

### Clone the Repository
```bash
git clone https://github.com/yourusername/go_jwt_auth.git
cd go_jwt_auth
```

### Environment Setup
Create a `.env` file in the root directory:
```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=go_jwt_auth
DB_PORT=5432
PORT=9000
SECRET_KEY=your_secret_key_here
```

### Install Dependencies
```bash
go mod download
```

### Run the Application
```bash
go run main.go
```



### Authentication Flow
1. User registers with email/password
2. User logs in and receives JWT token
3. Token is included in Authorization header for subsequent requests
4. Token is validated by middleware
5. Access is granted based on user role

### User Roles
- **ADMIN**: Full access to all endpoints
- **USER**: Limited access to own resources

## 🏗️ Project Structure
```
├── controllers/ - Request handlers
├── databases/ - Database connection
├── helpers/ - Utility functions
├── middleware/ - Request interceptors
├── models/ - Data structures
├── routes/ - API endpoint definitions
└── main.go - Application entry point
```

## 🤝 Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License
This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgements
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [JWT Go](https://github.com/golang-jwt/jwt)

---


