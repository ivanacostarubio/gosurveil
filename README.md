# Gosurveil


Surveillance toolkit written in GO. This is a client / server toolkit that allows you to spy on other computers. Nothing to fancy, just sending information over the wire. 


### Components:


#### Client

It sends the collected data into the server.

usage: `./client --server http://127.0.0.1:8000/log/`

#### Server

It runs on port 8000 and it serves static assets on the current folder.

It accepts clients's connections and store the collected files into the ./tmp folder.
