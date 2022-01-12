curl -v -X POST -H "Content-Type: application/json" --data '{"application": "first-client", "param1": 2, "param2": "0.05"}' 'localhost:5000/api/savestate'
curl -v -X POST -H "Content-Type: application/json" --data '{"application": "second-client", "param1": 23, "param2": "ghi"}' 'localhost:5000/api/savestate'

curl -v -X GET -H "Content-Type: application/json" --data '{"application": "first-client"}' 'localhost:5000/api/getstate'
curl -v -X GET -H "Content-Type: application/json" --data '{"application": "second-client"}' 'localhost:5000/api/getstate'

curl -v -X POST -H "Content-Type: application/json" --data '{"application": "first-client", "param1": 2, "param2": "0.05", "version" : 1}' 'localhost:5000/api/savestate'
curl -v -X GET -H "Content-Type: application/json" --data '{"application": "first-client"}' 'localhost:5000/api/getstate'

curl -v -X POST -H "Content-Type: application/json" --data '{"application": "first-client", "param1": 22, "param2": "changed", "version" : 1}' 'localhost:5000/api/savestate'
curl -v -X GET -H "Content-Type: application/json" --data '{"application": "first-client"}' 'localhost:5000/api/getstate'

