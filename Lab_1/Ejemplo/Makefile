central:
	go run Central/main.go

laboratorio:
	HOSTNAME_S := $(shell hostname)
	ifeq ($(HOSTNAME_S),dist149)
		go run Pripyat/lab.go
	endif