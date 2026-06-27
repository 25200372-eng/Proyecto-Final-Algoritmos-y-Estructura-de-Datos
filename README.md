# Proyecto-Final-Algoritmos-y-Estructura-de-Datos
Grupo 9, R-Tree, Consultas por rango (ventana) sobre un mapa de ciudades o puntos geográficos.
R-Tree: Consultas por Rango sobre Mapas Geográficos

Proyecto Final — Algoritmos y Estructura de Datos | Grupo 9

Implementación de un R-Tree para realizar consultas por rango (ventana) sobre un conjunto de ciudades o puntos geográficos, permitiendo búsquedas espaciales eficientes sobre grandes volúmenes de datos.


Descripción

Un R-Tree es una estructura de datos en árbol diseñada para indexar información espacial multidimensional. Este proyecto lo aplica al contexto geográfico: dado un mapa de ciudades o coordenadas, el sistema permite realizar consultas por ventana (range queries), es decir, encontrar todos los puntos que se encuentren dentro de un rectángulo de búsqueda definido por el usuario.

Este tipo de estructura es ampliamente utilizado en sistemas de información geográfica (GIS), bases de datos espaciales y motores de búsqueda por proximidad.


Características


Inserción dinámica de puntos geográficos (ciudades con coordenadas x, y)
Consultas por rango (ventana rectangular)
División de nodos con la heurística de Guttman (1984)
Visualización o reporte de los puntos encontrados en el rango consultado
Soporte para conjuntos de datos de ciudades reales o generados



Conceptos Clave

¿Qué es un R-Tree?

El R-Tree organiza los datos en Minimum Bounding Rectangles (MBR): cada nodo del árbol contiene un rectángulo que agrupa a sus hijos. Al buscar puntos en una ventana, se recorre el árbol descartando ramas cuyo MBR no intersecte con la ventana de búsqueda, lo que reduce drásticamente el número de comparaciones.

Consulta por Rango (Window Query)

Dado un rectángulo de búsqueda [x1, x2] × [y1, y2], la consulta retorna todos los puntos (px, py) tales que:

x1 ≤ px ≤ x2  y  y1 ≤ py ≤ y2


Estructura del Repositorio

Proyecto-Final-Algoritmos-y-Estructura-de-Datos/
│
├── ProyectoFinal.zip                  # Código fuente del proyecto
├── PPTX Grupo 9.pptx                  # Presentación del proyecto
├── Guttman (1984), "R-Trees...".pdf   # Paper original de R-Trees (Guttman)
├── The Priority R-Tree,...pdf         # Paper complementario (Priority R-Tree)
└── README.md                          # Este archivo


Instalación y Ejecución:

Clona el repositorio:
bash   git clone https://github.com/25200372-eng/Proyecto-Final-Algoritmos-y-Estructura-de-Datos.git
   cd Proyecto-Final-Algoritmos-y-Estructura-de-Datos


Descomprime el código fuente:
bash   unzip ProyectoFinal.zip


Compila y ejecuta según el lenguaje del proyecto (C++/Java/Python — ver carpeta descomprimida para instrucciones específicas).


Referencias: Guttman, A. (1984). R-Trees: A Dynamic Index Structure for Spatial Searching. SIGMOD '84.
Arge, L. et al. The Priority R-Tree: A Practically Efficient and Worst-Case Optimal R-Tree. (incluido en el repositorio)



Equipo — Grupo 9: Proyecto desarrollado como entrega final del curso de Algoritmos y Estructura de Datos.

Licencia:
Este proyecto es de uso académico. No se permite su redistribución sin autorización del equipo.
