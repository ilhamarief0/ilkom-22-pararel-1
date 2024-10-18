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


