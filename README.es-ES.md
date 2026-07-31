

# **Crunch: Tu herramienta definitiva para la compresión de datos**

![demo](./demo.gif "demo")

## **Introducción**

Crunch es una robusta herramienta de línea de comandos escrita en Go que te permite lograr una compresión de datos eficiente utilizando el venerable algoritmo de codificación Huffman. Agiliza el proceso, haciendo que comprimir tus archivos sea un juego de niños, manteniendo al mismo tiempo claridad y control.

## **Características**

- **Compresión sin pérdidas:** Crunch preserva la integridad de tus datos, garantizando una reconstrucción exacta al descomprimirlos.
- **Verbosidad personalizable:** Controla el nivel de detalle mostrado durante la compresión con la bandera `-v`, permitiéndote adaptar la salida a tus preferencias.
- **Manejo simplificado de archivos:** Crunch maneja sin problemas tanto los archivos de entrada como los de salida, haciendo que los flujos de trabajo de compresión sean sencillos.

## **Cómo compilar**
   ```bash
   go build -o crunch .
   ```
   O 
   ```bash
   ./build.sh
   ```

## **Uso**

1. **Codificar datos:**
   ```bash
   crunch -i input.txt -o compressed.dat
   ```
   Este comando comprime `input.txt` y guarda los datos comprimidos en `compressed.dat`.

2. **Decodificar datos:**
   ```bash
   crunch -i compressed.dat -o output.txt -d 
   ```
   Este comando descomprime los datos que fueron comprimidos con el comando del primer ejemplo.

3. **Modo verboso**
   ```bash
   crunch -i sample_data.txt -o encoded.dat -v
   ```
   Este ejemplo comprime `sample_data.txt` mientras proporciona mensajes informativos sobre el progreso del proceso de compresión.

## **Contribuciones**

¡Agradecemos las contribuciones a Crunch! Siéntete libre de bifurcar el repositorio en GitHub (el enlace llegará pronto) y enviar solicitudes de extracción (pull requests) para corrección de errores, mejoras o actualizaciones de la documentación.

## **Licencia**

Crunch se distribuye bajo la Licencia MIT (consulte el archivo LICENSE).

## **Descargo de responsabilidad**

Aunque Crunch ofrece una compresión efectiva, es posible que algunos tipos de archivos no experimenten una reducción significativa de tamaño debido a redundancias inherentes en los propios datos.

## **Mejoras futuras**

- Soporte para múltiples algoritmos de compresión (p. ej., LZMA)
- Integración con herramientas populares de gestión de archivos
- Funcionalidad de descompresión (próximamente)

**¡Agradezco tu interés en Crunch!**
