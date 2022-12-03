HOST = $(shell hostname)
broker:
ifeq ($(HOST),dist149)
	go run BrokerRasputin/main.go
else
	echo "Maquina virtual incorrecta para el broker"
endif

planetario:
ifeq ($(HOST),dist149)
	echo "maquina virtual incorrecta para el servidor planetario"
else
	go run ServidorPlanetario/main.go
endif

vanguardia:
ifeq ($(HOST), dist150)
	go run Vanguardia/main.go
else
	echo "Maquina virtual incorrecta para la vanguardia"
endif

guardianes:
ifeq ($(HOST),dist151)
	go run Guardianes/main.go
else
	echo "Maquina virtual incorrecta para los guardianes"
endif
