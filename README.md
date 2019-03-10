# customer-review-system
A simple application that handles signup, login of users. The users information is stored in the redis database(non-persistent).
The logged in users will be given access to the system. Already available products in the system are `ProductA`, `ProductB`, `ProductC`, `ProductD`. 
One will be able to rate those pre-defined products. The allowed rating points are `1`, `2`, `3`, `4`, `5`

# Run app locally
* Clone the repo using `https://github.com/Akhilachatlapalle/customer-review-system.git`. 
* Use `docker-compose up --d` to run this locally. 
* One can shutdownt the service using command `docker-compose down`.
* Application logs can be seen via `docker-compose logs -f app`

# Design Overview
* The design emulates the backend APIs that are related to a customer rating sites. All the details related to users and products are stored in redis database.
* Users are expected to signup to access the system. Once the user gives correct creadentials while login, user will receive a cookie in the response which will expire in 5 minutes.
* Now, the user will be given access to the products predefined in the system during startup.
* User can get the products list, rate the products and check the ratings for all products.

# API definition
The APIs, request body and expected responses are defined in `api.yaml`. Copy and paste the API contents to https://editor.swagger.io/ in browser for better readability and visbility
