# Prophet Challenge

## Requirements

1. Git client
2. Docker
3. docker-compose
4. Golang

## Install and Run
### How to start server
Run bash scripts:
    
```
./start.sh
```

### How to run test
Run unit test with generated report

```
./scripts/test.sh
```

Note: To create mock objects for interfaces, run:

```
./scripts/gen_mocks.sh
```
### Interact with the APIs

__Place a wager__

```curl -X POST  -d '{"total_wager_value":13,"odds":2212,"selling_percentage":50,"selling_price":122.23}' localhost:8080/wagers```

__Buy a wager__

```curl -X POST  -d '{"buying_price": 1.23}' localhost:8080/buy/1```

__Get list of wagers__

```curl "localhost:8080/wagers?page=1&limit=1"```