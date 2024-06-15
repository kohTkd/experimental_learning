# frozen_string_literal: true

module GrpcClients
  class HelloClient < ::GrpcClients::Base
    bind :auth, :hello, :hello_world
    define_message(:hello, :name)
  end
end
