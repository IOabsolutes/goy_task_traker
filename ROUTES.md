# API Routes Documentation

## Base URL
```
http://localhost:8080
```

## Health Check
- `GET /health` - Check API health status

## User Management Routes

### Users CRUD
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create new user
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user
- `PATCH /api/v1/users/:id/subscription` - Update user subscription

### Authentication
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/logout` - User logout

## Task Management Routes

### Tasks CRUD
- `GET /api/v1/tasks` - Get all tasks
- `GET /api/v1/tasks/:id` - Get task by ID
- `POST /api/v1/tasks` - Create new task
- `PUT /api/v1/tasks/:id` - Update task
- `DELETE /api/v1/tasks/:id` - Delete task
- `PATCH /api/v1/tasks/:id/status` - Update task status

### User-specific Tasks
- `GET /api/v1/users/:userId/tasks` - Get all tasks for a specific user

## Example Usage

### Create a Task
```bash
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user123",
    "title": "Complete project",
    "body": "Finish the task tracker API",
    "priority": "high",
    "start_date": "2024-01-01T00:00:00Z",
    "end_date": "2024-01-31T23:59:59Z"
  }'
```

### Get All Tasks
```bash
curl http://localhost:8080/api/v1/tasks
```

### Update Task Status
```bash
curl -X PATCH http://localhost:8080/api/v1/tasks/1/status \
  -H "Content-Type: application/json" \
  -d '{"status": "success"}'
```