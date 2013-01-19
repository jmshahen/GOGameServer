## Building the server ##
First you must set the GOPATH to the directory you have downloaded the git 
repo from.

<ol>
<li>Clone the repo to the folder
"C:\Users\Jon\Downloads\GOGameServer"</li>
<li>open a command prompt (WIN+R then "cmd" and press ENTER)</li>
<li>run the command: "set GOPATH=C:\Users\Jon\Downloads\GOGameServer" 
	(and replace the path with your directory)</li>
<li>Go to: "C:\Users\Jon\Downloads\GOGameServer\src\gameserver_server_example"
</li>
<li>run the command "go install"; this will create the executable in the folder
"C:\Users\Jon\Downloads\GOGameServer\bin\"</li>
<li>Do the same for the example client: 
	"C:\Users\Jon\Downloads\GOGameServer\src\gameserver_client_example"</li>
</ol>

## Running the example Server ##
Once there server and client are built open two command prompts and first run 
the server "gameserver_server_example.exe" and then in the other command prompt
run the client gameserver_client_example.exe" the output should be:

<pre>
C:\Jon\my documents\Go\GOGameServer\bin>gameserver_server_example.exe
GO Game Server 0.0.4 by Jonathan Shahen 2013
listening at 0.0.0.0:9989
connection from 127.0.0.1:58008
conn 127.0.0.1:58008 said 20 Random Number: 4583|


---

C:\Users\Jon\Downloads\GOGameServer\bin>gameserver_client_example.exe
GO Game Server Client Example 0.0.4 by Jonathan Shahen 2013
Message: Random Number: 4583|
in: 20
Response: Random Number: 4583|
Your Message: 
</pre>
