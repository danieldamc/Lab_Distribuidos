central:
	go run Central/main.go

HOST = $(shell hostname)
laboratorio:
ifeq ($(HOST),dist149)
	go run Laboratorio_01/main.go
endif
ifeq ($(HOST),dist150)
	go run Laboratorio_02/main.go
endif
ifeq ($(HOST),dist151)
	go run Laboratorio_03/main.go
endif
ifeq ($(HOST),dist152)
	go run Laboratorio_04/main.go
else
	echo "Maquinas virtuales incorrectas"
endif