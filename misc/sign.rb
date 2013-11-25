if ARGV.length == 0
  puts "usage: sign.rb {secret} {url} {params} {body}"
  puts <<-EOS
   eg: ruby sign.rb abcde /beatvector "a=1&b=2" "{\"data\":\"a lot of it\"}"
  EOS
  exit(0)
end

require 'rubygems'
require 'openssl'
key = ARGV[0]
data = ARGV[1].to_s + "?" + ARGV[2].to_s + ARGV[3].to_s
puts OpenSSL::HMAC.hexdigest(OpenSSL::Digest::Digest.new('sha1'), key, data)
