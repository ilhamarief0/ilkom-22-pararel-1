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
