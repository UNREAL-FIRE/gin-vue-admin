import json

def main():
    json_file = 'CET4_T.json'
    sql_file = 'output.sql'

    with open(json_file, 'r', encoding='utf-8') as f:
        data = json.load(f)
    
    with open(sql_file, 'w', encoding='utf-8') as f:
        for entry in data:
            name = entry['name'].replace("'", "''")
            trans = ','.join(entry['trans']).replace("'", "''")
            usphone = entry['usphone'].replace("'", "''")
            ukphone = entry['ukphone'].replace("'", "''")
            sql = f"INSERT INTO cet4_t (name, trans, usphone, ukphone) VALUES ('{name}', '{trans}', '{usphone}', '{ukphone}');\n"
            f.write(sql)

if __name__ == '__main__':
    main()
