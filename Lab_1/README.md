# Laboratorio 1 - Sistemas Distribuidos

## Integrantes
Daniel Martinez Castro 201973508-k  
Tomas Nieto Guerrero 201973519-2  

## Ejecución
Antes de ejecutar los programas se debe acceder a las máquinas virtuales que para nuestro grupo corresponden a:   
|dns|password|
|:-:|:-:|
|dist149.inf.santiago.usm.cl|D8JuLL2N|
|dist150.inf.santiago.usm.cl|kXG22pgU|
|dist151.inf.santiago.usm.cl|nWRrdWtc|
|dist152.inf.santiago.usm.cl|rPBD5Lhy|  

Ya ingresado en las VM el siguiente comando para llegar a la carpeta raíz de nuestro Laboratorio 1:

```
cd Lab_1/
```

Dentro de esta ubicación se pueden hacer los siguientes comandos:

Para ejecutar la central:
```bash
make central
```
Para ejecutar un laboratorio:
```bash
make laboratorio
```

En la máquina virtual dist149 se deben usar ambos comandos ya que esta aloja la central y un laboratorio (Pripyat), y en el caso de las otras máquinas virtuales solo se debe usar en comando para crear un laboratorio (Kampala, Pohang, Renca).