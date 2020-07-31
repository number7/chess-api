This program needs a CSV file that has four fields in this order:
  - ECO Code
  - Opening Name
  - Opening Variation
  - Move sequence for the opening

The program needs the DB info in addition.

Usage of this program:

openings-loader \
  -file FILENAME \
  -host localhost \
  -port 5432 \
  -uid USERNAME \
  -pwd PASSWORD \
  -dbname chess

