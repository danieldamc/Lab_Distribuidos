central:
	go run Central/main.go

laboratorio:
	HOSTNAME := $(shell hostname)
	ifeq ($(HOSTNAME), dist149)
		go run Pripyat/lab.go
	endif