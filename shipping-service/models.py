from sqlalchemy import column,integer,string,float
from database import Base

class Shipping(Base):
    __tablename__ ='shippings'

    id = column (integer,primary_key=True, index=True)
    Shipping_id = column(string,unique=True,index=True)
    destination = column(string)
    weight = column(float)
    status = column(string)
