# frozen_string_literal: true

module GrpcClients
  class InvalidServerError < StandardError; end

  SERVERS = {
    auth: ENV.fetch('AUTH_SERVER')
  }.freeze

  class Base
    class << self
      def bind(service, server)
        endpoint = SERVERS[server]
        raise InvalidServerError, "No matched server in #{SERVER.keys.join(', ')}" if endpoint.blank?

        client = Object.const_get("::#{service}::Stub").new(endpoint, *client_options)

        define_method :client do
          client
        end
      end

      def define_message(method_name, req_class)
        define_method(method_name) do |params|
          client.public_send(method_name, req_class.new(params))
        end
      end

      private

      def cleint_options
        return [] if Rails.env.production?

        [:this_channel_is_insecure]
      end
    end
  end
end
