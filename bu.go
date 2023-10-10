db.Migrator().CreateTable(&Users{})
// db.AutoMigrate(&Tasks{})

// Create table for `User`
db.Migrator().CreateTable(&Users{})

// Append "ENGINE=InnoDB" to the creating table SQL for `User`
db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&Users{})

// Check table for `User` exists or not
db.Migrator().HasTable(&Users{})
db.Migrator().HasTable("users")

// Drop table if exists (will ignore or delete foreign key constraints when dropping)
db.Migrator().DropTable(&Users{})
db.Migrator().DropTable("users")

// Add name field
db.Migrator().AddColumn(&Users{}, "Name")