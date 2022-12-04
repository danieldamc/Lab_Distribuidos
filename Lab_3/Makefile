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


docker-planetario:
ifeq ($(HOST),dist149) # Ports: 60500 49900 50001 49002
	echo "maquina virtual incorrecta para el planetario"
else
	docker run -it --rm -p 60500:60500 -p 49900:49900 go-app-dock:latest go run ServidorPlanetario/main.go
endif

docker-guardianes:
ifeq ($(HOST),dist151)
	docker run -it --rm go-app-dock:latest go run Guardianes/main.go
else
	echo "Maquina virtual incorrecta para guardianes"
endif

docker-broker:
ifeq ($(HOST),dist149)
	docker run -it --rm -p 49002:49002 -p 50001:50001 go-app-dock:latest go run BrokerRasputin/main.go
else
	echo "Maquina virtual incorrecta para broker"
endif

docker-vanguardia:
ifeq ($(HOST),dist150)
	docker run -it --rm go-app-dock:latest go run Vanguardia/main.go
else
	echo "Maquina virtual incorrecta para vanguardia"
endif