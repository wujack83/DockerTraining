package database

import (
    "fmt"
    "database/sql"
    "time"
    "log"

    _ "github.com/lib/pq"
    "github.com/google/uuid"

    "gitlab.com/andersph/docker-master-class/api/internal"
)

func SqlWriter(employee internal.Employee) (err error) {

    defer retryRecover()

    // load configuration
    c := internal.LoadConfig()

    psqldsn := fmt.Sprintf("host=%s port=%s user=%s "+
                            "password=%s dbname=%s sslmode=disable",
                            internal.GetEnv("DATABASE_HOST", c.Database.Host), 
                            internal.GetEnv("DATABASE_PORT", c.Database.Port), 
                            internal.GetEnv("DATABASE_USER", c.Database.User), 
                            internal.GetEnv("DATABASE_PASSWORD", ""), 
                            internal.GetEnv("DATABASE_DB", c.Database.Db))

    db, err := sql.Open("postgres", psqldsn)
    defer db.Close()
    checkErr(err)
    sqlStatementInsert := `
    INSERT INTO hr.employees
    VALUES ($1, $2, $3, $4, $5, $6, $7);`
    _, err = db.Exec(sqlStatementInsert,
                    uuid.New(),
                    employee.Name,
                    employee.Address,
                    employee.Mail, 
                    employee.DateOfBirth,
                    employee.Department,
                    employee.JobTitle)
    checkErr(err)


    return
}

func retryRecover() { 
    if r := recover(); r != nil {
        log.Println("Error: ", r)
    }
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func SqlWriterRetry(employee internal.Employee, try int, sleepSeconds int) (message string) {

    // load configuration
    c := internal.LoadConfig()

    if internal.GetEnv("IS_STDOUT", c.Flags.Is_StdOut) != "true" {
        for i := 0; i < try; i++ {
            log.Printf("%d. attempt to write to sql database ... \n", i+1)
            err := SqlWriter(employee)
            if err == nil {
                message = fmt.Sprintf("Employee entry was successfully created on database: %+v", employee)
                break
            }
            log.Printf("Writing to database failed, next trial in %d second(s)", sleepSeconds)
            time.Sleep(time.Duration(sleepSeconds) * time.Second)
        }

        if message == "" {
            message = fmt.Sprintf("Writing to database failure for employee: %+v", employee)
        }

    } else {
        message = fmt.Sprintf("Std_Out entry: %+v", employee)
    }

    return message
}