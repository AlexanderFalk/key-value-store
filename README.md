# Key-Value Database created in GoLang

### Author: Alexander Falk
### Release version: 1.0
### Date: 05-02-2018
### Language: Go

-----

This key-value database has been created in the programming language: **Golang**.  
The idea behind the database is to store all the data in a file with a **.db**-extension; in this case - **database.db**.  
Then there will be other file called an **index-file**, which will contain the key, the offset of the data in the data-file, and the length of the data.  
The data stored in the database file is purely bytes, where the index file is clean text separated by colons.  
  
As of **05-02-2018** you'll only be able to insert and read data. The option with update and delete will be in future releases.  
When you read from the database it will be read from a hashmap, which will be loaded into memory. 
  
## Usage

### Windows

* Download the program.exe file. 
* Press WindowsButton + R. Enter: "CMD". Press Enter.
* Navigate to the destination of the downloaded file
* To insert data, write: ``` Program.exe insert --key "<INSERT KEY>" --value "<INSERT VALUE>" ```
* To read data, write:  ``` Program.exe read --get "<INSERT KEY>" ``` 


### UNIX / Linux / MacOS

* Download program.go
* Open a terminal
* Navigate to the destination of the downloaded file
* Compile the program with Go: ```go build program.go```
* If Go is not installed, follow [Google's guide](https://golang.org/doc/install#osx)
* An executable will be created
* To insert data, write: ``` Program.exe insert --key "<INSERT KEY>" --value "<INSERT VALUE>" ```
* To read data, write:  ``` Program.exe read --get "<INSERT KEY>" ``` 

----
| Issues  | Definition |
| ------------- | ------------- |
| Same key  | If you add two entries with the same key, it will only find the latest one you inserted  |
| Keep running program  | The program is very static and can't be kept alive to keep inserting and getting items  |
| Content Cell  | Content Cell  |
