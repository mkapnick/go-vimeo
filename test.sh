for ((i=1;i<=100;i++)); do curl -w "@curl-format.txt" -o /dev/null -s http://localhost:4000/serve\?s\=http://storage.googleapis.com/vimeo-test/work-at-vimeo.mp4\&range\=0-$i; done
