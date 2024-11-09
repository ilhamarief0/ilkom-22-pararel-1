require 'sinatra'
require 'json'

# CREATE: Tambah item baru
post '/items' do
    data = JSON.parse(request.body.read)
    item = Item.new(name: data['name'], quantity: data['quantity'])
    if item.save
      item.to_json
    else
      halt 422, item.errors.full_messages.to_json
    end
  end

# READ: Dapatkan semua item
get '/items' do
  Item.all.to_json
end

# READ: Dapatkan item berdasarkan ID
get '/items/:id' do
  item = Item.find_by(id: params[:id])
  if item
    item.to_json
  else
    halt 404, { error: 'Item not found' }.to_json
end
# UPDATE: Perbarui item berdasarkan ID
put '/items/:id' do
  item = Item.find_by(id: params[:id])
  if item
    data = JSON.parse(request.body.read)
    item.update2(name: data['name'], quantity: data['quantity'])
    item.to_json
  else
    halt 404, { error: 'Item not found' }.to_json
  end
end

