## go-vimeo

### How to run
- `git clone https://github.com/mkapnick/go-vimeo.git`
- install redis (used for caching): `brew install redis`
- run redis: `redis-server /usr/local/etc/redis.conf`
- run this server locally: `go run main.go`
- Test it out!

### Testing
 - `curl http://localhost:4000/serve?s\=http://storage.googleapis.com/vimeo-test/work-at-vimeo.mp4\&range\=0`
 - `curl http://localhost:4000/serve?s\=http://storage.googleapis.com/vimeo-test/work-at-vimeo.mp4\&range\=0-100`
 - `curl http://localhost:4000/serve?s\=http://storage.googleapis.com/vimeo-test/work-at-vimeo.mp4\&range\=75-175`

### Bulk Testing
 - `chmod +x ./test.sh`
 - `./test.sh`
 - ^ This will run 100 requests with relatively small byte ranges
 - Run it back to back to see the caching in action
