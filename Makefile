run:
	go build -o TCPChat main.go

clean:
	rm ./TCPChat

removelog:
	> ./assets/logs.txt

ip:
	@echo "Your local IP address is:"
	@ip -o -4 addr list | grep -v docker | grep 'scope global' | awk '{print $$4}' | cut -d/ -f1