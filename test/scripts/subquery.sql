SELECT sub.full_name FROM (SELECT (foo.first_name || foo.last_name) AS full_name FROM foo JOIN bar ON foo.id = bar.id WHERE lower(bar.city) = chicago) sub;
