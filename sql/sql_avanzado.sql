# 1. Mostrar el título y el nombre del género de todas las series.
select s.title, g.name from series s inner join genres g
on s.genre_id = g.id;

# 2 Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
select e.title, a.first_name, a.last_name from episodes e inner join actor_episode ae on e.id = ae.episode_id
	inner join actors a on a.id = ae.actor_id;

# 3. Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
select count(ss.serie_id), s.title from seasons ss inner join series s on s.id = ss.serie_id
group by s.title;

# 4 Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
select g.name, count(m.genre_id) as total from movies m inner join genres g on m.genre_id = g.id
group by g.name having total > 3;

# 5 Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
select distinct a.first_name, a.last_name from movies m  inner join actor_movie am on m.id = am.movie_id
inner join actors a on am.actor_id = a.id where m.title like "La Guerra de las galaxias%";

