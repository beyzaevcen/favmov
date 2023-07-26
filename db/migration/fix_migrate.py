file_name = "./db/migration/000002_dump.up.sql"

id_sep = "SELECT pg_catalog.setval('public."
remove_line = "INSERT INTO public.schema_migrations"

data = []
with open(file_name, "r") as inp:
    data = inp.readlines()

for index, item in enumerate(data):
    if item.startswith(id_sep):
        end_before = item.split(id_sep)[1]
        table = end_before.split("_id_seq")[0]
        data[index] = id_sep + table + "_id_seq', (SELECT MAX(id) FROM public."+table+"));"
    elif item.startswith(remove_line):
        data[index] = ""
with open(file_name, "w") as out:
        out.writelines(data)


