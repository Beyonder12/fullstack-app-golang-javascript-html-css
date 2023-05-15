CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

go run main.go db.go user.go handlers.go


curl -X POST -H "Content-Type: application/json" -d '{"username": "testuser", "password": "testpassword"}' http://localhost:8000/register




curl -X POST -H "Content-Type: application/json" -d '{"username": "testuser", "password": "testpassword"}' http://localhost:8000/login
curl -X POST -H "Content-Type: application/json" -d "{\"username\": \"testuser\", \"password\": \"testpassword\"}" http://localhost:8000/login



Powershell

Invoke-WebRequest -Method POST -Uri 'http://localhost:8000/register' -ContentType 'application/json' -Body '{"username": "testuser", "password": "testpassword"}'

Invoke-WebRequest -Method POST -Uri 'http://localhost:8000/login' -ContentType 'application/json' -Body '{"username": "testuser", "password": "testpassword"}'
