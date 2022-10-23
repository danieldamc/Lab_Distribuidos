HOST = $(shell hostname)
namenode:
ifeq ($(HOST),dist149)
	go run NameNode_01/main.go
else
	echo "Maquina virtual incorrecta para el namenode"
endif


datanode:
ifeq ($(HOST),dist150)
	go run DataNode_01/main.go
endif
ifeq ($(HOST),dist151)
	go run DataNode_01/main.go
endif
ifeq ($(HOST),dist152)
	go run DataNode_01/main.go
else
	echo "maquina virtual incorrecta para el datanode"
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
