# poc_restful_api

## 🧰 Configuration

To install golang just follow the steps from website:
- https://golang.org/doc/install

To install docker and docker-compose just follow the steps from website:
- https://docs.docker.com/engine/install/
- https://docs.docker.com/compose/install/

To install postman just follow the steps from website:
- https://www.postman.com/downloads/

To import the collection and environment from repository on folder postman into your postman app please follow the steps:
- To open the Postman application, click on its icon on the taskbar.
<img src="https://github.com/miguelhbrito/stone_assignment/blob/master/images/postmanTaskbar.png" width="47" height="40">

- Click on the file tab and then click import.
<img src="https://github.com/miguelhbrito/stone_assignment/blob/master/images/postmanFileImport.png" width="255" height="296">

- Choose the import file method and then click on "Upload files".
<img src="https://github.com/miguelhbrito/stone_assignment/blob/master/images/postmanImportMethod.png" width="786" height="480">

- Choose the correct items to import and press open. Postman will automatically import the items.
<img src="https://github.com/miguelhbrito/stone_assignment/blob/master/images/postmanChooseFiles.png" width="788" height="480">

Please before start, check the file migrations/migrations.go and make the changes besides your OS.

Start database server postgresql:
``` powershell
make config-up
```
To stop database server postgresql:
``` powershell
make config-down
```

## 🛠 How to use

Start application:
``` powershell
make run-server
```

##### `/accounts` POST to create a account
##### `/accounts` GET  to list all accounts
##### `/accounts/{account_id}/balance` GET to get the balance from an account by id
##### `/login` POST to get login token
##### `/transfers` POST to create a new transfers between two accounts
##### `/transfers` GET to list all transfers

- First step is create a new account to login into system
<img src="https://github.com/miguelhbrito/stone_assignment/blob/master/images/postmanCreateAccount.png" width="620" height="365">

- Then login into system to get token auth
<img src="https://github.com/miguelhbrito/stone_assignment/blob/master/images/postmanLogin.png" width="620" height="325">

- Token is automatically saved
<img src="https://github.com/miguelhbrito/stone_assignment/blob/master/images/postmanLoginToken.png" width="595" height="322">

- After created two or more accounts, you are abble to tranfers some ammount between two accounts(the initial ammount to new account is R$100,00)
<img src="https://github.com/miguelhbrito/stone_assignment/blob/master/images/postmanCreateTransfer.png" width="617" height="319">

- Token is included in request's header
<img src="https://github.com/miguelhbrito/stone_assignment/blob/master/images/postmanCreateTransferHeader.png" width="557" height="313">


