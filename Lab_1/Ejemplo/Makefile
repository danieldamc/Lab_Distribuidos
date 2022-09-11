central:
	go run Central/main.go

HOST = $(shell hostname)
laboratorio:
ifeq ($(HOST),dist149)
	go run Pripyat/lab.go
endif
ifeq ($(HOST),dist150)
	go run Kampala/lab.go
endif
ifeq ($(HOST),dist151)
	go run Pohang/lab.go
endif
ifeq ($(HOST),dist152)
	go run Renca/lab.go
else
	echo "Maquinas virtuales incorrectas"
endif