SELECT sub.name FROM (SELECT foo.first_name ||   || foo.last_name FROM foo JOIN bar ON foo.id = bar.id WHERE lower(bar.city) = chicago) sub;
