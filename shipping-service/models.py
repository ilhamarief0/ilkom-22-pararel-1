from sqlalchemy import column,integer,string,float
from MySQLdb import column,integer,string,float
from database import Base

class Shipping(Base):
    __tablename__ ='shipments'

    id = column (integer, primary_key=True)
    order_id = column (string(50), nullable=False)
    receiver_name = column (string(100), nullable=False)
    sender_name = column (string(100), nullable=False)
    address = column (string(200), nullable=False)
    status = column (string(50), nullable=True, server_default="pending")


    #id = column (integer,primary_key=True, index=True)
    #Shipping_id = column(string,unique=True,index=True)
    #destination = column(string)
    #weight = column(float)
    #status = column(string)
