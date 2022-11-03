/*PARTE 1
Con la base de datos “movies”, se propone crear una tabla temporal llamada “TWD” y 
guardar en la misma los episodios de todas las temporadas de “The Walking Dead”.*/
CREATE temporary TABLE `The_Walking_Dead_Episodes`
SELECT e.title Episodio, son.title Temporada , s.title Serie FROM episodes e 
INNER JOIN seasons son ON son.id = e.season_id
INNER JOIN series s  ON s.id = son.serie_id
WHERE s.title = "The Walking Dead";
/*Realizar una consulta a la tabla temporal para ver los episodios de la primera temporada.*/
SELECT * FROM The_Walking_Dead_Episodes twd
WHERE twd.Temporada = "Primer Temporada" ;
/*En la base de datos “movies”, seleccionar una tabla donde crear un índice y luego chequear la creación del mismo.*/
CREATE INDEX season_title ON seasons (title) ;
SHOW INDEX FROM seasons ;
/* Analizar por qué crearía un índice en la tabla indicada y con qué criterio se elige/n el/los campos.*/
/*PARTE 2
Agregar una película a la tabla movies.*/
INSERT INTO movies 
(title,rating,awards,release_date,length,genre_id) VALUES 
("Hillhouse",8.5,4,"2010-05-10",310,2);
/*Agregar un género a la tabla genres.*/
INSERT INTO genres
(name,ranking, active) VALUES
("Anime",13,1) ;
/*Asociar a la película del punto 1. con el género creado en el punto 2.*/
UPDATE movies m SET m.genre_id = 13 
WHERE m.title = "Hillhouse";
/*Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el punto 1.*/
UPDATE actors a SET a.favorite_movie_id = 22
WHERE a.id = 2 ;
/*Crear una tabla temporal copia de la tabla movies.*/
CREATE TEMPORARY TABLE `movies_copy`
SELECT * FROM movies ;
/*Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.*/
DELETE FROM movies_copy mc WHERE mc.awards < 5 ; 
/*Obtener la lista de todos los géneros que tengan al menos una película.*/
SELECT g.name , m.title FROM genres g
INNER JOIN movies m ON m.genre_id = g.id ; 
/*Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.*/
SELECT CONCAT(a.first_name , " " , a.last_name) Actores FROM actors a 
INNER JOIN movies m ON m.id = a.favorite_movie_id
WHERE m.awards > 3 ;
/*Crear un índice sobre el nombre en la tabla movies.*/
CREATE INDEX movie_title ON movies (title) ;
/*Chequee que el índice fue creado correctamente.*/
SHOW INDEX FROM movies ;
/*En la base de datos movies ¿Existiría una mejora notable al crear índices? Analizar y justificar la respuesta.*/
SELECT * FROM actors ;
EXPLAIN SELECT * FROM actors a WHERE a.first_name = "Zoe" ;
SELECT * FROM movies ;
EXPLAIN SELECT * FROM movies m WHERE m.title = "Titanic" ;
/*
La query que utiliza indices analiza muchas menos filas al momento de realizar la consulta, traduciendose en una
mejor performance al momento de obtener datos.
*/
/*¿En qué otra tabla crearía un índice y por qué? Justificar la respuesta
Creo que los nombres en esta base de datos son buenos candidatos para recibir un indice, como 
el nombre de las peliculas y de las series ya que es un elemento que se consulta mucho al momento
de realizar filtros.
*/