SELECT * FROM foo f LEFT OUTER JOIN bar b ON f.tertiary_id = b.tertiary_id WHERE foo.department IN ('Sales', 'Corporate') AND bar.region = 'Northeast' LIMIT 4 OFFSET 20;
