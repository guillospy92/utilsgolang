# Tamaño y formato (png) de la imagen
set terminal png size 600,400
# Nombre del fichero de salida
set output "rendimiento_web1.png"
# Título
set title "1000 peticiones, 50 peticiones concurrentes"
# Proporción
set size ratio 0.6
# Lineas Eje-Y
set grid y
# Leyenda X
set xlabel "Nº de peticiones"
# Leyenda Y
set ylabel "Tiempo de respuesta (ms)"
# Graficar los datos
plot "resultado.tsv" using 9 smooth sbezier with lines title "www.meneame.net"
exit
