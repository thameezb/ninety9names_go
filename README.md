# ninety9names
Implementation of SA Bodhanya's Excel 99names application  

# DEPRECATION MESSAGE
Please note that this now solely acts as an API server for -> https://github.com/thameezb/ninety9names_flutter. As such no new features will be added.
Once the flutter app is shifted to a firebase DB, this repo will be completely deprecated.

## Utilities
- [migrate](./util/migrate/migrate.go)
    - Runs [migrations](./migrations) against PSQL DB 
        - Run command `go run ./util/migrate/migrate.go` to run all migrations found 
        - env var `DATABASE_URL` must be set
- [injest](./util/injest/injest.go)
    - Injests data from [original excel file](orig.xlsx) and writes to database
        - Run command `go run ./util/injest/injest.go` to run data injest 
        - env var `DATABASE_URL` must be set
        - *NOTE* This deletes all current data from `names` table.

## How to contribute
- Please see [issues](https://github.com/thameezb/ninety9names/issues/) to assit with currently required tasks
    -   Please submit all code on distinctly named branches and submit a pull request for code to be         reviewed
    - Try to adhear to distinct commits with small PRs 
 
### Commit messages and Branch names
- All commits should have a descriptive message attached 
- Branch names should be descriptive or linked to a detailed issue 

 Please consider the master branch as stable (at all times)

## Architecture and design
Please see [issue](https://github.com/thameezb/ninety9names/issues/8) for currently proposed architecture and design. Comments/ideas/thoughts are welcome 
