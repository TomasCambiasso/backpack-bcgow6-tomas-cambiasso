# 1 Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.alter
select emp.name, emp.puesto, dep.localidad from empleado emp inner join departamento dep on emp.depto_nro = dep.depto_nro;

# 2 Visualizar los departamentos con más de cinco empleados.
select dep.nombre_depto, count(emp.depto_nro) as cant from empleado emp inner join departamento dep on emp.depto_nro = dep.depto_nro
group by dep.nombre_depto; ## Creo que esto funciona, no estoy seguro

# 3 Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.
select emp.nombre, emp.salario, dep.nombre_depto from empleado emp inner join departamento dep on emp.depto_nro = dep.depto_nro
	where emp.nombre in (select emp2.name from empleado emp2 where emp2.nombre like "Mito Barchuck");

INSERT INTO `biblioteca`.`prestamo` (`idLector`, `idLibro`, `fechaPrestamo`, `fechaDevolucion`, `devuelto`) VALUES (1, 2, NULL, NULL, 0);
INSERT INTO `biblioteca`.`prestamo` (`idLector`, `idLibro`, `fechaPrestamo`, `fechaDevolucion`, `devuelto`) VALUES (2, 2, NULL, NULL, 1);
INSERT INTO `biblioteca`.`prestamo` (`idLector`, `idLibro`, `fechaPrestamo`, `fechaDevolucion`, `devuelto`) VALUES (3, 1, NULL, '2021-07-16', 0);
INSERT INTO `biblioteca`.`prestamo` (`idLector`, `idLibro`, `fechaPrestamo`, `fechaDevolucion`, `devuelto`) VALUES (4, 5, NULL, NULL, 1);
INSERT INTO `biblioteca`.`prestamo` (`idLector`, `idLibro`, `fechaPrestamo`, `fechaDevolucion`, `devuelto`) VALUES (1, 3, NULL, '2021-07-16', 0);
#### Biblioteca

# 1 Listar los datos de los autores.
select * from autor;

# 2 Listar nombre y edad de los estudiantes
select nombre, edad from estudiante;

# 3 ¿Qué estudiantes pertenecen a la carrera informática?
select nombre from estudiante where carrera like "informatica";

# 4  ¿Qué autores son de nacionalidad francesa o italiana?
Select * from autor where nacionalidad like "Nacionalidad1" or nacionalidad like "Nacionalidad4";

# 5 ¿Qué libros no son del área de internet?
select * from libro where area not like "internet";

# 6 Listar los libros de la editorial Salamandra.
select * from libro where editorial like "Salamandra";

# 7 Listar los datos de los estudiantes cuya edad es mayor al promedio.
select avg(edad) from estudiante;
select * from estudiante where edad > (select avg(edad) from estudiante);

# 8 Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
Select nombre from estudiante where apellido like "G%";

# 9 Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).
select nombre from libro l inner join libroautor la on l.idLibro = la.idLibro 
inner join autor a on la.idAutor = a.idAutor where l.titulo like "El Universo: Guía de viaje";

# 10 ¿Qué libros se prestaron al lector “Filippo Galli”?
select * from estudiante e inner join prestamo p on e.idLector = p.idLector
inner join libro l on p.idLibro = l.idLibro where e.nombre like "Fillipo" and e.apellido like "Gali";

# 11 Listar el nombre del estudiante de menor edad.
Select nombre, min(edad) from estudiante;

# 12 Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.
select distinct nombre from prestamo p inner join estudiante e on p.idLector = e.idLector;

# 13 Listar los libros que pertenecen a la autora J.K. Rowling.
select * from libro l inner join libroautor la on l.idLibro = la.idLibro 
inner join autor a on la.idAutor = a.idAutor where a.nombre like "j.k rowling";

# 14 Listar títulos de los libros que debían devolverse el 16/07/2021.
select * from libro l inner join prestamo p on l.idLibro = p.idLibro where p.fechaDevolucion = "2021-07-16";