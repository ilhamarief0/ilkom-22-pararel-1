require 'sinatra'
require 'json'

# Data sementara untuk menyimpan posts
posts = [
  { id: 0, title: 'title1', body: 'body1' },
  { id: 1, title: 'title2', body: 'body2' }
]

# Mendapatkan semua posts
get '/posts' do
  content_type :json
  posts.to_json
end

# Mendapatkan post berdasarkan ID
get '/posts/:id' do
  content_type :json
  id = params[:id].to_i
  post = posts.find { |p| p[:id] == id }
  halt(404, { message: 'Post not found' }.to_json) unless post
  post.to_json
end

# Membuat post baru
post '/posts' do
  content_type :json
  request_body = JSON.parse(request.body.read)
  new_id = posts.empty? ? 0 : posts.last[:id] + 1
  new_post = { id: new_id, title: request_body['title'], body: request_body['body'] }
  posts << new_post
  status 201
  new_post.to_json
end

# Mengupdate post berdasarkan ID
put '/posts/:id' do
  content_type :json
  id = params[:id].to_i
  post = posts.find { |p| p[:id] == id }
  halt(404, { message: 'Post not found' }.to_json) unless post

  request_body = JSON.parse(request.body.read)
  post[:title] = request_body['title'] if request_body['title']
  post[:body] = request_body['body'] if request_body['body']

  post.to_json
end

