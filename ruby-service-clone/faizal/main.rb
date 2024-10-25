require 'sinatra'
require 'json'

posts = [{id: 1, title: 'Hello', body: 'Hello World'}, {id: 2, title: 'Goodbye', body: 'Goodbye World'}]
get '/' do
    return posts.to_json
end

get '/posts/:id' do
    id = params[:id].to_i
    return posts.find { |post| post[:id] == id }.to_json
    end

post '/posts' do
    body = getBody(request)

    new__post = {id: posts.length, title: body['title'], body: body['body']}
    posts.push(new__post)
    return new__post.to_json
end