                                     # LOCKPICK HTB WRITEUP by breaakerr :)
## Forela necesita tu ayuda! Una gran porcion de nuestros servidores UNIX fueron golpeados por lo que creemos es un ransomware :o y necesitamos que nos salves jeje gracias marta.  
## ADVERTENCIA Y ESTO ES ENSERIO. HAY SOFTWARE DAÑINO QUE VIENE INCLUIDO EN EL ZIP QUE NOS PROPORCIONA HACKTHEBOX, POR LO QUE SE RECOMIENDA MANEJAR CON EXTREMA CAUTELA ESTOS ARCHIVOS APARTE DE HACERLO EN MAQUINAS VIRTUALES ISOLADAS QUE NO ESTEN RELACIONADAS CON NUESTRAS PRACTICAS HABITUALES DE HACKING O LO QUE SEA, LEAN EL DANGER.TXT FILE QUE APARTE AHÍ TAMBIÉN SE LES ADVIERTE. 

Bueno como ya fue comentado, esta box es sobre Analisis de Malware, así que vamos a ver que trae entre manos este software que encripto archivos tan importantes de nuestro amiguito Forela.
![Screenshot from 2024-04-25 14-16-10](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/51f20235-5731-4e18-b84d-a5d5a179e33c)


Despues de descomprimir el .zip "lockpick1" nos encontramos con 2 archivos de texto y otro .zip que al intentar descomprimir me pide una contraseña.

![Screenshot_1](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/2d47218a-4141-44e0-b2e7-06427b6b8f4e)

La contraseña la encontré dentro del .txt DANGER

Después de descomprimir bescrypt.zip nos encontramos con un archivo .exe que OBVIAMENTE no vamos a ejecutar jeje prefiero seguir teniendo mi maquina virtual a salvo, en este caso es un programa malicioso tipo ransomware.

![Screenshot_2](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/ce2546c9-3345-487e-a9d5-b2425441915f)


Todo el archivaje encriptado se encuentra en el directorio "forela-criticaldata"

![Screenshot_3](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/ea803e76-4577-48ab-9f98-6f6d2edd0cb5)


Desde ya que si tratamos de abrir o leer el contenido de alguno de estos archivos somos notificados del encriptado de los mismos. 

![Screenshot_4](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/57af0c7f-f721-48e9-9e39-2dd398282650)


Lo que va a ocurrir ahora es que vamos a tratar de desencriptar los archivos :D pero primero necesitamos saber que mecanismo se utilizo o su clave de cifrado.
Ahora si, que se armen los pinches chingadazos >:)
## Task 1 - Proporcionar el string de la clave de cifrado utilizada para cifrar los archivos. 
Como ya sabrán, un string es una cadena de caracteres individuales, lo digo por si alguien se confunde con esa palabra. Para esto vamos a analizar el exe bscrypt3.2 haciendo un poco de ingenieria inversa basica, yo voy a utilizar radare (porque no conozco otra forma o programa ajskaj)

![Screenshot_5](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/7f7755f6-f85c-40ee-a040-c0fd29e8da5c)

Si bien hay varios ficheritos vamos a enfocarnos en el que aparenta ser el mas importante, el llamado "main" y vamos a revisar que se trae entre datos (badumtsss)

![Screenshot_6](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/ba34262c-4694-4be3-ad7c-3d4416b22029)

Acá es donde me pongo serio y les invito a leer o simplemente copiar y pegar la clave en la tarea asignada y seguir como si nada (no lo hagan, lean >:C)
Líneas 1-3: Prólogo Estándar de la Función
![Screenshot_7](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/70c6836a-7b05-4016-8ed4-75ebe1e87b0f)

-55 push rbp: Esta línea guarda el registro del puntero base (rbp) actual en la pila. El puntero base se utiliza para acceder a las variables locales dentro de la función.
-4889e5 mov rbp, rsp: Esta línea establece el puntero base (rbp) en el puntero de la pila actual (rsp). Esto establece el inicio del nuevo marco de pila para esta función.
-4883ec10 sub rsp, 0x10: Esta línea asigna algo de espacio en la pila para las variables locales. Resta 0x10 (16 bytes) del puntero de la pila, creando efectivamente un nuevo "marco" en la pila para esta función.

Líneas 4-5: Cargando Direcciones de Cadenas

![Screenshot_8](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/3fd006d8-f269-46b0-bf47-16cf832bb717)

-488d05280a00. lea rax, str.bhUlIshutrea98liOp ; 0x217d ; "bhUlIshutrea98liOp": Esta línea usa la instrucción lea (obtener dirección efectiva). Calcula la dirección de la cadena "bhUlIshutrea98liOp" (que probablemente esté almacenada en otra parte de la memoria) y la almacena en el registro rax.
-488945f8 mov qword [var_8h], rax: Esta línea mueve la dirección almacenada en rax (que apunta a la cadena "bhUlIshutrea98liOp") a la ubicación de memoria en [rbp-0x8]. Esto efectivamente crea una variable local llamada var_8h y almacena la dirección de la cadena en ella.

Líneas 6-7: Cargando Direcciones de Cadenas (Otra Cadena)

![Screenshot_9](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/d2b768e3-952a-4936-bac5-ba0e5970141d)

-Esta parte sigue la misma lógica que las líneas 4-5. Calcula la dirección de la cadena "/forela-criticaldata/" y la almacena en la variable local var_10h en [rbp-0x10].

Líneas 8-11: Configurando Argumentos de Llamada de Función

![Screenshot_10](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/d333f077-c197-4ab1-9d68-01e443a155a6)

-488b55f8 mov rdx, qword [var_8h]: Esta línea carga el valor almacenado en var_8h (que es la dirección de "bhUlIshutrea98liOp") en el registro rdx.
-488b45f0 mov rax, qword [var_10h]: Esta línea carga el valor almacenado en var_10h (que es la dirección de "/forela-criticaldata/") en el registro rax.
-4889d6 mov rsi, rdx ; int64_t arg2: Esta línea mueve el contenido de rdx (dirección de "bhUlIshutrea98liOp") al registro rsi. Esto probablemente establece el segundo argumento para la llamada a la función.
-4889c7 mov rdi, rax ; char *arg1: Esta línea mueve el contenido de rax (dirección de "/forela-criticaldata/") al registro rdi. Esto probablemente establece el primer argumento para la llamada a la función.

Línea 12: Llamando a la Función

![Screenshot_11](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/9e67f920-c8e2-47c7-ba68-dc708c798df7)

-e85dfdffff call sym.process_directory: Esta línea es la llamada a la función real. Utiliza la instrucción call y especifica la dirección de la función a la que se llamará, que en este caso es sym.process_directory. Es probable que aquí sea donde ocurre el procesamiento principal con los argumentos proporcionados.

Líneas 13-15: Epílogo y Retorno

![Screenshot_12](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/3b31d488-3f8d-4ad5-aace-35feca698333)

-b800000000 mov eax, 0: Esta línea establece el registro eax en 0, lo que podría usarse como un valor de retorno para la función (aunque no está garantizado).
-c9 leave: Esta línea restaura el puntero base anterior (rbp) de la pila, dejando efectivamente el marco de pila actual.
-c3 ret: Esta línea ejecuta la instrucción ret, que devuelve el control al código que llamó a esta función.

O sea digamos, este fragmento de código define una función que toma dos argumentos de cadena, prepara los argumentos en la pila, llama a otra función (sym.process_directory) para hacer el trabajo real y luego regresa (potencialmente con un valor de 0).
TASK1: bhUlIshutrea98liOp

## Task 2: SOSPECHOSAMENTE un tal wbevansn1@cocolog-nifty.com nos pide el nombre y apellido con el que lo registramos. Creen que cometieron un error en el proceso de solicitud. Tenemos que confirmar nombre y apellido del solicitante. 
Acá hay algo que inevitablemente hay que hacer y es averiguar cual diablos es el mecanismo de encriptado partiendo de que tenemos la clave de cifrado utilizada. 
Como vieron en la tarea 1, hay un proceso llamado "sym.process_directory" que sirve para obtener los archivos en los directorios. Voy a dejar adjunto en el repositorio un archivo llamado "Mainfunc_1stcall" para que puedan hecharle un ojo. 

Al final de la función revisada se ve que llama a otra funcion: "call sym.encrypt_file". El procedimiento es el siguiente: agarra la información contenida en el fichero, la cifra byte a byte, luego borra el archivo original y lo reemplaza por una copia con el mismo nombre pero con la información encriptada :c

La verdad es que para el momento en el que estoy escribiendo esto ya tenia todo hecho, pero quedaba como un vacio en los pasos siguientes, asi que decidí que explicar esta nueva función seria lo mas óptimo jeje.

Para no atosigarlos con strings en assembler y saltos de espacios de memoria vamos a pasar el contenido del proceso de encriptamiento a un formato mas legible. Mediante el uso de varias herramientas de ingenieria inversa, como ghidra, autopsy y vscode terminé con un codigo que es algo así: 

#include <stdio.h> // para fopen, fclose, printf, etc.

#include <string.h> // para strlen, snprintf, etc.

#include <stdlib.h> // para malloc, free

bool encriptar_archivo(const char* archivo_entrada, const char* archivo_salida) 
{
  // Abrir el archivo de entrada
  FILE* archivo_entrada_ptr = fopen(archivo_entrada, "rb");
  if (archivo_entrada_ptr == nullptr) {
    printf("Error al abrir el archivo: %s\n", archivo_entrada);
    return false;
  }

  // Obtener el tamaño del archivo
  fseek(archivo_entrada_ptr, 0, SEEK_END);
  long int tam_archivo = ftell(archivo_entrada_ptr);
  rewind(archivo_entrada_ptr);

  // Reservar memoria para los datos del archivo
  void* datos_archivo = malloc(tam_archivo);
  if (datos_archivo == nullptr) {
    printf("Error al reservar memoria\n");
    fclose(archivo_entrada_ptr);
    return false;
  }

  // Leer los datos del archivo
  size_t bytes_leidos = fread(datos_archivo, 1, tam_archivo, archivo_entrada_ptr);
  if (bytes_leidos != tam_archivo) {
    printf("Error al leer el archivo\n");
    free(datos_archivo);
    fclose(archivo_entrada_ptr);
    return false;
  }

  // Cerrar el archivo de entrada
  fclose(archivo_entrada_ptr);

  // Realizar la encriptación (remplazar con la lógica de encriptación real)
  // ...

  // Abrir el archivo de salida
  char nombre_archivo_salida[256];
  snprintf(nombre_archivo_salida, sizeof(nombre_archivo_salida), "%s.cifrado", archivo_entrada);
  FILE* archivo_salida_ptr = fopen(nombre_archivo_salida, "wb");
  if (archivo_salida_ptr == nullptr) {
    printf("Error al crear el archivo de salida: %s\n", nombre_archivo_salida);
    free(datos_archivo);
    return false;
  }

  // Escribir los datos encriptados en el archivo de salida
  size_t bytes_escritos = fwrite(datos_archivo, 1, tam_archivo, archivo_salida_ptr);
  if (bytes_escritos != tam_archivo) {
    printf("Error al escribir el archivo\n");
    free(datos_archivo);
    fclose(archivo_salida_ptr);
    return false;
  }

  // Cerrar el archivo de salida
  fclose(archivo_salida_ptr);

  // Escribir un archivo de nota (opcional)
  // ...

  // Liberar la memoria reservada
  free(datos_archivo);

  // Eliminar el archivo original (opcional)
  if (remove(archivo_entrada) != 0) {
    printf("Error al eliminar el archivo original: %s\n", archivo_entrada);
  }

  return true;
}

En realidad el código está en inglés pero como no vamos a utilizarlo lo traduje a español así es mucho mas legible <3 Ahora si, como dije antes, tenemos la clave de cifrado, por lo que mediante un pequeño script en Go vamos a descifrar los ficheros yaaaaaaaaaaaaaaaaaaaaaaaaaaay!! :D

![Screenshot_13](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/c41ce56f-a336-4049-b532-fb236ad7ea54)

![Screenshot_14](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/511f7de0-00a3-4031-b4fd-9ebb87ea205a)


Despues de terminar mi descifrainador y ubicarlo en el directorio lockpick1 lo ejecuto y logramos descifrar todos los archivos cifrados :D

![Screenshot_15](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/6ab128f2-8891-4d6f-beea-7ce2d37556d4)


### Acá les va una pequeña explicación del código, pueden saltearse esto si les apetece jeje 
listFiles: Esta función recibe un directorio como argumento y devuelve una lista de nombres de archivos con la extensión .24bes dentro de ese directorio.

Utiliza ioutil.ReadDir para obtener información sobre los archivos en el directorio.

Itera sobre cada archivo en el directorio.

Comprueba si el archivo es un directorio (file.IsDir()). Si lo es, lo salta.

Comprueba si la extensión del archivo es .24bes utilizando filepath.Ext.

Si el archivo cumple con los criterios anteriores, agrega su nombre a la lista de archivos.

decrypt: Esta función recibe la ruta del directorio y el nombre del archivo a descifrar. Luego, descifra el archivo utilizando un algoritmo de cifrado XOR y escribe el resultado en un nuevo archivo en el directorio decrypt.

Genera la ruta al directorio de destino para los archivos descifrados.

Lee el contenido del archivo a descifrar usando ioutil.ReadFile.

Itera sobre cada byte del contenido del archivo y lo descifra usando el algoritmo de cifrado XOR.

Escribe el contenido descifrado en un nuevo archivo en el directorio decrypt usando ioutil.WriteFile.

Imprime un mensaje indicando que el archivo ha sido descifrado con éxito.

main: Esta es la función principal del programa. Es donde se ejecuta el flujo principal del programa.
Llama a listFiles para obtener la lista de archivos a descifrar.
Comprueba si hay archivos para descifrar. Si no hay ninguno, imprime un mensaje y termina la ejecución.
Itera sobre cada archivo y llama a decrypt para descifrarlo.
Muestra el progreso de la operación de descifrado.
Una vez que se completan todas las operaciones de descifrado, imprime un mensaje indicando que la operación ha finalizado.

Ahora si, dentro de la carpeta forela-criticaldata vamos a encontrar una carpeta llamada "descifrados" 

![Screenshot_16](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/8029e609-b671-45f7-b7ce-55de9beb088e)


#AHORA SI VAMOS A PODER RESPONDER TODAS LAS PREGUNTAS!!!! 

Retomando la pregunta original del correo de wbevansn1, vamos a buscar en la base de datos filtrando por su nombre a ver que encontramos. 

![Screenshot_17](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/4c31b451-ffd1-4e94-a86e-fd40e7168bd9)


## Task 3. Cual es la MAC y serial de la notebook de mi buen amigo Hart Manifould? 
Esto vamos a buscarlo en el fichero it_assets.xml

![Screenshot_18](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/1a8c0bbc-a7fb-49a5-a50d-c0ed84a909b4)

<MAC>E8-16-DF-E7-52-48</MAC><asset_type>laptop</asset_type><serial_number>1316262</serial_number>

## Task 4. Cual es la dirección de correo del atacante? 
Cada archivo cifrado tiene un mensaje de amenaza diciendo cual es el correo. En este caso best24@protonmail.com

![Screenshot_19](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/b19ce32e-f923-4737-ac96-a1ab707e7735)


## Task 5. Los CIBERPATRULLADORES de Londres sospechan de un insider que colaboro en el ataque, averiguar quien fue la persona que sacó el mayor porcentaje de ganancia con un solo movimiento.
Acá vamos a recurrir al archivo trading-firebase_bkup.json, al principio me puse a buscar uno por uno hasta que me dí cuenta de la cantidad INMENSA de información que había, asi que con ayuda de bash filtramos algo de información y obtenemos el resultado. 

![Screenshot_20](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/b59fe6d9-b059-4eaa-9645-d8a5e67e15c8)


## Task 6. Nuestro equipo quiere saber la dirección IP de cierta personita que se cree está compartiendo su cuenta, investigar a O’Hederscoll Karylin. Vamos a buscar en sales_forecast.xlsx y filtrar por el nombre de O’Hederscoll Karylin.
Para la siguiente tarea utilicé un programa llamado "visidata" para instalarlo (si tienen kali como yo) haganse primero un "sudo apt update" y despues "apt install visidata". El comando para utilizarlo es "vd *archivo*" Así pudimos encontrar a la famosa Karylin.

![Screenshot_21](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/3b2cf0c3-c3d1-4b26-8706-6215acea70e0)


## Task 7. Cual de las siguientes extensiones no fueron atacadas por el malware? .txt, .sql,.ppt, .pdf, .docx, .xlsx, .csv, .json, .xml.

![Screenshot_22](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/dc7f82bc-00b6-46d3-9fdc-e39c138528d6)

Tal como se ve en la imagen, la extensión que no fue encriptada fue .ppt

## Task 8, 9 y 10. Estas van de lo mismo, se quiere verificar0 la integridad de los archivos mediante su hash. Debemos dar el hash correspondiente a cada archivo después del descifrado.
![Screenshot_23](https://github.com/breaakerr/Sherlocks-Writeups/assets/141375378/424cd924-ef49-41b1-836b-7d6b4549fb87)

