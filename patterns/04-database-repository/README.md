# Database Repository Pattern

Deep dive into the Repository Pattern for clean database access in Go.

## What is the Repository Pattern?

The Repository Pattern provides an **abstraction layer** between your business logic and data access logic.

### Problem Without Repository Pattern

```go
// ❌ Business logic mixed with SQL
func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    // SQL directly in handler!
    db.Exec("INSERT INTO users (username, email) VALUES ($1, $2)",
        user.Username, user.Email)

    json.NewEncoder(w).Encode(user)
}
```

**Problems**:
- Hard to test (needs real database)
- Can't swap database easily
- SQL scattered throughout codebase
- Violates Single Responsibility Principle

### Solution With Repository Pattern

```go
// ✅ Clean separation

// 1. Interface (What's possible)
type UserRepository interface {
    Create(ctx context.Context, user *User) error
    GetByID(ctx context.Context, id int64) (*User, error)
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id int64) error
}

// 2. Implementation (How it works)
type PostgresUserRepository struct {
    db *sql.DB
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *User) error {
    query := `
        INSERT INTO users (username, email, password)
        VALUES ($1, $2, $3)
        RETURNING id, created_at
    `
    return r.db.QueryRowContext(ctx, query,
        user.Username, user.Email, user.Password,
    ).Scan(&user.ID, &user.CreatedAt)
}

// 3. Handler (Clean business logic)
func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    // Clean! No SQL here
    if err := app.store.Users.Create(r.Context(), &user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}
```

## Benefits

### 1. Testability
```go
// Mock repository for testing
type MockUserRepository struct {
    users map[int64]*User
}

func (m *MockUserRepository) Create(ctx context.Context, user *User) error {
    user.ID = int64(len(m.users) + 1)
    m.users[user.ID] = user
    return nil
}

// Test without real database!
func TestCreateUser(t *testing.T) {
    app := &application{
        store: store.Storage{
            Users: &MockUserRepository{users: make(map[int64]*User)},
        },
    }
    // Test handler with mock
}
```

### 2. Swappable Backends
```go
// Easy to switch databases
store := store.NewPostgresStorage(db)   // Production
store := store.NewMongoStorage(client)  // Migration
store := store.NewInMemoryStorage()     // Testing
```

### 3. Centralized Data Access
```go
// All SQL queries in one place
internal/store/
├── storage.go    # Interfaces
├── users.go      # User SQL queries
├── posts.go      # Post SQL queries
└── comments.go   # Comment SQL queries
```

## Implementation Example

### storage.go - Interface Definitions
```go
package store

type Storage struct {
    Users interface {
        Create(context.Context, *User) error
        GetByID(context.Context, int64) (*User, error)
        GetByEmail(context.Context, string) (*User, error)
        Update(context.Context, *User) error
        Delete(context.Context, int64) error
    }

    Posts interface {
        Create(context.Context, *Post) error
        GetByID(context.Context, int64) (*Post, error)
        GetByUserID(context.Context, int64) ([]*Post, error)
        Update(context.Context, *Post) error
        Delete(context.Context, int64) error
    }
}

func NewPostgresStorage(db *sql.DB) Storage {
    return Storage{
        Users: &UsersStorage{db},
        Posts: &PostsStorage{db},
    }
}
```

### users.go - User Repository
```go
package store

type User struct {
    ID        int64     `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"-"` // Hidden in JSON
    CreatedAt time.Time `json:"created_at"`
}

type UsersStorage struct {
    db *sql.DB
}

func (s *UsersStorage) Create(ctx context.Context, user *User) error {
    query := `
        INSERT INTO users (username, email, password)
        VALUES ($1, $2, $3)
        RETURNING id, created_at
    `
    return s.db.QueryRowContext(ctx, query,
        user.Username, user.Email, user.Password,
    ).Scan(&user.ID, &user.CreatedAt)
}

func (s *UsersStorage) GetByID(ctx context.Context, id int64) (*User, error) {
    query := `
        SELECT id, username, email, created_at
        FROM users
        WHERE id = $1
    `
    user := &User{}
    err := s.db.QueryRowContext(ctx, query, id).Scan(
        &user.ID, &user.Username, &user.Email, &user.CreatedAt,
    )
    if err == sql.ErrNoRows {
        return nil, ErrNotFound
    }
    return user, err
}
```

## Advanced Patterns

### Context for Timeouts
```go
func (s *UsersStorage) Create(ctx context.Context, user *User) error {
    // Context carries deadline, cancellation
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    return s.db.QueryRowContext(ctx, query, ...).Scan(...)
}
```

### Transactions
```go
func (s *Storage) CreateUserWithProfile(ctx context.Context, user *User, profile *Profile) error {
    tx, err := s.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()

    // Create user
    if err := s.createUser(ctx, tx, user); err != nil {
        return err
    }

    // Create profile
    if err := s.createProfile(ctx, tx, profile); err != nil {
        return err
    }

    return tx.Commit()
}
```

### Parameterized Queries (SQL Injection Prevention)
```go
// ✅ SAFE - Parameterized
query := "SELECT * FROM users WHERE email = $1"
rows, err := db.QueryContext(ctx, query, email)

// ❌ DANGEROUS - String concatenation
query := "SELECT * FROM users WHERE email = '" + email + "'"
rows, err := db.QueryContext(ctx, query)
// Vulnerable to: ' OR '1'='1
```

## Best Practices

1. **One Repository Per Entity**: UsersStorage, PostsStorage, CommentsStorage
2. **Interface First**: Define what's needed, then implement
3. **Context Everywhere**: Enable timeouts and cancellation
4. **Return Errors**: Don't panic, return descriptive errors
5. **Use Transactions**: For multi-step operations
6. **Parameterized Queries**: Prevent SQL injection
7. **RETURNING Clause**: Get auto-generated values (ID, timestamps)

## Related Patterns

- `03-rest-api-template`: See repository pattern in action
- Active Record vs Repository: Different approaches to data access
- Unit of Work: Coordinating multiple repositories in a transaction

## Resources

- [Martin Fowler - Repository Pattern](https://martinfowler.com/eaaCatalog/repository.html)
- [Go database/sql Tutorial](https://go.dev/doc/database/querying)
- [PostgreSQL Go Driver](https://github.com/lib/pq)
