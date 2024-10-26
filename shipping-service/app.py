from flask import Flask, jsonify, request
from flask_mysqldb import MySQL
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)

# MySQL config
app.config['MYSQL_HOST'] = 'localhost'
app.config['MYSQL_USER'] = 'root'
app.config['MYSQL_PASSWORD'] = ''
app.config['MYSQL_DB'] = 'ecommerce'

# SQLAlchemy config
app.config['SQLALCHEMY_DATABASE_URI'] = 'mysql://root:@localhost/ecommerce'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False  # Optional: suppresses warning

# Initialize MySQL and SQLAlchemy
mysql = MySQL(app)
db = SQLAlchemy(app)

# Define the Shipment model
class Shipment(db.Model):
    __tablename__ = 'shipments'
    
    id = db.Column(db.Integer, primary_key=True)
    order_id = db.Column(db.String(50), nullable=False)
    receiver_name = db.Column(db.String(100), nullable=False)
    sender_name = db.Column(db.String(100), nullable=False)
    address = db.Column(db.String(200), nullable=False)

    def __repr__(self):
        return f'<Shipment {self.id}>'

@app.route('/')
def root():
    return 'RESTful API'

@app.route('/person')
def person():
    return 'Halo Person'

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
        sql = "SELECT * FROM shipments WHERE id = %s"
        val = (request.args['id'],)
        cursor.execute(sql, val)

        # Get column names from cursor.description
        column_names = [i[0] for i in cursor.description]

        # Fetch data and format into list of dictionaries
        rows = cursor.fetchall()
        data = [dict(zip(column_names, row)) for row in rows]

        cursor.close()  # Pastikan close sebelum return
        return jsonify(data)

# Create the database tables in the application context
def create_tables():
    with app.app_context():
        db.create_all()

if __name__ == '__main__':
    # Create tables before starting the app
    create_tables()
    app.run(host='0.0.0.0', port=60, debug=True)
