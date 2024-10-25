from flask import Flask, jsonify, request
from flask_mysqldb import MySQL

app = Flask(__name__)

# MySQL config
app.config['MYSQL_HOST'] = 'localhost'
app.config['MYSQL_USER'] = 'root'
app.config['MYSQL_PASSWORD'] = ''
app.config['MYSQL_DB'] = 'shipping_service'
mysql = MySQL(app)

@app.route('/')
def root():
    return 'restful API'

@app.route('/person')
def person():
    return 'halo person'

@app.route('/shipments')
def shipments():
    cursor = mysql.connection.cursor()
    cursor.execute("SELECT * FROM shipments")

    # Get column names from cursor.description
    column_names = [i[0] for i in cursor.description]

    # Fetch data and format into list of dictionaries
    rows = cursor.fetchall()
    data = [dict(zip(column_names, row)) for row in rows]

    cursor.close()  # Pindahkan close sebelum return
    return jsonify(data)

@app.route('/detailshipments')
def detailshipments():
    if 'id' in request.args:
        cursor = mysql.connection.cursor()
        sql = "SELECT * FROM shipments WHERE shipment_id = %s"
        val = (request.args['id'],)
        cursor.execute(sql, val)

        # Get column names from cursor.description
        column_names = [i[0] for i in cursor.description]

        # Fetch data and format into list of dictionaries
        rows = cursor.fetchall()
        data = [dict(zip(column_names, row)) for row in rows]

        cursor.close()  # Pastikan close sebelum return
        return jsonify(data)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=60, debug=True)

