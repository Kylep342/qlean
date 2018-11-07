WITH v_foo AS (SELECT foo.first_name, foo.last_name, (foo.first_name || foo.last_name) AS full_name, foo.salary, bar.city FROM foo JOIN bar USING (id)) SELECT * FROM v_foo WHERE v_foo.salary < 40000;
