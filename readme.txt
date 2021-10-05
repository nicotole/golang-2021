El programa recibe una cadena y la fracciona, tomando los dos primeros caracteres como "tipo"
los dos siguientes como el largo definido para la secuencia posterior de caracteres de la cadena.

Si la cadena es de formato valido se retorna un resultado ej:{TX 3 ABC}.
Con la forma de la siguiente estructura:
    type Result struct {
        Type    string 
        Value   string
        Length  int
    }

Tomando de ejemplo la cadena: TX040001

TX son tomados como Type (Tipo de la cadena)
04 como Length (Largo de caracteres del Value)
0001 como Value (Valor)

Para obtener la estructura resultante de la cadena se debe utilizar el metodo publico:
GetStructure(<String Cadena>) del paquete Model. 

Ej: Model.GetStructure("TX040001")
    Este metodo retorna, para esta entrada, la estructura {TX 4 0001}

El metodo GetStructure(<String Cadena>) contempla cadenas invalidas y retorna nil, "Cadena invalida" 
y la cadena en caso de caso de encontrar una.

Una cadena es considerada invalida en caso de:
    + Demasiado corta como para ser fraccionada (minimo 5 caracteres)
    + Si el int no coincide con el largo del Bullet o Value
    + Si en la posicion del Length(En la cadena) en lugar de haber un int se encuentra un char
        Esta ultima se salvaguarda en caso de que el char sea convertible a int ej "01" = 1


+----By-------------------------------+ 
+    odelotsalocin@gmail.com          +
+    m.angeles.rasmussen@gmail.com    +
+-------------------------------------+
        


