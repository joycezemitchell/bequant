# Bequant
Service that collect data from cryptocompare.com using its API and stores it in a database (PostgreSQL))

## Installation

Run the following command
```bash
git clone https://github.com/joycezemitchell/bequant.git 
go build -o bequantserver server.go
go build -o sychdata synch.go
```

## Database Configuration
Open configration.json and update the

```bash
{
    "host":"xxxx",
    "port":5432,
    "user":"xxxx",
    "password":"xxxx",
    "dbname":"bequant"
}
```


## Usage
At the moment it only supports 
- BTC
- LINK
- USD
- EUR

Grab all BTC and LINK with USD and EUR currencies
```sh
http://bequant.allyapps.com/price?fsyms=BTC,LINK&tsyms=USD,EUR
```
Grab all BTC and LINK with USD currency only
```sh
http://bequant.allyapps.com/price?fsyms=BTC,LINK&tsyms=USD
```
Grab all BTC with EUR currencies
```sh
http://bequant.allyapps.com/price?fsyms=BTC&tsyms=EUR
```

Grab all LINK with USD currency 
```sh
http://bequant.allyapps.com/price?fsyms=LINK&tsyms=USD
```

Sample Postman output

![picture](images/postman.PNG)

## Grabbing new data from cryptocompare.com
Running **synch.go** will get/collect data from cryptocompare.com using its API and stores it in a database. This need to be running in the background and run as a daemon service. 

At the moment I have stop this service as it will run every 2 minutes and will max out my test server environment.
Please let me know if you need me to start this service.

Here is a screenshot of a the daemnon configuration

![picture](images/im1.PNG)

## Database Structure
I used **PostgreSql** as a backend database and created 2 tables called **raw** and **display**

![picture](images/raw.PNG)
![picture](images/display.PNG)


Sample database content

![picture](images/rawc.PNG)
![picture](images/displayc.PNG)






