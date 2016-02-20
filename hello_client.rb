#!/usr/bin/env ruby

this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)
# $LOAD_PATH.unshift(this_dir) unless $LOAD_PATH.include?(this_dir)

require 'grpc'
require 'helloworld_services'

def main
  stub = Helloworld::Greeter::Stub.new('localhost:5000', :this_channel_is_insecure)
  GRPC.logger.info('... connecting insecurely')

  message = stub.say_hello(Helloworld::HelloRequest.new(name: 'cncgl')).message
  p "Greeting: #{message}"
end

main
