require 'sinatra'
require 'json'

items = [] 

post '/items' do
  data = JSON.parse(request.body.read) # Parsing JSON dari body permintaan
  item = { name: data['name'], quantity: data['quantity'] } # Buat item baru sebagai hash

  items << item # Tambahkan item ke array items
  status 201 # Kode status untuk "Created"
  item.to_json # Kembalikan item yang disimpan dalam format JSON
end

get '/items' do
  content_type :json
  items.to_json # Kembalikan semua item dalam format JSON
end
 
