require 'eventmachine'
require 'faye'
# require 'sinatra/base'
 
# class ConsumerApp < Sinatra::Base
#   get '/'
#     ...
#   end
# end
 
# Thread.new {
  EM.run {
    faye_client = Faye::Client.new("http://localhost:3000/faye")
    faye_client.subscribe('/foo') do |message|
      puts "WOAH NELLY: #{message}"
    end
  }
# }
