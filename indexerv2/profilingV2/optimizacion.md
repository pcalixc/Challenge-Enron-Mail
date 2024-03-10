
#  Optimización


## Perfil de CPU:
### Version 1
-   **Duración**: 717.62s
-   **Total de muestras**: 98.78s
-   **Porcentaje de muestras totales**: 13.76%
### Version 2
-   **Duración**: 287.65s
-   **Total de muestras**: 124.54s
-   **Porcentaje de muestras totales**: 43.30%

### Análisis: 
-   En la segunda versión, se ha reducido significativamente el tiempo total de ejecución.
-   Se ha observado una redistribución del tiempo entre las diferentes funciones.
-   Se refleja una la disminución del tiempo dedicado a `runtime.memclrNoHeapPointers` y `runtime.memmove`.
-   La función `indexerv2/utils.ConvertEmailFileToStruct` ahora representa  parte importante del tiempo de ejecución en la segunda versión.


## Perfil de Memoria:
### Version 1 
- **Uso de Memoria**: 14483.73kB
### Version 2
- **Uso de Memoria**: 34.01MB

### Análisis: 
-   La segunda versión muestra un aumento significativo en el uso de memoria.
-   En ambas versiones, las operaciones de manipulación y serialización de datos, así como las operaciones de manipulación de cadenas, son responsables de una parte significativa del uso de memoria.

## Mejoras

1.  **Implementación de Concurrencia:** Utilizas goroutines para procesar archivos de correo electrónico de manera concurrente, lo que mejora significativamente la eficiencia de la indexación.
    
2.  **Uso de Channels:** La comunicación entre las goroutines se realiza a través de canales, lo que facilita la coordinación y la transmisión de datos de manera segura entre las diferentes partes de la aplicación.
    
3.  **Utilización de Workers y Worker Pool:** He implementado una estructura de trabajadores y un pool de trabajadores para administrar el procesamiento concurrente de archivos. Esto permite una gestión eficiente de los recursos y una mejor escalabilidad.
   
6.  **Diseño Modulizado:** La aplicación está dividida en funciones y métodos bien definidos, lo que mejora la legibilidad, el mantenimiento y la escalabilidad del código.








