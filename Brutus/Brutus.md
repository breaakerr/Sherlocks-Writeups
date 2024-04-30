# HTB SHERLOCKS WRITEUP, DFIR BRUTUS, by breaakerr :)

## Segundo writeup de la serie Sherlocks de Hack the Box, ahora con BRUTUS, como se daran cuenta solamente por el nombre esta box trata sobre un ataque de fuerza bruta a un servidor de Confluence mediante el servicio SSH, despues de ganar acceso al servidor el atacante realizo algunas otras pillerias que vamos a ir viendo a medida que vayamos explorando el archivo más importante en cuestion, el auth.log, que ya de por si se suele utilizar para analizar ataques de fuerza bruta pero ahora vamos a ir un poco más allá explorando este archivo viendo cosas como habilidades de persistencia, escalado de privilegios y ejecución de comandos. Una MUY FACIL pero interesante box. 

### Primero lo primero y antes de pasar a las preguntas vamos a descomprimir los archivos que vienen incluidos en el .zip para ir viendo que tenemos de entrada. Y son dos ficheritos, auth.log y wtmp.
![alt text](Screenshot_1.png)

### Con esto ya listo para empezar a trabajar tenemos entonces varios topicos a analizar y enfocar de entrada, esos topicos son 
    *Persistencia
    *RCE
    *Escalado de privilegios
    *Fuerza bruta
    *Servidor Confluence
    
### Sabiendo esto vamos a avanzar entonces con las preguntas. 
## Task 1. Analizando el fichero auth.log, puedes identificar la dirección IP que uso el atacante para efectuar el ataque de fuerza bruta? Desde ya que si y ahora veremos como. Primero y principal vamos a abrir el fichero auth.log para ver que es lo que hay dentro y como está organizada la información (obviamente van a ser todas lineas, si fuese en formato json sería mas atractivo visualmente hablando)
![alt text](Screenshot_2.png)   

La tarea 1 nos pide averiguar la IP pero ahora que lo pienso vamos a matar varios pajaros de un tiro con la ayuda del buen bash. 