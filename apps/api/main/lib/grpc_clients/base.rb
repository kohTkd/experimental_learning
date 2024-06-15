# frozen_string_literal: true

module GrpcClients
  class InvalidServerError < StandardError; end
  class RequiredParamsLackError < StandardError; end

  SERVERS = {
    auth: ENV.fetch('AUTH_SERVER')
  }.freeze

  class Base
    class << self
      def bind(server, *identifiers)
        endpoint = SERVERS[server]
        raise InvalidServerError, "No matched server in #{SERVER.keys.join(', ')}" if endpoint.blank?

        service = identifiers.map { |identifier| identifier.to_s.camelize }.join('::')
        client = Object.const_get("::#{service}::Stub").new(endpoint, *client_options)

        define_method :client do
          client
        end

        define_method :request_class do
          service_prefix = identifiers.first.to_s.camelize
          Object.const_get("::#{service_prefix}::#{service_prefix}Request")
        end
      end

      def define_message(method_name, *keys)
        define_method(method_name) do |params|
          unless keys.all? { |key| params.key?(key) }
            lack = keys.filter { |params| !params.key?(key) }

            raise RequiredParamsLackError, "Required params are lack: #{lack.join(', ')}"
          end

          client.public_send(method_name, request_class.new(params))
        end
      end

      private

      def client_options
        return [] if Rails.env.production?

        [:this_channel_is_insecure]
      end
    end
  end
end
