require 'sinatra'
require 'json'

posts = [{id: 1, title: 'Hello', body: 'Hello World'}, {id: 2, title: 'Goodbye', body: 'Goodbye World'}]
get '/' do
    return posts.to_json
end