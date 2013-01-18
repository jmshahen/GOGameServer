## Building the server ##
First you must set the GOPATH to the directory you have downloaded the git 
repo from.

An example is cloning the repo to the folder 
"C:\Users\Jon\Downloads\GOGameServer", you will need to open a command prompt 
(WIN+R then "cmd" and press ENTER). Now that the command prompt is open run the 
command: "set GOPATH=C:\Users\Jon\Downloads\GOGameServer" (and replace the path 
with your directory). Nect go to the folder with the server: 
"C:\Users\Jon\Downloads\GOGameServer\src\gameserver_server_example" and run
the command "go install" this will create the executable in the folder
"C:\Users\Jon\Downloads\GOGameServer\bin\". Do the same for the example client.

## Running the example Server ##
Once there server and client are built open two command prompts and first run 
the server "gameserver_server_example.exe" and then in the other command prompt
run the client the output should be:


C:\Users\Jon\Downloads\GOGameServer\bin>gameserver_server_example.exe
listening at 0.0.0.0:9989
connection from 127.0.0.1:51991
conn 127.0.0.1:51991 said 17 Doing some stuff

conn 127.0.0.1:51991 said 6 quit


---

C:\Users\Jon\Downloads\GOGameServer\bin>gameserver_client_example.exe
Response: Doing some stuff
