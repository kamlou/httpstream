require 'sinatra'
require 'sinatra/streaming'
require 'json'
require 'date'

last_successful_read = 0

get "/" do
  
  content_type :json
  status 202
  headers "Transfer-Encoding" => "chunked"
  
  stream do |out|
    out.errback do
      puts "last_read: #{last_successful_read}"
    end
    begin
      (last_successful_read+1..last_successful_read+1000).to_a.each do |p|
        h = {id: p, data: "this is my data at #{p}", timestamp: DateTime.now.to_s}
        out << fmt("#{h.to_json}\r\n")
        last_successful_read = p
        sleep(p%5)
      end
      out << fmt("")
    rescue StandardError => e
      puts "ex: #{e.message}"
    end
  end
end

def fmt(s)
  "#{s.size.to_s(16)}\r\n#{s}\r\n"
end
