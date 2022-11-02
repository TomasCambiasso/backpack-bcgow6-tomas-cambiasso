#1,2,3, 4  Agregar una película a la tabla movies, Agregar un género a la tabla genres, asociarlos.
# Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el punto 1.
insert into movies (title, rating, awards, release_date, length) values ("PeliculaMia", 89.1, 2, "2022-11-02 12:00:00", 140);
insert into genres (name, ranking, active) values ("Disco", 100, 1);
update movies set genre_id = (select id from genres where name like "Disco") where id = 22;
update actors set favorite_movie_id = 22 where id = 4;

# 5 Crear tabla temporal y copiar movies
CREATE Temporary TABLE `movies_copy1` 
select * from movies;

# 6 Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
delete from movies_copy1 where awards < 5;
select * from movies_copy1;

# 7 Obtener la lista de todos los géneros que tengan al menos una película.
select name from genres g inner join movies m on g.id = m.genre_id;

# 8 Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
select * from actors a inner join movies m on a.favorite_movie_id = m.id
where m.awards > 3;
  
# 9 Crear un índice sobre el nombre en la tabla movies.
create index test_idx on movies(rating);

# 11 No creo porque ya existen 2 indices bastante utiles que son el de la PK y el de FK. Los demas datos no son necesariamente unicos, tienen imagen pequeña y agregar indices sobre
# ellos no agregaria ningun valor, solo agregaria costo.

# 12 En ninguna. Todas ya tienen bastante cubierto los posibles indices. 