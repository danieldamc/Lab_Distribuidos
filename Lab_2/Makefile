HOST = $(shell hostname)
namenode:
ifeq ($(HOST),dist149)
	go run NameNode/main.go
else
	echo "Maquina virtual incorrecta para el namenode"
endif


datanode:
ifeq ($(HOST),dist149)
	echo "maquina virtual incorrecta para el datanode"
else
	go run DataNode/main.go
endif

combine:
ifeq ($(HOST),dist149)
	go run Combine/main.go
else
	echo "Maquina virtual incorrecta para combine"
endif

rebeldes:
ifeq ($(HOST),dist150)
	go run Rebeldes/main.go
else
	echo "Maquina virtual incorrecta para rebeldes"
endif
