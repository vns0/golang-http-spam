# golang-http-spam

This Go package manages a SQLite database for a spam bot application.
## Configuration

Before using the utility, configure the database settings to suit your environment:

1. Update the `AdminUserID` variable with the desired admin user ID.
2. Review the `InitDB` function to ensure the database connection settings match your environment.

```go
func InitDB() {
    // Update the database connection string as needed
    var err error
    db, err = gorm.Open(sqlite.Open("spam_bot.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&User{}, &AttackHistory{})
}
```
Notes
- Ensure that SQLite is installed and accessible in your environment.
- Modify the database connection string in InitDB to point to the desired SQLite database file.
- Customize the admin user ID in the AdminUserID variable.
- Review and adapt the code according to your specific requirements and use case.