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
end

# UPDATE: Perbarui item berdasarkan ID
put '/items/:id' do
  item = Item.find_by(id: params[:id])
  if item
    data = JSON.parse(request.body.read)
    item.update(name: data['name'], quantity: data['quantity'])
    item.to_json
  else
    halt 404, { error: 'Item not found' }.to_json
  end
end

# DELETE: Hapus item berdasarkan ID
delete '/items/:id' do
  item = Item.find_by(id: params[:id])
  if item
    item.destroy
    { message: 'Item deleted' }.to_json
  else
    halt 404, { error: 'Item not found' }.to_json
  end
end


