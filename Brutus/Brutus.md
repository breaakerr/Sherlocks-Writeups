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

La tarea 1 nos pide averiguar la IP pero ahora que lo pienso vamos a matar varios pajaros de un tiro con la ayuda del buen bash. 

    cat auth.log | grep sshd

![alt text](Screenshot_2.png)  

Y acá vemos MUCHAS cosas interesantes. Veamos en detalle:

    *AuthorizedKeysCommand: Es una directiva de configuración de SSH que especifica un comando que debe ejecutarse para obtener las claves públicas autorizadas para la autenticación del usuario.

    */usr/share/ec2-instance-connect/eic_run_authorized_keys: Es el comando que se está ejecutando como parte de la autorización de las claves públicas. En sistemas basados en AWS, EC2 Instance Connect proporciona un método seguro para conectarse a instancias de EC2 mediante la autenticación de claves SSH utilizando IAM (Identity and Access Management).

    *root: Indica que el usuario para el que se está intentando realizar la autenticación de claves públicas es el usuario root, que es el usuario administrativo en sistemas Linux.

    *SHA256:4vycLsDMzI+hyb9OP3wd18zIpyTqJmRq/QIZaLNrg8A: Esto parece ser una huella digital de la clave pública que se está utilizando para intentar la autenticación. El formato SHA256: sugiere que es el hash SHA-256 de la clave pública.

    *failed, status 22: Indica que la operación de obtener las claves públicas autorizadas ha fallado, y el código de estado 22 sugiere que ha habido un error relacionado con la validez de la clave o con el proceso de autenticación en sí.
