APP?=eta

clean:
	rm -f ${APP}

build: clean
	go build -o ${APP}

run: build
	./${APP}

cars:
	curl 'https://dev-api.wheely.com/fake-eta/cars?lat=55.752992&lng=37.618333&limit=3'

predict:
	curl -X POST -H 'Content-Type: application/json' -d '{"target": {"lat": 55.752992, "lng": 37.618333}, "source": [{"lat":55.7575429, "lng":37.6135117},{"lat":55.74837156167371, "lng":37.61180107665421},{"lat":55.7532706,"lng":37.6076902}]}' 'https://dev-api.wheely.com/fake-eta/predict'

test:
	go test -v -race ./...

test-cover:
	go test -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html