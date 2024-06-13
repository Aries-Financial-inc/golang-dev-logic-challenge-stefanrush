# Aries Financial Challenge - Stefan Rush

## Solution Instructions

1. Ensure your system has `go`, `make`, `docker`, and `docker-compose` installed. The docker dependencies are optional and only used to run a KeyDB (Redis) database to store the analysis results.

2. Clone the repository and navigate to the root directory.

3. Install the dependencies using the following command:

    ```bash
    make update
    ```

4. Start the KeyDB store using the following command (optional):

    ```bash
    docker-compose up -d
    ```

    This will start a KeyDB store running on `localhost:6379`

5. Start the server using the following command:

    ```bash
    make run
    ```

    This will start the server on `localhost:8080`

6. You should now be able to access the following endpoints on `localhost:8080`:

    - `POST /analysis`: Used to submit any number of options contracts for analysis

        - The request body should use the following format:

            ```json
            [
                 {
                    "type": "call",
                    "long_short": "long",
                    "strike_price": 100,
                    "expiration_date": "2025-12-17T00:00:00Z",
                    "bid": 10.05,
                    "ask": 12.04
                },
                 {
                    "type": "put",
                    "long_short": "short",
                    "strike_price": 100,
                    "expiration_date": "2025-12-17T00:00:00Z",
                    "bid": 10.05,
                    "ask": 12.04
                }
            ]
            ```

        - The response with contain the analysis results along with a UUID for retrieving the results on-demand from the KeyDB store

    - `GET /analysis/{contract_id}`: Used to retrieve the analysis results for a given contract ID (`{contract_id}` should be replaced with a UUID returned from `POST /analysis`)

        - Note that this endpoint will only work if the KeyDB store is running

    - `GET /analysis/{contract_id}/graph`: Used to render a line graph of the analysis results for a given contract ID (`{contract_id}` should be replaced with a UUID returned from `POST /analysis`)

        - Note that this endpoint will only work if the KeyDB store is running

## Original Problem Statement

### Objective

Design and implement a backend service for options contracts risk and reward analysis using GoLang.

### Brief

Aries Financial is looking for a Lead GoLang Developer that can create a backend service to generate a risk & reward graph for options contracts. The service should accept an input of up to four options contracts and output X & Y values for a risk & reward graph where X is the price of the underlying at the time of expiry and Y is the profit/loss at that price. It should also return the following: max profit, max loss, and all break even points.

Here are the tasks you need to accomplish:

1. **Task 1: Options Contract Model**
    - Implement an OptionsContract model with the following fields: type (call or put), strike_price, bid, ask, expiration_date, long_short

2. **Task 2: Analysis Endpoint**
    - Implement an endpoint that accepts an array of up to four options contracts and returns the following:
        - An array of X & Y values for the risk & reward graph
        - The maximum possible profit
        - The maximum possible loss
        - All break even points

3. **Task 3: Analysis Logic**
    - Implement logic to calculate the X & Y values for the risk & reward graph
        - X values should be the price of the underlying at the time of expiry
        - Y values should be the profit or loss at that price
    - Implement logic to calculate the maximum possible profit, maximum possible loss, and all break even points

4. **Task 4: Testing**
    - Write unit tests for the options contract model validation
    - Write unit tests for the analysis endpoint
    - Write integration tests that simulate a user submitting options contracts and receiving the analysis results

### Evaluation Criteria

- Correctness and completeness of the code.
- Use of GoLang idioms and best practices.
- Structure and organization of the code.
- Quality of the tests and coverage of the code.

### CodeSubmit

Please organize, design, test, and document your code as if it were
going into production - then push your changes to the master branch.

Have fun coding! 🚀

The Aries Financial Team
