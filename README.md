# how to map entity with singular table name
    - Solution : https://stackoverflow.com/questions/44589060/how-to-set-singular-name-for-a-table-in-gorm

#NOTE

- Harus declare default value di sisi database

        nextval('sequence_name'::regclass)

- 