require 'sinatra'
require 'json'

# In-memory data storage
items = [
  { id: 1, name: 'Item 1', description: 'This is item 1' },
  { id: 2, name: 'Item 2', description: 'This is item 2' }
]

# Routes
# Get all items
get '/items' do
  content_type :json
  items.to_json
end

# Get a single item by id
get '/items/:id' do
  content_type :json
  item = items.find { |i| i[:id] == params[:id].to_i }
  if item
    item.to_json
  else
    status 404
    { error: "Item not found" }.to_json
  end
end

# Create a new item
post '/items' do
  content_type :json
  new_item = {
    id: items.size + 1,
    name: params[:name],
    description: params[:description]
  }
  items << new_item
  status 201
  new_item.to_json
end

# Update an item
put '/items/:id' do
  content_type :json
  item = items.find { |i| i[:id] == params[:id].to_i }
  if item
    item[:name] = params[:name] if params[:name]
    item[:description] = params[:description] if params[:description]
    item.to_json
  else
    status 404
    { error: "Item not found" }.to_json
  end
end

# Delete an item
delete '/items/:id' do
  content_type :json
  item = items.find { |i| i[:id] == params[:id].to_i }
  if item
    items.delete(item)
    status 204
  else
    status 404
    { error: "Item not found" }.to_json
  end
end
