# Laboratorio 2 - Sistemas Distribuidos

## Integrantes
Daniel Martinez Castro 201973508-k  
Tomas Nieto Guerrero 201973529-2  
Joaquin Calderon Donoso 201973571-3

## Ejecución
Antes de ejecutar los programas se debe acceder a las máquinas virtuales que para nuestro grupo corresponden a:   
|dns|password|
|:-:|:-:|
|dist149.inf.santiago.usm.cl|D8JuLL2N|
|dist150.inf.santiago.usm.cl|kXG22pgU|
|dist151.inf.santiago.usm.cl|nWRrdWtc|
|dist152.inf.santiago.usm.cl|rPBD5Lhy|  

Se deben ejecutar 2 consolas a la dist149, 2 a la dist150, 1 a la dist 151, 1 a la dist152.

Ya ingresado en las VM el siguiente comando para llegar a la carpeta raíz de nuestro Laboratorio 2:

```
cd Tarea2-Grupo38/
```

En la maquina virtual dist149 se encuentra el NameNode y Combine por lo tanto se deben ejecutar los siguientes comandos en consolas diferentes. 

Para ejecutar el namenode:
```bash
make namenode
```
Para ejecutar Combine: 
```bash
make combine
```

En la maquina virtual dist150 se encuentra un DataNode y Rebeldes por lo tanto se deben ejecutar los siguientes comandos en consolas diferentes.

Para ejecutar un Datanode:
```bash
make datanode
```
Para ejecutar Rebeldes:
```bash
make rebeldes
```

En las maquinas virtuales dist151 y dist152 se debe ejecutar un DataNode en cada una con:
```bash
make datanode
``` 
