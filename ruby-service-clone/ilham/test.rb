require 'sinatra'
require 'json'

posts = [{id: 0, title: 'title1', body: 'body1'}, {id: 1, title: 'title2', body: 'body2'}]

get '/posts' do
  return posts.to_json
end

get '/posts/:id' do
  id = params[:id].to_i
  return posts[id].to_json
end