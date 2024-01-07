# frozen_string_literal: true

CONTAINERS = %w[api front].freeze
COMMANDS = %w[build run up stop restart reboot].freeze

File.open("#{__dir__}/../makefiles/docker.mk", 'w') do |f|
  CONTAINERS.each do |container|
    f.puts ".PHONY: #{COMMANDS.map { |command| "docker.#{container}.#{command}" }.join(' ')}"
    COMMANDS.each do |command|
      f.puts "docker.#{container}.#{command}:"
      f.puts "\tmake docker.#{command} container='#{container}'"
    end
    f.puts ''
  end
end
