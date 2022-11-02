/*Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.*/
SELECT CONCAT(e.nombre , " " , e.apellido) Nombre , e.puesto , d.localidad FROM empleado e
INNER JOIN departamento d ON d.depto_nro = e.depto_nro;
/*Visualizar los departamentos con más de cinco empleados.*/
SELECT d.nombre_depto , COUNT(e.nombre) FROM departamento d 
INNER JOIN empleado e ON e.depto_nro = d.depto_nro
GROUP BY d.nombre_depto;
/*Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.*/
SELECT CONCAT(e.nombre , " " , e.apellido) Nombre , e.salario , d.nombre_depto FROM empleado e
INNER JOIN departamento d ON d.depto_nro = e.depto_nro
WHERE e.pusto = (SELECT e.puesto FROM empleado WHERE code_emp = "E-0006");
/*Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.*/
SELECT CONCAT(e.nombre , " " , e.apellido) Nombre , e.puesto , e.salario FROM empleado e
INNER JOIN departamento d ON d.depto_nro = e.depto_nro
WHERE d.nombre_depto = "Contabilidad"
ORDER BY e.Nombre ;
/*Mostrar el nombre del empleado que tiene el salario más bajo.*/
SELECT CONCAT(e.nombre , " " , e.apellido) Nombre FROM empleado e 
ORDER BY e.salario ASC 
LIMIT 1;
/*Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.*/
SELECT CONCAT(e.nombre , " " , e.apellido) Nombre FROM empleado e
INNER JOIN departamento d ON d.depto_nro = e.depto_nro
WHERE d.nombre_depto = "Ventas"
ORDER BY e.salario 
LIMIT 1;