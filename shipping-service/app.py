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
    status = db.Column(db.String(50), nullable=True, server_default="pending")

    def __repr__(self):
        return f'<Shipment {self.id}>'


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
    
@app.route('/add_shipments', methods=['POST'])
def add_shipments():
    data = request.get_json()
    order_id = data.get('order_id')
    receiver_name = data.get('receiver_name')
    sender_name = data.get('sender_name')
    address = data.get('address')
    status = data.get('status', 'pending') 

    cursor = mysql.connection.cursor()
    sql = """
    INSERT INTO shipments (order_id, receiver_name, sender_name, address, status)
    VALUES (%s, %s, %s, %s, %s)
    """
    val = (order_id, receiver_name, sender_name, address, status)
    cursor.execute(sql, val)
    mysql.connection.commit()
    cursor.close()

    return jsonify({"message": "Shipping record added successfully!"}), 201

@app.route('/edit_shipments/<int:id>', methods=['PUT'])
def edit_shipments(id):
    data = request.get_json()
    order_id = data.get('order_id')
    receiver_name = data.get('receiver_name')
    sender_name = data.get('sender_name')
    address = data.get('address')
    status = data.get('status')

    cursor = mysql.connection.cursor()
    sql = """
    UPDATE shipments 
    SET order_id = %s, receiver_name = %s, sender_name = %s, address = %s, status = %s
    WHERE id = %s
    """
    val = (order_id, receiver_name, sender_name, address, status, id)
    cursor.execute(sql, val)
    mysql.connection.commit()
    cursor.close()

    return jsonify({"message": "Shipping record updated successfully!"}), 200


@app.route('/delete_shipping/<int:id>', methods=['DELETE'])
def delete_shipping(id):
    cursor = mysql.connection.cursor()
    sql = "DELETE FROM shipments WHERE id = %s"
    val = (id,)
    cursor.execute(sql, val)
    mysql.connection.commit()
    cursor.close()

    return jsonify({"message": "Shipping record deleted successfully!"}), 200


# Create the database tables in the application context
def create_tables():
    with app.app_context():
        db.create_all()

if __name__ == '__main__':
    # Create tables before starting the app
    create_tables()
    app.run(host='0.0.0.0', port=60, debug=True)
